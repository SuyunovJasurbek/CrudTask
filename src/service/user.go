package service

import (
	"context"
	"time"

	"github.com/SuyunovJasurbek/CrudTask/models"
)

func (s *Service) CreateUser(ctx context.Context, usr models.UserHandler) (int ,  error) {
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

	id ,err := s.repo.UserI.CreateUser(ctx, user)
	if err != nil {
		return 0 , err
	}
	return  id , nil
}

func (s *Service) UpdateUser(ctx context.Context, Id int , usr models.UserHandler) (string ,error) {
	var user = models.UserUpdate{
		FullName:  usr.FullName,
		NickName:  usr.NickName,
		Photo:     usr.Photo,
		Birthday:  usr.Birthday,
		Location:  usr.Location,
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	updated_at, err := s.repo.UserI.UpdateUser(ctx, Id , user)
	if err != nil {
		return "" , err
	}
	return  updated_at , nil
}