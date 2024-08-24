package middleware

import (
	"gotalk/internal/json"
	"net/http"
)

func writeUnauthed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write(json.Json{
		Status: http.StatusUnauthorized,
		Message: http.StatusText(http.StatusUnauthorized),
	}.ToBytes())
}
