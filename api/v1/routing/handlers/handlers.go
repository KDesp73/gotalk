package handlers

import "net/http"

func Greeter(w http.ResponseWriter, r *http.Request){
	name := r.PathValue("name")
	w.Write([]byte("Hello, " + name))
}

func Pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

