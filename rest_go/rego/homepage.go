package rego

import (
	"net/http"
)

func FirstPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./resources/html/client.html")
}
