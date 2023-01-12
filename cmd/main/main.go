package main

import (
	"net/http"
	"log"

	"github.com/aguazul-marco/goProjects/pkg/routes"
	"github.com/gorilla/mux"
)


func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}