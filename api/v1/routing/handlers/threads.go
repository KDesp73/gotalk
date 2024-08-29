package handlers

import (
	"fmt"
	"gotalk/api/v1/errors"
	"gotalk/api/v1/response"
	"gotalk/internal/json"
	"gotalk/internal/state"
	"gotalk/internal/threads"
	"net/http"
	"strings"
)

func GetThreads(w http.ResponseWriter, r* http.Request) {
	j := json.Json {
		Status: 200,
		Message: "Threads retreived successfully",
		Data: json.NestedJson {},
	}
	for _, thread := range state.Instance.Threads.Items {
		j.Data.Threads = append(j.Data.Threads, *thread)
	}

	response.Success(w, j)
}

func DeleteThread(w http.ResponseWriter, r *http.Request) {
	threadid := r.PathValue("threadid")
	threadid = strings.TrimSpace(threadid)

	if threadid == "" {
		response.Error(w, errors.NOT_SET("Thread"))
		return
	}

	succ := state.Instance.Threads.RemoveThread(threadid)
	
	if !succ {
		response.Error(w, errors.NOT_FOUND(fmt.Sprintf("Thread '%s'", threadid)))
		return
	}
	
	response.Success(w, json.Json{
		Status: 204,
		Message: "Thread deleted successfully",
	})
}


func NewThread(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		response.Error(w, errors.FAILED("parsing form"))
		return
	}

	title := r.FormValue("title")
	title = strings.TrimSpace(title)

	if title == "" {
		response.Error(w, errors.NOT_SET("Thread title"))
		return
	}

	if state.Instance.Threads.TitleExists(title) {
		response.Error(w, errors.DUPLICATE("Thread title"))
		return
	}

	id := state.Instance.Threads.PushThread(&threads.Thread{Title: title})

	response.Success(w, json.Json {
		Status: 201,
		Message: "Thread created successfully",
		Data: json.NestedJson{
			Key: id,
		},
	})
}
