package auth

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(api fiber.Router) {
	api.Post("/login", LoginHandler)
}
