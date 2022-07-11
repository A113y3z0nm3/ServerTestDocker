package rego

import (
	"net/http"
)

func getCat(w http.ResponseWriter, r *http.Request) {
	cat := random("./resources/pictures/cats/")
	http.ServeFile(w, r, ("./resources/pictures/cats/" + cat))
}
