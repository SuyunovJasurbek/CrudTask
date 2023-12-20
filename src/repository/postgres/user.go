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
// single
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

// multi
func (r *UserPostgres) CreateUsers(ctx context.Context, users []models.UserService) ( error) {
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES ", userTable, userFeilds)
	vals := []interface{}{}
	for _, user := range users {
		t := fmt.Sprintf("('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s'),",
			user.FullName,
			user.NickName,
			user.Photo,
			user.Birthday,
			user.Location,
			user.CreatedAt,
			user.UpdatedAt,
			user.DeletedAt)
		query += t
		vals = append(vals,
			user.FullName,
			user.NickName,
			user.Photo,
			user.Birthday,
			user.Location,
			user.CreatedAt,
			user.UpdatedAt,
			user.DeletedAt)
	}
	query = query[:len(query)-1]
	fmt.Println(query)
	_, err := r.db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error multi create users method: %v", err)
		return nil 
	}

	return nil
}
func (r *UserPostgres) UpdateUsers(ctx context.Context, users []models.UsersUpdate) ( error) {
	query := fmt.Sprintf("UPDATE %s SET full_name = temp.full_name, nick_name = temp.nick_name, photo = temp.photo, birth_day = temp.birth_day, location = temp.location, updated_at = temp.updated_at FROM (VALUES ", userTable)
	vals := []interface{}{}
	for _, user := range users {
		t := fmt.Sprintf("(%d, '%s', '%s', '%s', '%s', '%s', '%s'),",
			user.ID,
			user.FullName,
			user.NickName,
			user.Photo,
			user.Birthday,
			user.Location,
			user.UpdatedAt)
		query += t
		vals = append(vals,
			user.ID,
			user.FullName,
			user.NickName,
			user.Photo,
			user.Birthday,
			user.Location,
			user.UpdatedAt)
	}
	query = query[:len(query)-1]
	query += ") AS temp(id, full_name, nick_name, photo, birth_day, location, updated_at) WHERE temp.id = users.id"
	fmt.Println(query)
	_, err := r.db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error multi update users method: %v", err)
		return nil 
	}

	return nil
}
func ( r *UserPostgres) DeleteUsers(ctx context.Context, Ids []models.UserIds, deleted_at string) ( error) {
	query := fmt.Sprintf("UPDATE %s SET deleted_at = $1 WHERE id IN (", userTable)
	vals := []interface{}{}
	for _, user := range Ids {
		t := fmt.Sprintf("%d,", user.ID)
		query += t
		vals = append(vals, user.ID)
	}
	query = query[:len(query)-1]
	query += ")"
	fmt.Println(query)
	_, err := r.db.ExecContext(ctx, query, deleted_at)
	if err != nil {
		log.Printf("Error multi delete users method: %v", err)
		return nil 
	}

	return nil
}