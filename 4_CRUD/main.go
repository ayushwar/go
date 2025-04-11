package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Movie struct defines the structure of each movie in the list
type Movie struct {
	ID       string    `json:"id"`       // Unique ID of the movie
	Isbn     string    `json:"isbn"`     // ISBN of the movie (as string)
	Title    string    `json:"title"`    // Title of the movie
	Director *Director `json:"director"` // Pointer to a Director struct
}

// Director struct defines the structure of the director of a movie
type Director struct {
	Firstname string `json:"firstname"` // First name of the director
	Lastname  string `json:"lastname"`  // Last name of the director
}

// movies is our in-memory "database"
var movies []Movie

// getmovie handles GET requests and returns all movies in JSON format
func getmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies) // Encode the entire slice of movies and return
}

// creatmovie handles POST requests and adds a new movie to the list
func creatmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie) // Decode request JSON body into movie struct
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	movie.ID = fmt.Sprintf("%d", len(movies)+1) // Assign a new ID based on the length
	movies = append(movies, movie)              // Add movie to the list

	json.NewEncoder(w).Encode(movie) // Return the added movie
}

// updatemovie handles PUT requests and updates an existing movie by ID
func updatemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get route parameters (like ID)

	for index, item := range movies {
		if item.ID == params["id"] {
			// Remove the old movie entry
			movies = append(movies[:index], movies[index+1:]...)

			var updatedMovie Movie
			err := json.NewDecoder(r.Body).Decode(&updatedMovie) // Decode new movie data
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			updatedMovie.ID = params["id"]         // Keep the same ID
			movies = append(movies, updatedMovie)  // Add the updated movie to the list

			json.NewEncoder(w).Encode(updatedMovie) // Return updated movie
			return
		}
	}

	// If movie not found, return error
	http.Error(w, "Movie not found", http.StatusNotFound)
}

// delmovie handles DELETE requests to remove a movie by ID
func delmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get route parameters

	for index, item := range movies {
		if item.ID == params["id"] {
			// Remove the movie at the given index
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies) // Return updated movie list
}

// main function sets up the server and routes
func main() {
	r := mux.NewRouter()

	// Add some initial movies to the in-memory slice
	movies = append(movies, Movie{
		ID: "1", Isbn: "87870", Title: "One Punch",
		Director: &Director{Firstname: "Ayush", Lastname: "Ahirwar"},
	})
	movies = append(movies, Movie{
		ID: "2", Isbn: "65478", Title: "Kick Me",
		Director: &Director{Firstname: "Ahirwar", Lastname: "Ayush"},
	})

	// Define API routes and their handler functions
	r.HandleFunc("/movies", getmovie).Methods("GET")
	r.HandleFunc("/movies/{id}", getmovie).Methods("GET")
	r.HandleFunc("/movies", creatmovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updatemovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", delmovie).Methods("DELETE")

	fmt.Println("Starting server on localhost:8000")

	log.Fatal(http.ListenAndServe(":8000", r)) // Start the server on port 8000
}

