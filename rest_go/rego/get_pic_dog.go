package rego

import (
	"net/http"
)

func GetDog(w http.ResponseWriter, r *http.Request) {
	dog := Random("./resources/pictures/dogs/")
	http.ServeFile(w, r, ("./resources/pictures/dogs/" + dog))
}
