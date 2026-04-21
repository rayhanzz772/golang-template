package main

import (
	"belajar-coding/go/config"
	"belajar-coding/go/model"
	"belajar-coding/go/router"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	if err := config.InitDB(); err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}
	config.DB.AutoMigrate(&model.User{})
	// seed.SeedUsers()
	app := fiber.New()
	router.SetupRoutes(app)
	log.Fatal(app.Listen(os.Getenv("PORT")))
}
