package controller

import (
	"fmt"
	"shopBackend/app/model"
	"shopBackend/app/service"
	"shopBackend/helpers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		var user model.User
		// check data type
		if err := helpers.DataContentType(ctx, &user); err != nil {
			helpers.RespondJSON(ctx, 400, helpers.StatusCodeFromInt(400), err.Error(), nil)
			return
		}
		// check validate field
		if err := validator.New().Struct(&user); err != nil {
			listErrors := helpers.ValidateErrors(err.(validator.ValidationErrors))
			helpers.RespondJSON(ctx, 400, helpers.StatusCodeFromInt(400), listErrors, nil)
			return
		}
		// Create
		if err := uc.service.Register(&user); err != nil {
			statusCode, message := helpers.DBError(err)
			helpers.RespondJSON(ctx, statusCode, helpers.StatusCodeFromInt(statusCode), message, nil)
			return
		}
		helpers.RespondJSON(ctx, 201, helpers.StatusCodeFromInt(201), nil, nil)
		return
	}
}
