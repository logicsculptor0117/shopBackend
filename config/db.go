package config

import (
	"fmt"
	"os"
	"shopBackend/app/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("POSTGRES_DB"))
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Role{})
	fmt.Println("Migration complete")
	DB = db

	//create role table
	if errRole := InitRole(); errRole != nil {
		fmt.Println("Dont create Role")
	}
}

func InitRole() error {
	var adminRole model.Role
	adminRole.Name = "admin"
	if err := DB.Create(&adminRole).Error; err != nil {
		fmt.Println("Error Database: Dont create admin role")
	}

	var userRole model.Role
	userRole.Name = "user"
	if err := DB.Create(&userRole).Error; err != nil {
		fmt.Println("Error Database: Dont create user role")
	}

	return nil
}
