package service

import "context"


func (s *Service)CreateUser() int  {
	return s.repo.UserI.CreateUser(context.Background())
}