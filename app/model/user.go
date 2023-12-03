package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id        uint   `json:"ID"       form:"ID"       gorm:"primary_key"`
	Name      string `json:"name"     form:"name"     gorm:"unique;not null"   validate:"required,min=4,max=32"`
	Email     string `json:"email"    form:"email"    gorm:"unique"            validate:"required,email"`
	Phone     string `json:"phone"    form:"phone"    gorm:"unique"            validate:"required,len=10"`
	Password  string `json:"password" form:"password" gorm:"not null"          validate:"required,min=4,max=32"`
	Address   string `json:"address"  form:"address"  gorm:""                  validate:""`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	RoleId uint
}

type Role struct {
	Id        uint   `json:"ID" form:"ID" gorm:"primary_key"`
	Name      string `json:"name" form:"name" gorm:"unique;not null" validate:"required,min=4,max=32"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	User User `gorm:"foreignKey:RoleId;references:Id"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
