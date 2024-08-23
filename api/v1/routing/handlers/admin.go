package handlers

import (
	"fmt"
	"net/http"
)

func IsAdmin(w http.ResponseWriter, r *http.Request){
	user := r.PathValue("user")
	w.Write([]byte(fmt.Sprintf("%s you are an admin!", user)))
}

func Sudo(w http.ResponseWriter, r *http.Request) {
	user := r.PathValue("user")
	w.Write([]byte(fmt.Sprintf("%s is now an admin", user)))
}
