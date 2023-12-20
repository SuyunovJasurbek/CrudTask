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

func (r *UserPostgres) CreateUser(ctx context.Context, entity models.UserService) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES ($1, $2, $3, $4, $5, $6, $7, $8 )  RETURNING id", userTable, userFeilds)
	err := r.db.QueryRowContext(ctx, query,
		entity.FullName,
		entity.NickName,
		entity.Photo,
		entity.Birthday,
		entity.Location,
		entity.CreatedAt,
		entity.UpdatedAt,
		entity.DeletedAt,
	).Scan(&id)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return 0, err
	}
	return id, nil
}

func (r *UserPostgres) UpdateUserById(ctx context.Context, Id int, entity models.UserUpdate) (string, error) {
	var updated_at string
	query := fmt.Sprintf("UPDATE %s SET full_name = $1, nick_name = $2, photo = $3, birth_day = $4, location = $5, updated_at = $6 WHERE id = $7 RETURNING updated_at ", userTable)
	fmt.Println(query)
	err := r.db.QueryRowContext(ctx, query,
		entity.FullName,
		entity.NickName,
		entity.Photo,
		entity.Birthday,
		entity.Location,
		entity.UpdatedAt,
		Id,
	).Scan(&updated_at)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return "", err
	}
	return updated_at, nil
}
func (r *UserPostgres) DeleteUserById(ctx context.Context, Id int, deleted_at string) (string, error) {
	query := fmt.Sprintf("UPDATE %s SET deleted_at = $1 WHERE id = $2 RETURNING deleted_at ", userTable)
	err := r.db.QueryRowContext(ctx, query,
		deleted_at,
		Id,
	).Scan(&deleted_at)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return "", err
	}
	return deleted_at, nil
}

func (r *UserPostgres) GetUserById(ctx context.Context, Id int) (*models.GetUser, error) {
	var user models.GetUser
	query := fmt.Sprintf("SELECT id , full_name , nick_name , photo , birth_day,location, created_at , updated_at FROM %s WHERE id = $1 ", userTable)
	err := r.db.GetContext(ctx, &user, query, Id)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return nil, err
	}
	return &user, nil
}

func (r *UserPostgres) GetUsers(ctx context.Context) ([]*models.GetUser, error) {
	var users []*models.GetUser
	query := fmt.Sprintf("SELECT id , full_name , nick_name , photo , birth_day,location, created_at , updated_at FROM %s ", userTable)
	err := r.db.SelectContext(ctx, &users, query)
	if err != nil {
		log.Printf("Error getting users: %v", err)
		return nil, err
	}
	return users, nil
}

func (r *UserPostgres) GetUsersByFieldCreatedAt(ctx context.Context, created_at string) ([]*models.GetUser, error) {
	var users []*models.GetUser
	query := fmt.Sprintf("SELECT id , full_name , nick_name , photo , birth_day,location, created_at , updated_at FROM %s WHERE created_at = $1 ", userTable)
	err := r.db.SelectContext(ctx, &users, query, created_at)
	if err != nil {
		log.Printf("Error getting users createdat method: %v", err)
		return nil, err
	}
	return users, nil
}

func (r *UserPostgres) GetUsersByFieldUpdatedAt(ctx context.Context, updated_at string) ([]*models.GetUser, error) {
	var users []*models.GetUser
	query := fmt.Sprintf("SELECT id , full_name , nick_name , photo , birth_day,location, created_at , updated_at FROM %s WHERE updated_at = $1 ", userTable)
	err := r.db.SelectContext(ctx, &users, query, updated_at)
	if err != nil {
		log.Printf("Error getting users updatedat method: %v", err)
		return nil, err
	}
	return users, nil
}

func (r *UserPostgres) GetUserFullnameSort(ctx context.Context) ([]*models.GetUser, error) {
	var users []*models.GetUser
	query := fmt.Sprintf("SELECT id , full_name , nick_name , photo , birth_day,location, created_at , updated_at FROM %s ORDER BY full_name ASC , id DESC ", userTable)
	err := r.db.SelectContext(ctx, &users, query)
	if err != nil {
		log.Printf("Error getting users fullname sort method: %v", err)
		return nil, err
	}
	return users, nil
}