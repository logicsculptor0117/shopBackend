package main

import (
	"shopBackend/app/router"
	"shopBackend/config"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
}

func main() {
	r := gin.Default()

	// Routers
	router.UserRouter(r)

	// Start the server
	if err := r.Run(); err != nil {
		panic(err)
	}

}
