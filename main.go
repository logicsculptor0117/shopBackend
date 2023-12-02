package main

import (
	"fmt"

	"shopBackend/app/config"
)

func init() {
	config.LoadEnv()
}

func main() {
	fmt.Println("server started")
}
