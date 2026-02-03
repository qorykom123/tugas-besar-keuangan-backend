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
	// Load env (local only, Render/Railway pakai system env)
	_ = godotenv.Load()

	app := fiber.New()

	// DB
	config.InitDB()
	config.DB.AutoMigrate(&model.User{})

	// CORS
	config.SetupCORS(app)

	// Logger
	app.Use(logger.New())

	// Root health check
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Backend Tugas Besar Keuangan running 🚀",
		})
	})

	// Routes
	router.SetupRoutes(app)

	// PORT (Render/Railway wajib pakai env PORT)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("Server running on port:", port)
	log.Fatal(app.Listen(":" + port))
}
