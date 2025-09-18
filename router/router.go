package router

import (
	"belajar-coding/go/config"
	"belajar-coding/go/modules/auth"
	"belajar-coding/go/modules/user"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

const (
	authPrefix = "/auth"
	apiPrefix  = "/api"
)

func SetupRoutes(app *fiber.App) {
	// Public route (login, register)
	authGroup := app.Group(authPrefix)
	auth.RegisterAuthRoutes(authGroup)

	// Protected route
	api := app.Group(apiPrefix)

	// pasang middleware JWT
	api.Use(jwtware.New(jwtware.Config{
		SigningKey:    []byte(config.App.JWT.Secret),
		ErrorHandler:  jwtError,
		SigningMethod: "HS256",                // Specify signing method
		TokenLookup:   "header:Authorization", // Define token location
		AuthScheme:    "Bearer",               // Specify auth scheme
	}))

	// semua route di bawah ini wajib pakai token
	user.RegisterUserRoutes(api)
}

func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"success": false,
		"message": "Unauthorized or invalid token",
		"error":   err.Error(), // Add specific error for debugging
		"code":    "AUTH_ERROR",
	})
}
