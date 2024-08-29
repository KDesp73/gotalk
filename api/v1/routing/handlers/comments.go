package handlers

import (
	"gotalk/api/v1/errors"
	"gotalk/api/v1/response"
	"gotalk/internal/json"
	"gotalk/internal/state"
	"net/http"
	"strings"
)


func GetComments(w http.ResponseWriter, r* http.Request) {
	err := r.ParseForm()
	if err != nil {
		response.Error(w, errors.FAILED("parsing form"))
		return
	}

	threadid := r.FormValue("threadid")
	threadid = strings.TrimSpace(threadid)

	if threadid == "" {
		response.Error(w, errors.NOT_FOUND("Thread id"))
		return
	}

	j := json.Json {
		Status: 200,
		Message: "Comments retreived successfully",
		Data: json.NestedJson {},
	}
	for _, comment := range state.Instance.Threads.Get(threadid).Comments {
		j.Data.Comments = append(j.Data.Comments, *comment)
	}

	response.Success(w, j)
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := r.ParseForm()
	if err != nil {
		response.Error(w, errors.FAILED("parsing form"))
		return
	}

	threadid := r.FormValue("threadid")

	if strings.TrimSpace(threadid) == "" {
		response.Error(w, errors.NOT_FOUND("Thread id"))
		return
	}

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

func PostComment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		response.Error(w, errors.FAILED("parsing form"))
		return
	}

	userid := r.PathValue("userid")
	content := r.FormValue("content")
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

	thread.PushComment(user.Name, content)

	response.Success(w, json.Json {
		Status: 201,
		Message: "Comment posted",
	})
}
