package router

import (
	"tugasbesar/config"
	"tugasbesar/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	/* ===============================
	   PUBLIC ROUTES
	=============================== */
	api.Post("/login", handler.Login)
	api.Post("/register", handler.CreateUser)

	/* ===============================
	   PROTECTED ROUTES (JWT)
	=============================== */
	protected := api.Group("", config.JWTMiddleware())

	// ===== KEUANGAN =====

	// semua role boleh lihat
	protected.Get("/keuangan", handler.GetAllKeuangan)
	protected.Get("/keuangan/:id", handler.GetKeuanganByID)

	// cuma admin boleh ubah data
	protected.Post("/keuangan",
		config.RequireRole("admin"),
		handler.InsertKeuangan,
	)

	protected.Put("/keuangan/:id",
		config.RequireRole("admin"),
		handler.UpdateKeuangan,
	)

	protected.Delete("/keuangan/:id",
		config.RequireRole("admin"),
		handler.DeleteKeuangan,
	)
}
