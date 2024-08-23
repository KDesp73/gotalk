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

		if !strings.HasPrefix(authorization, "Bearer ") && !strings.HasPrefix(authorization, "Admin ") {
			writeUnauthed(w)
			return
		}

		isAdmin := strings.HasPrefix(authorization, "Admin ")
		var prefix = "Bearer "
		if isAdmin {
			prefix = "Admin "
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
