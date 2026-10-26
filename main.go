package main

import (
	"belajar-coding/go/config"
	"belajar-coding/go/model"
	seed "belajar-coding/go/seeder"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"belajar-coding/go/router"

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

	command := "server"
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	switch command {
	case "server", "serve":
		runServer()
	case "seed":
		seed.SeedUsers()
		fmt.Println("Seed completed")
	case "migrate":
		if err := config.DB.AutoMigrate(&model.User{}); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
		fmt.Println("Migration completed")
	case "help", "-h", "--help":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", command)
		printUsage()
		os.Exit(1)
	}
}

func runServer() {
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

func printUsage() {
	fmt.Print(`Usage: app <command>

Commands:
  server   Start the HTTP server (default)
  seed     Seed the database with sample data
  migrate  Run database auto-migration
  help     Show this help message

Examples:
  app server
  app seed
  app migrate
`)
}
