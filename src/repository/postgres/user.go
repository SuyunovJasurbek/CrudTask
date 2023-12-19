package postgres

import (
	"context"
	"fmt"

	"github.com/SuyunovJasurbek/CrudTask/models"
	"github.com/jmoiron/sqlx"
)

const (
	userTable  = "users"
	userFeilds = "full_name , nick_name , photo , birth_day , location , created_at , updated_at , deleted_at"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(ctx context.Context, entity models.UserService) error {
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) ", userTable, userFeilds)
	fmt.Println(query)
	r.db.QueryRowContext(ctx, query,
		entity.FullName,
		entity.NickName,
		entity.Photo,
		entity.Birthday,
		entity.Location,
		entity.CreatedAt,
		entity.UpdatedAt,
		entity.DeletedAt)

	return nil
}
