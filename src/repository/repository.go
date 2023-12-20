package repository

import (
	"context"

	"github.com/SuyunovJasurbek/CrudTask/models"
	"github.com/SuyunovJasurbek/CrudTask/src/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	UserI
}

type UserI interface {
	CreateUser( ctx context.Context,entity models.UserService) (int , error)
	UpdateUser(ctx context.Context, Id int , entity models.UserUpdate) ( string , error)
	// DeleteUser(ctx context.Context)
	// GetIdUser(ctx context.Context)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserI: postgres.NewUserPostgres(db),
	}
}
