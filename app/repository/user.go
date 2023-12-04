package repository

import (
	"shopBackend/app/model"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

type UserRepoInterface interface {
	FindRoleByName(roleName string) (*model.Role, error)
	Create(user *model.User) error
	FindUser(input string) (*model.User, error)
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

func (ur *UserRepo) FindRoleByName(roleName string) (*model.Role, error) {
	var role *model.Role
	if err := ur.db.Where("name = ?", roleName).First(&role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (ur *UserRepo) Create(user *model.User) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepo) FindUser(input string) (*model.User, error) {
	var user *model.User
	if err := ur.db.Where("name=? OR email=?", input, input).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
