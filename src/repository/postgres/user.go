package postgres

import (
	"context"
	"fmt"
	"log"

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

func (r *UserPostgres) CreateUser(ctx context.Context, entity models.UserService) ( int , error ) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES ($1, $2, $3, $4, $5, $6, $7, $8 )  RETURNING id", userTable, userFeilds)
	 err := r.db.QueryRowContext(ctx, query, entity.FullName, entity.NickName, entity.Photo, entity.Birthday, entity.Location, entity.CreatedAt, entity.UpdatedAt, entity.DeletedAt).Scan(&id)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return 0 , err
	}
	return id, nil
}

func (r *UserPostgres) UpdateUser(ctx context.Context, Id int , entity models.UserUpdate) (string, error) {
	var updated_at string
	query := fmt.Sprintf("UPDATE %s SET full_name = $1, nick_name = $2, photo = $3, birth_day = $4, location = $5, updated_at = $6 WHERE id = $7 RETURNING updated_at ", userTable)
	fmt.Println(query)
	err := r.db.QueryRowContext(ctx, query, entity.FullName, entity.NickName, entity.Photo, entity.Birthday, entity.Location, entity.UpdatedAt, Id ).Scan(&updated_at)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return "" ,  err
	}
	return updated_at,  nil
}