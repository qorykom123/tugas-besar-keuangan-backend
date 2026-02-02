package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	_ = godotenv.Load()
	dsn := os.Getenv("SUPABASE_DSN")
	if dsn == "" {
		log.Fatal("SUPABASE_DSN tidak ditemukan. Pastikan .env berisi DSN Supabase.")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal konek ke database: %v", err)
	}
	DB = db
	fmt.Println("✅ Koneksi ke PostgreSQL (Supabase) BERHASIL")
}
func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("DB belum diinisialisasi. Panggil config.InitDB() lebih dulu.")
	}
	return DB
}
