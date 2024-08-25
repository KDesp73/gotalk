package handlers

import (
	"gotalk/api/state"
	"gotalk/api/v1/errors"
	"gotalk/api/v1/response"
	"gotalk/internal/json"
	"gotalk/internal/users"
	"gotalk/internal/utils"
	"net/http"
	"os"
	"strings"
)

// /ping
func Pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

// /
func ServeIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	content, err := os.ReadFile("./docs/index.html")

	if err != nil {
		response.Error(w, errors.FAILED("serving index.html"))
		return
	}

	w.Write(content)
}

// /comment?threadid={threadid}&userid={userid}&content=content
func PostComment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		response.Error(w, errors.FAILED("parsing form"))
		return
	}

	content := r.FormValue("content")
	userid := r.FormValue("userid")
	threadid := r.FormValue("threadid")

	if strings.TrimSpace(content) == "" {
		response.Error(w, errors.NOT_FOUND("Content"))
		return
	}
	if strings.TrimSpace(threadid) == "" {
		response.Error(w, errors.NOT_FOUND("Thread id"))
		return
	}
	if strings.TrimSpace(userid) == "" {
		response.Error(w, errors.NOT_FOUND("User id"))
		return
	}

	thread := state.Instance.Threads.Get(threadid)

	if thread.ID != threadid {
		response.Error(w, errors.FAILED("finding thread"))
		return
	}
	
	if thread == nil {
		response.Error(w, errors.INVALID("Thread id"))
		return
	}

	user := state.Instance.Users.Get(userid)

	if user == nil {
		response.Error(w, errors.INVALID("Thread id"))
		return
	}

	thread.AppendComment(user.Name, content)

	response.Success(w, json.Json {
		Status: 201,
		Message: "Comment posted",
	})
}

// /register?name={name}&email={email}
func Register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		response.Error(w, errors.FAILED("parsing form"))
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")

	if strings.TrimSpace(name) == "" {
		response.Error(w, errors.NOT_FOUND("Name"))
		return
	}
	if strings.TrimSpace(email) == "" {
		response.Error(w, errors.NOT_FOUND("Email"))
		return
	}
	if !utils.IsValidEmail(email) {
		response.Error(w, errors.INVALID("email"))
		return
	}
	if state.Instance.Users.NameExists(name) {
		response.Error(w, errors.DUPLICATE("name"))
		return
	}
	if state.Instance.Users.EmailExists(email) {
		response.Error(w, errors.DUPLICATE("email"))
		return
	}

	key := state.Instance.Users.PushUser(&users.User{
		Name: name,
		Email: email,
		Type: users.USER_DEFAULT,
		SignUpTime: utils.CurrentTimestamp(),
	})

	response.Success(w, json.Json{
		Status: 201,
		Message: "Registration complete",
		Data: json.NestedJson{
			Key: key,
		},
	})
}
