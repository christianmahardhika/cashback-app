package main

import (
	"cashback-admin/internal/core/model/entity"
	"cashback-admin/internal/pkg/httpserver"
)

func main() {
	app := httpserver.InitServer()

	// router
	app.AddRESTHandlers(entity.Tenant{})
	app.Run()
}
