package middleware

import (
	"cashback-admin/internal/pkg/errors"
	"net/http"

	gofrHTTP "gofr.dev/pkg/gofr/http"
)

type contextKey string

const JWTToken contextKey = "jwt_token"

func ValidateJWT() gofrHTTP.Middleware {
	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token == "" {
				errors.Unauthorized(w, "missing Authorization header")
				return
			}

			inner.ServeHTTP(w, r)
		})
	}
}
