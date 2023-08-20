package helper

import (
	"errors"
	"fmt"
	"os"

	e "root/check_error"

	"github.com/joho/godotenv"
)

func init() {
	loadEnv()
}

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		e.ErrorHandler(errors.New("error loading .env file"))
	}
}

func Port() string {
	return fmt.Sprintf(":%s", os.Getenv("PORT"))
}
