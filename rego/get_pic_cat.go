package rego

import (
	"net/http"
	"math/rand"
	"time"
)

func GetCat(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UTC().UnixNano())
	cat := Random("rego/resources/pictures/cats")
	http.ServeFile(w, r, ("rego/resources/pictures/cats/" + cat))
}
