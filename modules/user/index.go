package user

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(api fiber.Router) {
	api.Post("/users", create)
	api.Get("/users", list)
	api.Get("/users/:id", show)
	api.Put("/users/:id", update)
	api.Delete("/users/:id", delete)
}
