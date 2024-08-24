package handlers

import (
	"fmt"
	"gotalk/internal/json"
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
