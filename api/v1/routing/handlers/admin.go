package handlers

import (
	"fmt"
	"gotalk/api/state"
	"gotalk/internal/json"
	"gotalk/internal/threads"
	"net/http"
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
}


func NewThread(w http.ResponseWriter, r *http.Request) {
	id := state.Instance.Threads.PushThread(&threads.Thread{})

	w.Write(json.Json {
		Status: http.StatusAccepted,
		Message: "Thread created",
		Data: json.NestedJson{
			Key: id,
		},
	}.ToBytes())
}
