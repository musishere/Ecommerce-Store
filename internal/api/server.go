package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/musishere/Mustafa-Ecommerce-App/configs"
)

func StartServer(appconfig configs.Appconfig) {
	app := fiber.New()

	app.Listen(appconfig.ServerPort)
}
