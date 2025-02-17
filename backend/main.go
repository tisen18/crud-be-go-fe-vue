package main

import (
	"klinik_crud/config"
	"klinik_crud/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Inisialisasi aplikasi Fiber
	app := fiber.New()

	// Konfigurasi CORS untuk mengizinkan akses dari frontend
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",                            // Mengizinkan akses dari semua origin
		AllowMethods: "GET, POST, PUT, DELETE",       // Mengizinkan method HTTP yang digunakan
		AllowHeaders: "Origin, Content-Type, Accept", // Mengizinkan header yang diperlukan
	}))

	// Middleware CORS default
	app.Use(cors.New())

	// Menginisialisasi koneksi database
	config.ConnectDB()

	// Mendaftarkan semua routes
	routes.SetupRoutes(app)

	// Menjalankan server pada port 9090
	log.Fatal(app.Listen(":9090"))
}
