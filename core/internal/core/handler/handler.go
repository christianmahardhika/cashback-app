package handler

import (
	"cashback-core/internal/core/model/request"
	"cashback-core/internal/core/usecase"

	"github.com/go-playground/validator/v10"
	"gofr.dev/pkg/gofr"
)

type Handler struct {
	uc       usecase.Usecase
	validate *validator.Validate
}

func NewHandler(uc usecase.Usecase, validate *validator.Validate) *Handler {
	return &Handler{
		uc:       uc,
		validate: validate,
	}
}

func (h *Handler) Calculate(c *gofr.Context) (interface{}, error) {
	// validate request
	var payload request.Calculate
	err := c.Bind(&payload)
	if err != nil {
		return nil, err
	}

	err = h.validate.Struct(payload)
	if err != nil {
		return nil, err
	}

	// call usecase
	result, err := h.uc.Calculate(c)
	if err != nil {
		return nil, err
	}

	return result, nil
}
