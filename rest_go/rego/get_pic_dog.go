package rego

import (
	"net/http"
)

func getDog(w http.ResponseWriter, r *http.Request) {
	dog := random("./resources/pictures/dogs/")
	http.ServeFile(w, r, ("./resources/pictures/dogs/" + dog))
}
