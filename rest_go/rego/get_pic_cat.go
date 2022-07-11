package rego

import (
	"net/http"
	"time"
)

func GetCat(w http.ResponseWriter, r *http.Request) {
	cat := Random("/home/user/GoProjects/ServerTestDocker/rest_go/resources/pictures/cats", time.Now().UTC().UnixNano())
	http.ServeFile(w, r, ("/home/user/GoProjects/ServerTestDocker/rest_go/resources/pictures/cats/" + cat))
}
