package routes

import (
	"klinik_crud/controllers"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes mendaftarkan semua endpoint API untuk manajemen pasien
func SetupRoutes(app *fiber.App) {
	// Membuat grup route dengan prefix /api
	api := app.Group("/api")

	// Mendaftarkan endpoint-endpoint untuk CRUD pasien
	api.Get("/pasien", controllers.GetAllPasien)        // Endpoint untuk mengambil semua data pasien
	api.Post("/pasien", controllers.CreatePasien)       // Endpoint untuk membuat pasien baru
	api.Put("/pasien/:id", controllers.UpdatePasien)    // Endpoint untuk memperbarui data pasien
	api.Delete("/pasien/:id", controllers.DeletePasien) // Endpoint untuk menghapus data pasien
}
