package rego

import (
	"net/http"
)

func GetCat(w http.ResponseWriter, r *http.Request) {
	cat := Random("/home/user/GoProjects/ServerTestDocker/rest_go/resources/pictures/cats")
	http.ServeFile(w, r, ("/home/user/GoProjects/ServerTestDocker/rest_go/resources/pictures/cats/" + cat))
}
