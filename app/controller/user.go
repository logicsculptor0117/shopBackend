package controller

import (
	"fmt"
	"net/http"
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

func (uc *UserController) LoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginUser model.LoginUser
		// check data type
		if err := helpers.DataContentType(ctx, &loginUser); err != nil {
			helpers.RespondJSON(ctx, 400, helpers.StatusCodeFromInt(400), err.Error(), nil)
			return
		}
		// check validate field
		if err := validator.New().Struct(&loginUser); err != nil {
			listErrors := helpers.ValidateErrors(err.(validator.ValidationErrors))
			helpers.RespondJSON(ctx, 400, helpers.StatusCodeFromInt(400), listErrors, nil)
			return
		}
		// login
		err, token := uc.service.Login(&loginUser)
		if err != nil {
			statusCode, message := helpers.DBError(err)
			helpers.RespondJSON(ctx, statusCode, helpers.StatusCodeFromInt(statusCode), message, nil)
			return
		}
		// response token
		ctx.SetSameSite(http.SameSiteLaxMode)
		ctx.SetCookie("Authorization", token, 3600*12, "/", "", false, false)
		helpers.RespondJSON(ctx, 201, helpers.StatusCodeFromInt(201), nil, token)
		return
	}
}

func (uc *UserController) ReadUserHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, isUser := ctx.MustGet("user").(*model.User)
		if !isUser {
			helpers.RespondJSON(ctx, 500, helpers.StatusCodeFromInt(500), "No provide token", nil)
			return
		}
		helpers.RespondJSON(ctx, 200, helpers.StatusCodeFromInt(200), nil, user.ReadUser())
		return
	}
}
