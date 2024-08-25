package handlers

import (
	"fmt"
	"gotalk/api/state"
	"gotalk/api/v1/response"
	"gotalk/api/v1/errors"
	"gotalk/internal/encryption"
	"gotalk/internal/json"
	"gotalk/internal/threads"
	"net/http"
	"strings"
)

func IsAdmin(w http.ResponseWriter, r *http.Request){
	user := r.PathValue("user")

	response.Success(w, json.Json{
		Status: 200,
		Message: fmt.Sprintf("%s you are an admin!", user),
	})
}

func UndoSudo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		response.Error(w, errors.FAILED("parsing form"))
		return
	}

	id := r.FormValue("id")
	succ := state.Instance.Users.Sudo(encryption.Hash(id), true)

	if !succ {
		response.Error(w, errors.FAILED("revoking admin privileges"))
		return
	}

	response.Success(w, json.Json{
		Status: 200,
		Message: "Admin privileges revoked",
	})
}

func Sudo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		response.Error(w, errors.FAILED("parsing form"))
		return
	}

	id := r.FormValue("id")

	succ := state.Instance.Users.Sudo(encryption.Hash(id), false)

	if !succ {
		response.Error(w, errors.FAILED("granding admin privileges"))
		return
	}

	response.Success(w, json.Json{
		Status: 200,
		Message: "Admin privileges granted",
	})
}

func DeleteThread(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		response.Error(w, errors.FAILED("parsing form"))
		return
	}

	threadid := r.FormValue("threadid")

	if strings.TrimSpace(threadid) == "" {
		response.Error(w, errors.NOT_FOUND("Thread"))
		return
	}

	succ := state.Instance.Threads.RemoveThread(threadid)
	
	if !succ {
		response.Error(w, errors.FAILED("deleting thread"))
		return
	}
	
	response.Success(w, json.Json{
		Status: 204,
		Message: "Thread deleted successfully",
	})
}


func NewThread(w http.ResponseWriter, r *http.Request) {
	id := state.Instance.Threads.PushThread(&threads.Thread{})

	response.Success(w, json.Json {
		Status: 201,
		Message: "Thread created successfully",
		Data: json.NestedJson{
			Key: id,
		},
	})
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		response.Error(w, errors.FAILED("parsing form"))
		return
	}

	id := r.FormValue("id")
	threadid := r.FormValue("threadid")

	thread := state.Instance.Threads.Get(threadid)
	index := thread.SearchCommentID(id)
	succ := thread.RemoveComment(index)

	if !succ {
		response.Error(w, errors.NOT_FOUND("Comment"))
		return
	}

	response.Success(w, json.Json {
		Status: 204,
		Message: "Comment removed successfully",
	})
}
