package service

import (
	"shopBackend/app/model"
	"shopBackend/app/repository"
)

type UserService struct {
	repo repository.UserRepoInterface
}

type UserServiceInterface interface {
	Regist(user *model.User) error
}

func NewUserService(repo repository.UserRepoInterface) *UserService {
	return &UserService{repo}
}

func (s *UserService) Regist(user *model.User) error {
	return nil
}
