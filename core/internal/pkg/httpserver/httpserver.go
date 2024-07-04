package httpserver

import (
	"cashback-core/internal/pkg/middleware"

	"gofr.dev/pkg/gofr"
)

func InitServer() *gofr.App {
	app := gofr.New()
	app.UseMiddleware(middleware.CheckTenant())
	// app.EnableAPIKeyAuthWithValidator(auth.ApiKeyValidator)
	// app.Migrate(migrations.All())
	return app
}
