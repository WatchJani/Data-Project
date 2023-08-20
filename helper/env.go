package helper

import (
	"errors"

	e "root/check_error"

	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load(".env.local"); err != nil {
		e.ErrorHandler(errors.New("error loading .env file"))
	}
}
