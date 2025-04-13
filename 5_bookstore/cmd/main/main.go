package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/ayushwar/go/5_bookstore/pkg/routes"
)

func main()  {
	r:=mux.NewRouter()
	routes.registerbookroutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:8000",r))

	
}