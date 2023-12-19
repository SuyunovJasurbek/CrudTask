package service

import "github.com/SuyunovJasurbek/CrudTask/src/repository"

type User interface {
	CreateUser (string ) int 
}

type Service struct {
	User 
}

func NewService( rep *repository.Repository) *Service {
	return &Service{User: NewUserService(rep.UserI)}
}