package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/musishere/Mustafa-Ecommerce-App/configs"
	"github.com/musishere/Mustafa-Ecommerce-App/internal/api/rest"
	"github.com/musishere/Mustafa-Ecommerce-App/internal/api/rest/handlers"
)

func StartServer(appconfig configs.Appconfig) {
	app := fiber.New()
	app.Get("/health", HealthCheck)

	rh := &rest.RestHandler{
		App: app,
	}

	setUpRoutes(rh)
	app.Listen(appconfig.ServerPort)
}

func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "Healthy!"})
}

func setUpRoutes(restHandlers *rest.RestHandler) {
	handlers.SetUpUserRoutes(restHandlers)
}
