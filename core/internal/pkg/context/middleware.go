package ctx

import (
	"cashback-core/internal/pkg/middleware"
	"context"
)

func GetTenantIDKey(c context.Context) string {
	val, ok := c.Value(middleware.TenantIDKey).(string)
	if !ok {
		return ""
	}
	return val
}
