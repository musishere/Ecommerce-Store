package helper

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/musishere/Mustafa-Ecommerce-App/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Secret string
}

func (secret Auth) CreateHashedPassword(password string) (string, error) {
	if len(password) < 6 {
		return "", errors.New("password must be at least 6 characters")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", errors.New("Password hashing failed")
	}

	return string(hashedPassword), nil
}

func (secret Auth) GenerateToken(id uint, email string, role string) (string, error) {
	if id == 0 || role == "" || email == "" {
		return "", errors.New("Please provide valid user or role")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"role":    role,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret.Secret))
	if err != nil {
		log.Println(err)
		return "", errors.New("Token signing failed")
	}
	return tokenString, nil
}

func (secret Auth) VerifyPassword(password string, hashedPassword string) error {
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Println(err)
		return errors.New("Password does not match")
	}
	return nil
}

func (secret Auth) VerifyToken(token string) (domain.User, error) {
	t := strings.Split(token, "")
	if len(t) != 2 {
		return domain.User{}, errors.New("Invalid token")
	}

	tokenStr := t[0]
	if tokenStr != "Bearer" {
		return domain.User{}, errors.New("Invalid token")
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret.Secret), nil
	})

	if err != nil {
		log.Println(err)
	}
	return domain.User{}, nil
}

func (secret Auth) Authorize(ctx *fiber.Ctx) {}

func (secret Auth) GetCurrentUser(ctx *fiber.Ctx) domain.User {
	return domain.User{}
}
