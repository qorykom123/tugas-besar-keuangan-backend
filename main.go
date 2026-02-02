package main

import (
	"tugasbesar/config"
	"tugasbesar/model"
	"tugasbesar/router"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// ✅ LOAD .env
	if err := godotenv.Load(); err != nil {
		log.Println(".env tidak ditemukan, pakai system env")
	}

	app := fiber.New()

	// Inisialisasi koneksi DB
	config.InitDB()

	// Auto migrate tabel
	config.DB.AutoMigrate(&model.User{})

	// CORS
	config.SetupCORS(app)

	// Logger
	app.Use(logger.New())

	// Routes
	router.SetupRoutes(app)

	// Run server
	app.Listen(":3000")
}
