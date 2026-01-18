package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/musishere/Mustafa-Ecommerce-App/configs"
)

func StartServer(appconfig configs.Appconfig) {
	app := fiber.New()
	app.Get("/health", HealthCheck)
	app.Listen(appconfig.ServerPort)
}

func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "Healthy!"})
}
