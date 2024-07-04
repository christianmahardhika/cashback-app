package repository

import (
	"cashback-admin/internal/core/model/entity"

	"gofr.dev/pkg/gofr"
)

type repository struct {
	c *gofr.Context
}

func NewRepository(c *gofr.Context) *repository {
	return &repository{c: c}
}

func (r *repository) FindUserByID(id string) (entity.User, error) {
	// find user by id
	result := r.c.SQL.QueryRow("SELECT id, name, email, created_at, updated_at, deleted_at FROM users WHERE id = $1", id)

	var user entity.User
	err := result.Scan(&user)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
