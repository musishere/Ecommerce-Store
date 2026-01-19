package api

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/musishere/Mustafa-Ecommerce-App/configs"
	"github.com/musishere/Mustafa-Ecommerce-App/internal/api/rest"
	"github.com/musishere/Mustafa-Ecommerce-App/internal/api/rest/handlers"
	"github.com/musishere/Mustafa-Ecommerce-App/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(appconfig configs.Appconfig) {
	app := fiber.New()
	app.Get("/health", HealthCheck)

	db, err := gorm.Open(postgres.Open(appconfig.Dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	log.Printf("Database Connected")
	//run migration
	db.AutoMigrate(&domain.User{})
	rh := &rest.RestHandler{
		App: app,
		DB:  db,
	}

	setUpRoutes(rh)
	app.Listen(appconfig.ServerPort)
}

func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "Healthy!"})
}

func setUpRoutes(restHandlers *rest.RestHandler) {
	//user handler
	handlers.SetUpUserRoutes(restHandlers)

	//	transaction handler

	// catalog handler
}
