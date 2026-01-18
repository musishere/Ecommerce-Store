package configs

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Appconfig struct {
	ServerPort string
}

func SetUpEnvironmentVariables() (cfg Appconfig, err error) {
	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}
	httpPort := os.Getenv("HTTP_PORT")

	if len(httpPort) < 1 {
		return Appconfig{}, errors.New("environment variable HTTP_PORT is not set")
	}

	return Appconfig{httpPort}, nil
}
