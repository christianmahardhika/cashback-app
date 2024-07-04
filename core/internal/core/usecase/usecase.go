package usecase

import (
	"cashback-core/internal/core/model/response"
	ctx "cashback-core/internal/pkg/context"

	"gofr.dev/pkg/gofr"
)

type Usecase struct {
}

func NewUsecase() *Usecase {
	return &Usecase{}
}

func (u *Usecase) Calculate(c *gofr.Context) (interface{}, error) {
	tenantID := ctx.GetTenantIDKey(c)
	// do the calculation logic here

	c.Logger.Info("Success Calculate Cashback of Tenant: ", tenantID)
	return response.Calculate{
		Cashback: 100,
	}, nil
}
