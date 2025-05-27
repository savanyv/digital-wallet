package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/savanyv/digital-wallet/shared/utils/jwt"
)

func AuthMiddlewares(jwtService jwt.JWTService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthorized",
			})
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token format",
			})
		}

		token := bearerToken[1]
		claims, err := jwtService.ValidateToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		c.Locals("user_id", claims.UserID)

		return c.Next()
	}
}
