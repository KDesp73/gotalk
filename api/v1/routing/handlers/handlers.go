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

// /comment?threadid={threadid}&userid={userid}&content=content
func PostComment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, errors.PARSING_FORM_FAILED.ToString(), errors.STATUS_BAD_REQUEST)
		return
	}

	content := r.FormValue("content")
	userid := r.FormValue("userid")
	threadid := r.FormValue("threadid")

	if strings.TrimSpace(content) == "" {
		// I miss C macros
		http.Error(w, errors.NOT_FOUND("Content").ToString(), errors.STATUS_BAD_REQUEST)
		return
	}
	if strings.TrimSpace(threadid) == "" {
		http.Error(w, errors.NOT_FOUND("Thread id").ToString(), errors.STATUS_BAD_REQUEST)
		return
	}
	if strings.TrimSpace(userid) == "" {
		http.Error(w, errors.NOT_FOUND("User id").ToString(), errors.STATUS_BAD_REQUEST)
		return
	}

	thread := state.Instance.Threads.Get(threadid)
	
	if thread == nil {
		http.Error(w, errors.INVALID_THREAD_ID.ToString(), errors.STATUS_BAD_REQUEST)
		return
	}

	user := state.Instance.Users.Get(userid)

	if user == nil {
		http.Error(w, errors.INVALID_THREAD_ID.ToString(), errors.STATUS_BAD_REQUEST)
		return
	}

	thread.AppendComment(user.Name, content)

	w.Write(json.Json {
		Status: 200,
		Message: "Comment posted",
	}.ToBytes())
}

// /register?name={name}&email={email}
func Register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		http.Error(w, errors.PARSING_FORM_FAILED.ToString(), errors.STATUS_BAD_REQUEST)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")

	if strings.TrimSpace(name) == "" {
		http.Error(w, errors.NOT_FOUND("Name").ToString(), errors.STATUS_BAD_REQUEST)
		return
	}
	if strings.TrimSpace(email) == "" {
		http.Error(w, errors.NOT_FOUND("Email").ToString(), errors.STATUS_BAD_REQUEST)
		return
	}
	if !utils.IsValidEmail(email) {
		http.Error(w, errors.INVALID_EMAIL.ToString(), errors.INVALID_EMAIL.Status)
		return
	}
	if state.Instance.Users.EmailExists(email) {
		http.Error(w, errors.DUPLICATE_EMAIL.ToString(), errors.DUPLICATE_EMAIL.Status)
		return
	}

	key := state.Instance.Users.PushUser(&users.User{
		Name: name,
		Email: email,
		Type: users.USER_DEFAULT,
		SignUpTime: utils.CurrentTimestamp(),
	})

	w.Write(json.Json{
		Status: http.StatusAccepted,
		Message: "Registration complete",
		Data: json.NestedJson{
			Key: key,
		},
	}.ToBytes())
}
