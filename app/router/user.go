package router

import (
	"fmt"
	"shopBackend/app/controller"
	"shopBackend/app/repository"
	"shopBackend/app/service"
	"shopBackend/config"
	middleware "shopBackend/moddleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	userRepo := repository.NewUserRepo(config.DB)
	userService := service.NewUserService(userRepo)
	userMiddleware := middleware.NewUserMiddleware(userRepo)
	userController := controller.NewUserController(userService)

	router.POST("/user/register", userController.RegisterHandler())
	router.POST("/user/login", userController.LoginHandler())

	router.GET("/user", userMiddleware.AuthMiddleware(), userController.ReadUserHandler())

	fmt.Println(userService, userMiddleware)
}
