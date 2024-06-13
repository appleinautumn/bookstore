package middleware

import (
	"net/http"

	"gotu/bookstore/internal/util"
)

func Authenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userIDHeader := r.Header.Get("user_id")

		if userIDHeader == "" {
			util.WriteError(w, http.StatusUnauthorized, "User must be authenticated")
			return
		}

		next.ServeHTTP(w, r)
	})
}
