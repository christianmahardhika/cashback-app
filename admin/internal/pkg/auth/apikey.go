package auth

import (
	"cashback-admin/internal/core/model/entity"

	"gofr.dev/pkg/gofr/container"
)

func ApiKeyValidator(c *container.Container, apiKey string) bool {
	// find api key in database

	result := c.SQL.QueryRow("SELECT id, name, api_key, created_at, updated_at, deleted_at FROM tenants WHERE api_key = $1")

	var tenant entity.Tenant
	err := result.Scan(&tenant)
	if err != nil {
		c.Logger.Error("error while fetching api key from database", err)
		return false
	}
	if tenant.ApiKey != apiKey {
		return false
	}

	return true
}
