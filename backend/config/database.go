package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Variable global untuk koneksi database
var DB *sql.DB

// ConnectDB menginisialisasi koneksi ke database MySQL
func ConnectDB() {
	// Memuat konfigurasi dari file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Mengambil konfigurasi database dari environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Membuat string koneksi database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Membuka koneksi ke database
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	// Memverifikasi koneksi database
	if err = DB.Ping(); err != nil {
		log.Fatal("Database ping failed: ", err)
	}

	log.Println("Connected to database successfully")
}
