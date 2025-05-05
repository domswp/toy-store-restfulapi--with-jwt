package config

import (
	"log"
	"os"
)

// Default fallback kalau .env tidak ditemukan
var (
	DefaultPort  = "8080"
	JWTSecretKey = "defaultSecretKey"
	DBDriver     = "sqlite3"
	DBName       = "toystore.db"
)

func InitConfig() {
	// Cek dan update dari ENV kalau tersedia
	if port := os.Getenv("PORT"); port != "" {
		DefaultPort = port
	}
	if secret := os.Getenv("JWT_SECRET"); secret != "" {
		JWTSecretKey = secret
	}
	if driver := os.Getenv("DB_DRIVER"); driver != "" {
		DBDriver = driver
	}
	if dbname := os.Getenv("DB_NAME"); dbname != "" {
		DBName = dbname
	}

	log.Println("✅ Konfigurasi berhasil dimuat dari .env")
}
