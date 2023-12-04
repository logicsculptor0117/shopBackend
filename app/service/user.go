package service

import (
	"errors"
	"shopBackend/app/model"
	"shopBackend/app/repository"
	middleware "shopBackend/moddleware"
)

type UserService struct {
	repo repository.UserRepoInterface
}

type UserServiceInterface interface {
	Register(user *model.User) error
	Login(loginUser *model.LoginUser) (error, string)
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

func (us *UserService) Login(loginUser *model.LoginUser) (error, string) {
	user, err := us.repo.FindUser(loginUser.Input)
	if err != nil {
		return errors.New("name or email isn't already exist"), ""
	}
	if user.ComparePassword(loginUser.Password) == false {
		return errors.New("incorrect password"), ""
	}
	// create token
	token, errCreate := middleware.CreateToken(user.Id)
	if errCreate != nil {
		return errCreate, ""
	}
	return nil, token
}
