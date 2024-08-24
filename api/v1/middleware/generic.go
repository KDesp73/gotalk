package middleware

import (
	"gotalk/api/state"
	"gotalk/internal/encryption"
	"log"
	"net/http"
	"strings"
)

func EnsureAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Checking if user is admin")
		authorization := r.Header.Get("Authorization")
		prefix := "Bearer "

		if !strings.HasPrefix(authorization, prefix) {
			writeUnauthed(w)
			return
		}

		encodedToken := strings.TrimPrefix(authorization, prefix)
		hashedToken := encryption.Hash(encodedToken)

		if !state.Instance.Users.IsAdmin(hashedToken) {
			writeUnauthed(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func LoadUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Loading user")
		next.ServeHTTP(w, r)
	})
}
