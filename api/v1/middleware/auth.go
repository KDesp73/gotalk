package middleware

import (
	"gotalk/internal/state"
	"log"
	"net/http"
	"strings"
)

const AuthUserID = "middleware.auth.userID"

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

		if !state.Instance.Users.IsAdmin(encodedToken) {
			writeUnauthed(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func EnsureAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		prefix := "Bearer "

		if !strings.HasPrefix(authorization, prefix) {
			writeUnauthed(w)
			return
		}

		encodedToken := strings.TrimPrefix(authorization, prefix)

		match := false
		for _, user := range state.Instance.Users.Users {
			if encodedToken == user.Key {
				match = true
			}
		}

		if !match {
			writeUnauthed(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
