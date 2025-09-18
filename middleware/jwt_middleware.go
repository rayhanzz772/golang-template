package middleware

import (
	"belajar-coding/go/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing or invalid token",
			})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ParseToken(tokenStr)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		// simpan claims ke context biar bisa diakses di controller
		c.Locals("userID", claims.UserID)
		c.Locals("email", claims.Email)

		return c.Next()
	}
}
