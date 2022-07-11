package rego

import (
	"net/http"
)

func GetCat(w http.ResponseWriter, r *http.Request) {
	cat := Random("./resources/pictures/cats/")
	http.ServeFile(w, r, ("./resources/pictures/cats/" + cat))
}
