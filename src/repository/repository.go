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
	UpdateUserById(ctx context.Context, Id int , entity models.UserUpdate) ( string , error)
	DeleteUserById(ctx context.Context, Id int, deleted_at string) (string, error)
	GetUserById(ctx context.Context, Id int ) ( *models.GetUser , error)
	GetUsers(ctx context.Context) ( []*models.GetUser , error)
	GetUsersByFieldCreatedAt(ctx context.Context, created_at string) ([]*models.GetUser, error)
	GetUsersByFieldUpdatedAt(ctx context.Context, updated_at string) ([]*models.GetUser, error)
	GetUserFullnameSort(ctx context.Context) ([]*models.GetUser, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserI: postgres.NewUserPostgres(db),
	}
}
