package main

import (
	"fmt"

	"shopBackend/app/config"
)

func init() {
	config.LoadEnv()
	config.ConnectDB()
}

func main() {
	fmt.Println("server started")
}
