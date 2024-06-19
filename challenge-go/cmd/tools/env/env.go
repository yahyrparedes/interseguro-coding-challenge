package env

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"os"
)

const (
	localEnv = "local"
	devEnv   = "dev"
	testEnv  = "test"
	prodEnv  = "prod"
)

func NewEnv() *Env {
	initializeEnv()
	return NewDataEnv()
}

func initializeEnv() {
	if IsLocal() {
		err := godotenv.Load(".env.local")
		if err != nil {
			log.Fatal("Error load environment local: " + err.Error())
			//panic("Environment variable not defined: " + err.Error())
		}
	}
}

func GetSecureEnv(key string) string {
	var data = os.Getenv(key)
	if data == "" {
		panic("Environment variable not defined: " + key)
	}
	return data
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

func IsLocal() bool {
	return os.Getenv(AppEnv) == localEnv || os.Getenv("APP_ENV") == ""
}

func IsDev() bool {
	return os.Getenv(AppEnv) == devEnv
}

func IsTest() bool {
	return os.Getenv(AppEnv) == testEnv
}

func IsProd() bool {
	return os.Getenv(AppEnv) == prodEnv
}
