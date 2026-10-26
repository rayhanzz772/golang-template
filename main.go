package main

import (
	"belajar-coding/go/config"
	"belajar-coding/go/model"
	"belajar-coding/go/router"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	config.LoadConfig()

	if len(config.App.JWT.Secret) == 0 {
		log.Fatal("JWT_SECRET is required")
	}

	if err := config.InitDB(); err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}
	if err := config.DB.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	app := fiber.New()
	router.SetupRoutes(app)

	go func() {
		addr := fmt.Sprintf("%s:%s", config.App.Server.Host, config.App.Server.Port)
		if err := app.Listen(addr); err != nil && !errors.Is(err, fiber.ErrServiceUnavailable) {
			log.Fatalf("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	if err := app.Shutdown(); err != nil {
		log.Fatalf("Graceful shutdown failed: %v", err)
	}
	log.Println("Server stopped")
}
