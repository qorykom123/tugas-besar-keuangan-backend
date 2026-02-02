package handler

import (
	"os"
	"strconv"

	"tugasbesar/config"
	"tugasbesar/model"
	"tugasbesar/repository"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Login User
func Login(c *fiber.Ctx) error {
	var req model.LoginRequest

	// parsing body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Format data salah",
		})
	}

	// cari user berdasarkan username
	user, err := repository.FindUserByUsername(req.Username)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Username atau password salah",
		})
	}

	// compare password (plain vs hash)
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Username atau password salah",
		})
	}

	// ambil expiry JWT dari env
	expMin, _ := strconv.Atoi(os.Getenv("JWT_EXPIRES_MINUTES"))

	// generate token
	token, err := config.GenerateToken(
		user.ID,
		user.Username,
		user.Role,
		expMin,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal membuat token",
			"error":   err.Error(),
		})
	}

	// response
	return c.JSON(fiber.Map{
		"message": "Login berhasil",
		"token":   token,
		"user": fiber.Map{
			"username": user.Username,
			"role":     user.Role,
		},
	})
}
