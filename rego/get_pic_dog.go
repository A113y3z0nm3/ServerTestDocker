package rego

import (
	"net/http"
)

func GetDog(w http.ResponseWriter, r *http.Request) {
	dog := Random("rego/resources/pictures/dogs")
	http.ServeFile(w, r, ("rego/resources/pictures/dogs/" + dog))
}
