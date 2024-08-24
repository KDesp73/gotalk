package middleware

import (
	"gotalk/api/state"
	"gotalk/internal/encryption"
	"net/http"
	"strings"
)

const AuthUserID = "middleware.auth.userID"

func EnsureAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		prefix := "Bearer "

		if !strings.HasPrefix(authorization, prefix) {
			writeUnauthed(w)
			return
		}

		encodedToken := strings.TrimPrefix(authorization, prefix)

		hashedToken := encryption.Hash(encodedToken)

		match := false
		for _, user := range state.Instance.Users.Users {
			if hashedToken == user.Key {
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
