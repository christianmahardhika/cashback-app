package httpserver

import (
	"cashback-admin/migrations"

	"gofr.dev/pkg/gofr"
)

func InitServer() *gofr.App {
	app := gofr.New()
	// app.UseMiddleware(middleware.CheckTenant())

	// app.EnableAPIKeyAuthWithValidator(auth.ApiKeyValidator)
	app.Migrate(migrations.All())
	return app
}
