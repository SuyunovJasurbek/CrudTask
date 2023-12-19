package service

import (
	"context"
	"time"

	"github.com/SuyunovJasurbek/CrudTask/models"
)

func (s *Service) CreateUser(ctx context.Context, usr models.UserHandler) error {
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

	err := s.repo.UserI.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return  nil
}
