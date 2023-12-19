package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	UserI
}

type UserI interface {
	CreateUser(ctx context.Context)int 
	// UpdateUser(ctx context.Context)
	// DeleteUser(ctx context.Context)
	// GetIdUser(ctx context.Context)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserI: NewUserPostgres(db),
	}
}