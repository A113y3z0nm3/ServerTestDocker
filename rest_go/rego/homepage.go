package rego

import (
	"net/http"
)

func firstPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./resources/html/client.html")
}
