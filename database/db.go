package database

import (
	"database/sql"
	"fmt"
	"log"

	"toystore-api/config"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error

	// Gunakan config dari file .env
	DB, err = sql.Open(config.DBDriver, config.DBName)
	if err != nil {
		log.Fatalf("❌ Gagal konek ke database: %v", err)
	}

	// Tes koneksi
	err = DB.Ping()
	if err != nil {
		log.Fatalf("❌ Database tidak bisa diakses: %v", err)
	}

	fmt.Println("📦 Database terhubung dengan sukses ✔️")

	// Inisialisasi tabel (optional, bisa juga lewat migrasi eksternal)
	initToyTable()
}

func initToyTable() {
	query := `
	CREATE TABLE IF NOT EXISTS toys (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		price INTEGER NOT NULL,
		stock INTEGER NOT NULL
	);`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("❌ Gagal membuat tabel toys: %v", err)
	}
}
