package handler

import (
	"tugasbesar/model"
	"tugasbesar/repository"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser / Register User
func CreateUser(c *fiber.Ctx) error {
	var req model.CreateUserRequest

	// parsing body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Format data salah",
		})
	}

	// validasi sederhana
	if req.Username == "" || req.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Username dan password wajib diisi",
		})
	}

	// cek username sudah dipakai atau belum
	if _, err := repository.FindUserByUsername(req.Username); err == nil {
		return c.Status(409).JSON(fiber.Map{
			"message": "Username sudah digunakan",
		})
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal hash password",
			"error":   err.Error(),
		})
	}

	// buat user baru
	role := "user" // default aman

	if req.Role == "admin" {
		role = "admin"
	}

	newUser := model.User{
		Username: req.Username,
		Password: string(hash),
		Role:     role,
	}

	created, err := repository.CreateUser(newUser)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal membuat user",
			"error":   err.Error(),
		})
	}

	// response (tanpa password)
	return c.Status(201).JSON(fiber.Map{
		"message": "User berhasil dibuat",
		"data": fiber.Map{
			"id":       created.ID,
			"username": created.Username,
			"role":     created.Role,
		},
	})
}
