package main

import (
	"fmt"
	"money-transfer-grpc-user-service/configs"
	"os"
)

func main() {
	if err := configs.LoadEnvVariables(); err != nil {
		fmt.Println("Error loading environment variables:", err)
		os.Exit(1)
	}

	name := os.Getenv("name")
	fmt.Println("Hello " + name)
}
