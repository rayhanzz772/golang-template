package router

import (
	"belajar-coding/go/middleware"
	"belajar-coding/go/modules/auth"
	"belajar-coding/go/modules/user"

	"github.com/gofiber/fiber/v2"
)

const (
	authPrefix = "/auth"
	apiPrefix  = "/api"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/status", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	api := app.Group(apiPrefix)

	// Public route (login, register)
	auth.RegisterAuthRoutes(api)

	// Protected route
	api.Use(middleware.JWTProtected())

	// semua route di bawah ini wajib pakai token
	auth.RegisterProtectedAuthRoutes(api)
	user.RegisterUserRoutes(api)
}
