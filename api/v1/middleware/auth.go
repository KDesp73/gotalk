package middleware

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"
)

const AuthUserID = "middleware.auth.userID"

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		prefix := "Bearer "

		if !strings.HasPrefix(authorization, prefix) {
			writeUnauthed(w)
			return
		}

		encodedToken := strings.TrimPrefix(authorization, prefix)

		token, err := base64.StdEncoding.DecodeString(encodedToken)
		if err != nil {
			writeUnauthed(w)
			return
		}

		userID := string(token)
		log.Printf("Authorization successful (userID: %s)", userID)

		next.ServeHTTP(w, r)
	})
}
