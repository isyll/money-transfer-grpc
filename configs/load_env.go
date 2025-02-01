package configs

import (
	"errors"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() error {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
		return errors.New("")
	}
	return err
}
