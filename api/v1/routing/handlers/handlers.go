package handlers

import (
	"gotalk/api/state"
	"gotalk/api/v1/errors"
	"gotalk/internal/json"
	"gotalk/internal/users"
	"gotalk/internal/utils"
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
	threadid := r.FormValue("threadid")

	if strings.TrimSpace(content) == "" {
		// I miss C macros
		http.Error(w, errors.NOT_FOUND("Content").ToString(), errors.STATUS_CODE)
		return
	}
	if strings.TrimSpace(threadid) == "" {
		http.Error(w, errors.NOT_FOUND("Thread id").ToString(), errors.STATUS_CODE)
		return
	}

	thread := state.Instance.Threads.Get(threadid)
	
	if thread == nil {
		http.Error(w, errors.INVALID_THREAD_ID.ToString(), errors.STATUS_CODE)
		return
	}

	// thread.AppendComment()

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

	// TODO: validate email

	key := state.Instance.Users.PushUser(&users.User{
		Name: name,
		Email: email,
		Type: users.USER_DEFAULT,
		SignUpTime: utils.CurrentTimestamp(),
	})

	w.Write(json.Json{
		Status: 200,
		Message: "Registration complete",
		Data: json.NestedJson{
			Key: key,
		},
	}.ToBytes())
}
