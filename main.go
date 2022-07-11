package main

import (
	"log"
	"net/http"

	rg "ServerTestDocker/rest_go/rego"

	"github.com/gorilla/mux"
)

func main() {
	log.Printf("Server started")
	r := mux.NewRouter()
	r.HandleFunc("/", rg.firstPage).Methods("GET")
	r.HandleFunc("/picture_cat", rg.getCat).Methods("GET")
	r.HandleFunc("/picture_dog", rg.getDog).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
