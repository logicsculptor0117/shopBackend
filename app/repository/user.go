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
