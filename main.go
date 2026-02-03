package main

import (
	"log"
	"os"

	"tugasbesar/config"
	"tugasbesar/model"
	"tugasbesar/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Load env (local only, Railway pakai system env)
	_ = godotenv.Load()

	app := fiber.New()

	// DB
	config.InitDB()
	config.DB.AutoMigrate(&model.User{})

	// CORS
	config.SetupCORS(app)

	// Logger
	app.Use(logger.New())

	// Routes
	router.SetupRoutes(app)

	// Railway PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // local
	}

	log.Println("Server running on port:", port)
	log.Fatal(app.Listen(":" + port))
}
