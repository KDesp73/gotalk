package response

import (
	"gotalk/internal/json"
	"net/http"
)

func Success(w http.ResponseWriter, j json.Json) {
	w.WriteHeader(j.Status)
	w.Write(j.ToBytes())
}

func Error(w http.ResponseWriter, j json.Json) {
	http.Error(w, j.ToString(), j.Status)
}
