package middleware

import (
	"context"
	"net/http"

	gofrHTTP "gofr.dev/pkg/gofr/http"
)

type contextKey string

const TenantIDKey contextKey = "tenant_id"

func CheckTenant() gofrHTTP.Middleware {
	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// check tenant
			// if tenant not found return 404
			tenant := r.Header.Get("X-Tenant-ID")
			if tenant == "" {
				http.Error(w, "tenant not found", http.StatusNotFound)
				return
			}
			// if tenant found set tenant in context
			ctx := context.WithValue(r.Context(), TenantIDKey, tenant)
			r = r.WithContext(ctx)

			inner.ServeHTTP(w, r)
		})
	}
}
