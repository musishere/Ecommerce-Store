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

func SetUpAuth(s string) Auth {
	return Auth{
		Secret: s,
	}
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

func (secret Auth) VerifyToken(authHeader string) (domain.User, error) {
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return domain.User{}, errors.New("invalid authorization header")
	}

	tokenStr := parts[1]

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret.Secret), nil
	})

	if err != nil || !token.Valid {
		return domain.User{}, errors.New("invalid token")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return domain.User{}, errors.New("invalid user_id claim")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return domain.User{}, errors.New("invalid email claim")
	}

	userType, ok := claims["type"].(string)
	if !ok {
		return domain.User{}, errors.New("invalid type claim")
	}

	return domain.User{
		ID:       uint(userID),
		Email:    email,
		UserType: userType,
	}, nil
}

func (secret Auth) Authorize(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")

	if token == "" {
		return fiber.ErrUnauthorized
	}

	user, err := secret.VerifyToken(token)
	if err != nil {
		log.Println(err)
		return fiber.ErrUnauthorized
	}

	// optionally store user in context
	ctx.Locals("user", user)

	return ctx.Next()
}

func (secret Auth) GetCurrentUser(ctx *fiber.Ctx) domain.User {
	user := ctx.Locals("user")

	return user.(domain.User)
}
