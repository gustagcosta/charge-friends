package shared

import (
	"backend/config"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Protected(ctx *fiber.Ctx) error {
	env, err := config.LoadConfig()
	if err != nil {
		fmt.Println("error: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed at register user",
		})
	}

	var tokenString string
	authorization := ctx.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	}

	if tokenString == "" {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte(env.JWT_KEY), nil
	})

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "invalidate token"})
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "invalid token claim"})
	}

	ctx.Locals("userId", claims["id"].(float64))
	return ctx.Next()
}
