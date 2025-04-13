package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"

)
func formpagehandeler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "statics/form.html")
		return
	}

	if r.Method=="POST"{
		err:=r.ParseForm()
		if err!=nil{
			log.Fatal(err)

		}

		name:=r.FormValue("name")
		age:=r.FormValue("age")

		fmt.Fprintln(w,"form submitted succesfully")

		fmt.Fprintln(w,"name=",name)
		fmt.Fprintln(w,"age",age)


	}

}






func indexhandeler(w http.ResponseWriter , r* http.Request){
if r.URL.Path!="/index"{
	http.Error(w,"404 not found",http.StatusNotFound)
	return

}
if r.Method!="GET"{
	http.Error(w,"not a valid",http.StatusNotFound)
	return
}
http.ServeFile(w,r,"statics/index.html")

}



func main(){
	err:=godotenv.Load()
	if err != nil {
        log.Fatal("Error loading .env file")

    }


	port:=os.Getenv("PORT")
	if port==""{
		port="8080"
	}

	fileserver:=http.FileServer(http.Dir("statics"))
	http.Handle("/",fileserver)
	http.HandleFunc("/index",indexhandeler)
	http.HandleFunc("/form",formpagehandeler)
	fmt.Println("serever is starting on port 8000")
	err =http.ListenAndServe(":"+port,nil)
	if err!=nil{
		log.Fatal(err)
	}


}