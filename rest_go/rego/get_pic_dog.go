package rego

import (
	"net/http"
	"time"
)

func GetDog(w http.ResponseWriter, r *http.Request) {
	dog := Random("/home/user/GoProjects/ServerTestDocker/rest_go/resources/pictures/dogs", time.Now().UTC().UnixNano())
	http.ServeFile(w, r, ("/home/user/GoProjects/ServerTestDocker/rest_go/resources/pictures/dogs/" + dog))
}
