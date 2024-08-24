package handlers

import (
	"gotalk/api/v1/errors"
	"gotalk/internal/json"
	"net/http"
	"strings"
)

func Greeter(w http.ResponseWriter, r *http.Request){
	name := r.PathValue("name")
	w.Write([]byte("Hello, " + name))
}

func Pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func PostComment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, errors.PARSING_FORM_FAILED.ToString(), errors.STATUS_CODE)
		return
	}

	content := r.FormValue("content")

	if strings.TrimSpace(content) == "" {
		// I miss C macros
		http.Error(w, errors.NOT_FOUND("Content").ToString(), errors.STATUS_CODE)
		return
	}

	// TODO: Post comment logic

	w.Write(json.Json {
		Status: 200,
		Message: "Comment posted",
	}.ToBytes())
}

func Register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		http.Error(w, errors.PARSING_FORM_FAILED.ToString(), errors.STATUS_CODE)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")

	if strings.TrimSpace(name) == "" {
		http.Error(w, errors.NOT_FOUND("Name").ToString(), errors.STATUS_CODE)
		return
	}
	if strings.TrimSpace(email) == "" {
		http.Error(w, errors.NOT_FOUND("Email").ToString(), errors.STATUS_CODE)
		return
	}

	// TODO: Register logic

	w.Write(json.Json{
		Status: 200,
		Message: "Registration complete",
	}.ToBytes())
}
