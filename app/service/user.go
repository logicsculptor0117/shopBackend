package service

import (
	"errors"
	"shopBackend/app/model"
	"shopBackend/app/repository"
)

type UserService struct {
	repo repository.UserRepoInterface
}

type UserServiceInterface interface {
	Register(user *model.User) error
}

func NewUserService(repo repository.UserRepoInterface) *UserService {
	return &UserService{repo}
}

func (s *UserService) Register(user *model.User) error {
	roleName := "user"
	role, err := s.repo.FindRoleByName(roleName)
	if err != nil {
		return errors.New("dont find role name")
	}
	user.RoleId = role.Id
	if errHashPass := user.HashPassword(); errHashPass != nil {
		return errors.New("cant hash password")
	}
	return s.repo.Create(user)
}
