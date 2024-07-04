package httpserver

import (
	"cashback-admin/internal/pkg/middleware"
	"cashback-admin/migrations"

	"gofr.dev/pkg/gofr"
)

func InitServer() *gofr.App {
	app := gofr.New()
	app.UseMiddleware(middleware.ValidateJWT())

	app.Migrate(migrations.All())
	return app
}
