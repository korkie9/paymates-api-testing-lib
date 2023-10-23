package env

import (
	"github.com/joho/godotenv"
	"os"
	"paymates-mock-db-updater/checkError"
)

func DotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	checkError.ErrCheck(err)
	return os.Getenv(key)
}
