package service

import "github.com/SuyunovJasurbek/CrudTask/src/repository"

type UserService struct {
	user_repo repository.UserI
}

func NewUserService(repo repository.UserI) *UserService {
	return &UserService{
		user_repo: repo,
	}
}
func (s *UserService)CreateUser (string ) int  {
	panic("s")
}