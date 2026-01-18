package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/musishere/Mustafa-Ecommerce-App/internal/api/rest"
	"github.com/musishere/Mustafa-Ecommerce-App/internal/service"
)

type UserHandler struct {
	svc service.UserService
}

func SetUpUserRoutes(restHandler *rest.RestHandler) {
	app := restHandler.App
	svc := service.UserService{}
	handler := UserHandler{
		svc: svc,
	}
	//	public endpoints
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	//private endpoints
	app.Get("/verify", handler.GetVerificationCode)
	app.Post("/verify", handler.Verify)
	app.Post("/profile", handler.CreateProfile)
	app.Get("/profile", handler.GetProfile)
	app.Post("/cart", handler.AddToCart)
	app.Get("/cart", handler.GetCart)
	app.Get("/order", handler.GetOrders)
	app.Get("/order/:id", handler.GetOrder)
	app.Post("/become-seller", handler.BecomeSeller)

}

func (handler *UserHandler) Register(ctx *fiber.Ctx) error {

	handler.svc.Register()
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "User Registered!"})
}
func (handler *UserHandler) Login(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "User Login!"})
}
func (handler *UserHandler) Verify(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "Verify!"})
}
func (handler *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "Verification Code!"})
}
func (handler *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "get profile!"})
}
func (handler *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "create profile!"})
}
func (handler *UserHandler) AddToCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "add to cart!"})
}
func (handler *UserHandler) GetCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "get cart!"})
}
func (handler *UserHandler) CreateOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "Create order!"})
}
func (handler *UserHandler) GetOrders(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "GetOrders!"})
}
func (handler *UserHandler) GetOrder(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "Getorder!"})
}
func (handler *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "Seller!"})
}
