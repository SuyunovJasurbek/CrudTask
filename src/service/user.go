package service

import (
	"context"
	"time"

	"github.com/SuyunovJasurbek/CrudTask/models"
)

func (s *Service) CreateUser(ctx context.Context, usr models.UserHandler) (int, error) {
	var user = models.UserService{
		FullName:  usr.FullName,
		NickName:  usr.NickName,
		Photo:     usr.Photo,
		Birthday:  usr.Birthday,
		Location:  usr.Location,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		DeletedAt: "",
	}

	id, err := s.repo.UserI.CreateUser(ctx, user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Service) UpdateUser(ctx context.Context, Id int, usr models.UserHandler) (string, error) {
	var user = models.UserUpdate{
		FullName:  usr.FullName,
		NickName:  usr.NickName,
		Photo:     usr.Photo,
		Birthday:  usr.Birthday,
		Location:  usr.Location,
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	updated_at, err := s.repo.UserI.UpdateUserById(ctx, Id, user)
	if err != nil {
		return "", err
	}
	return updated_at, nil
}

func (s *Service) DeleteUser(ctx context.Context, Id int) (string, error) {
	delete_at := time.Now().Format("2006-01-02 15:04:05")
	deleted_at, err := s.repo.UserI.DeleteUserById(ctx, Id, delete_at)
	if err != nil {
		return "", err
	}
	return deleted_at, nil
}

func (s *Service) GetUser(ctx context.Context, Id int) (*models.GetUser, error) {
	user, err := s.repo.UserI.GetUserById(ctx, Id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) GetUsers(ctx context.Context) ([]*models.GetUser, error) {
	users, err := s.repo.UserI.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Service) GetUsersByFieldCreatedAt(ctx context.Context, created_at string) ([]*models.GetUser, error) {
	users, err := s.repo.UserI.GetUsersByFieldCreatedAt(ctx, created_at)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Service) GetUsersByFieldUpdatedAt(ctx context.Context, updated_at string) ([]*models.GetUser, error) {
	users, err := s.repo.UserI.GetUsersByFieldUpdatedAt(ctx, updated_at)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Service) GetUserFullnameSort(ctx context.Context) ([]*models.GetUser, error) {
	users, err := s.repo.UserI.GetUserFullnameSort(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Service) CreateUsers(ctx context.Context, users []models.UserHandler) ( error) {
	var usersService []models.UserService
	for _, usr := range users {
		var user = models.UserService{
			FullName:  usr.FullName,
			NickName:  usr.NickName,
			Photo:     usr.Photo,
			Birthday:  usr.Birthday,
			Location:  usr.Location,
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
			DeletedAt: "",
		}
		usersService = append(usersService, user)
	}
	 err := s.repo.UserI.CreateUsers(ctx, usersService)
	if err != nil {
		return nil
	}
	return  nil
}