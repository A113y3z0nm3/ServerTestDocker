package rego

import (
	"net/http"
)

func GetDog(w http.ResponseWriter, r *http.Request) {
	dog := Random("/home/user/GoProjects/ServerTestDocker/rest_go/resources/pictures/dogs")
	http.ServeFile(w, r, ("/home/user/GoProjects/ServerTestDocker/rest_go/resources/pictures/dogs/" + dog))
}
