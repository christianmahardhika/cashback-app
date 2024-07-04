package main

import (
	"cashback-core/internal/core/handler"
	"cashback-core/internal/core/usecase"
	"cashback-core/internal/pkg/httpserver"

	"github.com/go-playground/validator/v10"
)

func main() {
	app := httpserver.InitServer()

	// init usecase
	uc := usecase.NewUsecase()
	v := validator.New()
	// init handlers
	h := handler.NewHandler(*uc, v)
	// router
	app.POST("/cashback", h.Calculate)
	app.Run()
}
