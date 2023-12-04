package repository

import (
	"shopBackend/app/model"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

type UserRepoInterface interface {
	FindRoleByName(roleName string) (*model.Role, error)
	Create(user *model.User) error
	FindUser(input string) (*model.User, error)
	GetUserFromToken(*jwt.Token) (*model.User, error)
	Update(user *model.User) error
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

func (ur *UserRepo) GetUserFromToken(token *jwt.Token) (*model.User, error) {
	claims := token.Claims.(jwt.MapClaims)
	// query user from id
	var user *model.User
	if err := ur.db.Where("id=?", claims["id"]).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepo) Update(user *model.User) error {
	return ur.db.Save(user).Error
}
