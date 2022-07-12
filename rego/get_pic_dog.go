package rego

import (
	"net/http"
	"math/rand"
	"time"
)

func GetDog(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UTC().UnixNano())
	dog := Random("rego/resources/pictures/dogs")
	http.ServeFile(w, r, ("rego/resources/pictures/dogs/" + dog))
}
