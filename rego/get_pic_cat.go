package rego

import (
	"net/http"
)

func GetCat(w http.ResponseWriter, r *http.Request) {
	cat := Random("rego/resources/pictures/cats")
	http.ServeFile(w, r, ("rego/resources/pictures/cats/" + cat))
}
