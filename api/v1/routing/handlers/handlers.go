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
	content, err := os.ReadFile("./docs/index.html")

	if err != nil {
		response.Error(w, errors.FAILED("serving index.html"))
		return
	}

	w.Write(content)
}
