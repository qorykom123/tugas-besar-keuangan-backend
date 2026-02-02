package config

import (
	"os"
	"strings"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// JWTMiddleware: pasang untuk route yang butuh auth
func JWTMiddleware() fiber.Handler {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default_secret_change_me"
	}

	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(secret)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			msg := strings.ToLower(err.Error())
			// 1) Token tidak ada / format salah
			if strings.Contains(msg, "missing") || strings.Contains(msg,

				"malformed") {

				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Token tidak ada / format salah",
					"hint":    "Gunakan header: Authorization: Bearer <token>",
				})
			}
			// 2) Token invalid / expired
			if strings.Contains(msg, "expired") || strings.Contains(msg,

				"invalid") {

				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Token tidak valid atau sudah kedaluwarsa",
					"hint":    "Silakan login ulang untuk mendapatkan token baru",
				})
			}
			// 3) Fallback (error lain)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
				"error":   err.Error(),
			})
		},
	})
}

// GenerateToken: buat token saat login berhasil
func GenerateToken(userID, username string, role string, expiresMinutes int) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default_secret_change_me"
	}
	if expiresMinutes <= 0 {
		expiresMinutes = 60

	}
	claims := jwt.MapClaims{
		"sub":      userID,
		"username": username,
		"role":     role,
		"exp": time.Now().Add(time.Duration(expiresMinutes) *
			time.Minute).Unix(),
		"iat": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// Helper untuk ambil claims dari ctx (opsional, kalau butuh role check, dll)
func GetClaims(c *fiber.Ctx) (jwt.MapClaims, bool) {
	user := c.Locals("user")
	if user == nil {
		return nil, false
	}
	tok, ok := user.(*jwt.Token)
	if !ok {
		return nil, false
	}
	claims, ok := tok.Claims.(jwt.MapClaims)
	return claims, ok
}
