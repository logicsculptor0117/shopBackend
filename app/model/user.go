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

type LoginUser struct {
	Input    string `json:"input" form:"input" validate:"required,min=4,max=32"`
	Password string `json:"password" form:"password" validate:"required,min=4,max=32"`
}

type ReadUser struct {
	Name    string
	Email   string
	Phone   string
	Address string
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (user *User) ComparePassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return false
	}
	return true
}

func (u *User) ReadUser() interface{} {
	return ReadUser{Name: u.Name, Email: u.Email, Phone: u.Phone, Address: u.Address}
}
