package controller

import (
	"fmt"
	"shopBackend/app/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.UserServiceInterface
}

func NewUserController(service service.UserServiceInterface) *UserController {
	return &UserController{service}
}

func (uc *UserController) RegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("POST /user/register of UserController started")
		return
	}
}
