package main

import (
	"log"

	"github.com/musishere/Mustafa-Ecommerce-App/configs"
	"github.com/musishere/Mustafa-Ecommerce-App/internal/api"
)

func main() {
	cfg, err := configs.SetUpEnvironmentVariables()

	if err != nil {
		log.Fatalf("Error setting up environment variables: %v", err)
	}
	api.StartServer(cfg)

}
