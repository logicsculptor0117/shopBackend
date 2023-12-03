package repository

import (
	"shopBackend/app/model"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

type UserRepoInterface interface {
	Create(user *model.User) error
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

func (ur *UserRepo) Create(user *model.User) error {
	return ur.db.Create(user).Error
}
