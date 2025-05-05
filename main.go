package main

import (
	"log"
	"net/http"

	// Import package lokal (buatan sendiri)
	"toystore-api/config"   // untuk konfigurasi global dari file .env
	"toystore-api/database" // untuk koneksi ke database
	"toystore-api/routes"   // untuk routing RESTful API

	// Import eksternal package
	"github.com/joho/godotenv" // untuk load file .env
)

func main() {
	// ✅ Load variabel dari file .env ke environment Go
	// Kalau tidak ada .env, akan pakai nilai default di config/config.go
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  .env file not found, using default config")
	}

	// ✅ Inisialisasi konfigurasi dari environment (PORT, JWT_SECRET, dll)
	config.InitConfig()

	// ✅ Buka koneksi ke database (misalnya: SQLite, MySQL, dsb)
	database.InitDB()

	// ✅ Setup routing untuk semua endpoint API
	r := routes.SetupRouter()

	// ✅ Jalankan server di port yang diambil dari config
	log.Printf("🚀 Server running on http://localhost:%s", config.DefaultPort)
	log.Fatal(http.ListenAndServe(":"+config.DefaultPort, r))
}
