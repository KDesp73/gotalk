package handlers

import (
	"fmt"
	"gotalk/api/state"
	"gotalk/api/v1/errors"
	"gotalk/internal/json"
	"gotalk/internal/threads"
	"net/http"
	"strings"
)

func IsAdmin(w http.ResponseWriter, r *http.Request){
	user := r.PathValue("user")

	w.Write(json.Json{
		Status: 200,
		Message: fmt.Sprintf("%s you are an admin!", user),
	}.ToBytes())
}

func Sudo(w http.ResponseWriter, r *http.Request) {
	user := r.PathValue("user")
	w.Write(json.Json{
		Status: 200,
		Message: fmt.Sprintf("%s is now an admin", user),
	}.ToBytes())
}

func DeleteThread(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, errors.PARSING_FORM_FAILED.ToString(), errors.STATUS_BAD_REQUEST)
		return
	}

	threadid := r.FormValue("threadid")

	if strings.TrimSpace(threadid) == "" {
		http.Error(w, errors.NOT_FOUND("Thread id").ToString(), errors.STATUS_BAD_REQUEST)
		return
	}

	succ := state.Instance.Threads.RemoveThread(threadid)
	
	if !succ {
		http.Error(w, errors.FAILED("delete thread (threadid doesn't exist)").ToString(), errors.STATUS_FAIL)
		return
	}
	
	w.Write(json.Json{
		Status: http.StatusOK,
		Message: "Thread deleted successfully",
	}.ToBytes())
}


func NewThread(w http.ResponseWriter, r *http.Request) {
	id := state.Instance.Threads.PushThread(&threads.Thread{})

	w.Write(json.Json {
		Status: http.StatusOK,
		Message: "Thread created successfully",
		Data: json.NestedJson{
			Key: id,
		},
	}.ToBytes())
}
