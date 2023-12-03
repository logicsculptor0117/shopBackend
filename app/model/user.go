package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint   `json:"ID"       form:"ID"       gorm:"primary_key"`
	Name      string `json:"name"     form:"name"     gorm:"unique;not null"   validate:"required,min=4,max=32"`
	Email     string `json:"email"    form:"email"    gorm:"unique"            validate:"required,email"`
	Phone     string `json:"phone"    form:"phone"    gorm:"unique"            validate:"required,len=10"`
	Password  string `json:"password" form:"password" gorm:"not null"          validate:"required,min"`
	Address   string `json:"address"  form:"address"  gorm:""                  validate:""`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
