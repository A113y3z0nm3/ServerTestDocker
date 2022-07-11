package main

import (
	"log"
	"net/http"

	rg "ServerTestDocker/rego"

	"github.com/gorilla/mux"
)



func main() {
	log.Printf("Server started")
	r := mux.NewRouter()
	r.HandleFunc("/", rg.FirstPage).Methods("GET")
	r.HandleFunc("/picture_cat", rg.GetCat).Methods("GET")
	r.HandleFunc("/picture_dog", rg.GetDog).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
