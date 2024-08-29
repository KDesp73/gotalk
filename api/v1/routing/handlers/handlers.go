package handlers

import (
	"gotalk/api/v1/errors"
	"gotalk/api/v1/response"
	"net/http"
	"os"
)

func Pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	var file string
	if r.URL.Path == "/dark" {
		file = "./docs/dark/index.html"
	} else if r.URL.Path == "/" || r.URL.Path == "/light" {
		file = "./docs/index.html"
	} else {
		file = "./docs/404/index.html"
	}

	content, err := os.ReadFile(file)

	if err != nil {
		response.Error(w, errors.FAILED("serving index.html"))
		return
	}

	w.Write(content)
}
