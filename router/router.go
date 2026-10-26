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
	// Public route (login, register)
	authGroup := app.Group(authPrefix)
	auth.RegisterAuthRoutes(authGroup)

	// Protected route
	api := app.Group(apiPrefix)
	api.Use(middleware.JWTProtected())

	// semua route di bawah ini wajib pakai token
	user.RegisterUserRoutes(api)
}
