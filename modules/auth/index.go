package auth

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(api fiber.Router) {
	api.Post("/auth/login", LoginHandler)
}

func RegisterProtectedAuthRoutes(api fiber.Router) {
	api.Get("/auth/me", MeHandler)
}
