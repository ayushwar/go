package main

import (
	// "fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	PORTsTRING := os.Getenv("PORT")
	// if PORTsTRING == "" {
	// 	log.Fatal("not found port")
	// }

	router := chi.NewRouter()
	srv := &http.Server{
		Handler: router,
		Addr: ":"+PORTsTRING,
	}

	println("serevr start")
	err:=srv.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}
	// fmt.Println("port",PORTsTRING)
}
