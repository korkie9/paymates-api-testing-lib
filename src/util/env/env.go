package util

import (
	"github.com/joho/godotenv"
	"os"
	"paymates-mock-db-updater/src/check_error"
)

func DotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	check_error.ErrCheck(err)
	return os.Getenv(key)
}
