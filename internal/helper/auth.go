package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/musishere/Mustafa-Ecommerce-App/internal/domain"
)

type Auth struct {
	Secret string
}

func (secret Auth) CreateHashedPassword(password string) (string, error) {

	return "", nil
}

func (secret Auth) GenerateToken(id uint, email string, password string) (string, error) {
	return "", nil
}

func (secret Auth) VerifyPassword(password string, hashedPassword string) error {

	return nil
}

func (secret Auth) VerifyToken(token string) (domain.User, error) {
	return domain.User{}, nil
}

func (secret Auth) Authorize(ctx *fiber.Ctx) {}

func (secret Auth) GetCurrentUser(ctx *fiber.Ctx) domain.User {
	return domain.User{}
}
