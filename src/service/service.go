package service

import "github.com/SuyunovJasurbek/CrudTask/src/repository"


type Service struct {
	repo repository.Repository
}

func NewService( repo *repository.Repository) *Service {
	return &Service{repo: *repo}
}