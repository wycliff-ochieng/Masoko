package middleware

import "net/http"

func CartMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	}
}

func GetUserIDFromContext() error { return nil }
