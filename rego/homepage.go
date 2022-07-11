package rego

import (
	"net/http"
)

func FirstPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "rego/resources/html/client.html")
}
