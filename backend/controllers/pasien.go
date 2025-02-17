package controllers

import (
	"klinik_crud/config"
	"klinik_crud/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetAllPasien menangani request untuk mengambil semua data pasien
func GetAllPasien(c *fiber.Ctx) error {
	// Query untuk mengambil semua data pasien dari database
	rows, err := config.DB.Query("SELECT id, nama, alamat, no_hp, penyakit FROM pasien")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	// Slice untuk menyimpan hasil query
	var pasienList []models.Pasien
	// Iterasi setiap baris hasil query
	for rows.Next() {
		var pasien models.Pasien
		// Scan data dari baris ke struct Pasien
		if err := rows.Scan(&pasien.Id, &pasien.Nama, &pasien.Alamat, &pasien.NoHP, &pasien.Penyakit); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		pasienList = append(pasienList, pasien)
	}

	return c.JSON(pasienList)
}

// CreatePasien menangani request untuk membuat data pasien baru
func CreatePasien(c *fiber.Ctx) error {
	var pasien models.Pasien
	// Parse body request ke struct Pasien
	if err := c.BodyParser(&pasien); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Query untuk menyimpan data pasien baru
	_, err := config.DB.Exec("INSERT INTO pasien (nama, alamat, no_hp, penyakit) VALUES (?, ?, ?, ?)",
		pasien.Nama, pasien.Alamat, pasien.NoHP, pasien.Penyakit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Pasien berhasil ditambahkan"})
}

// UpdatePasien menangani request untuk memperbarui data pasien
func UpdatePasien(c *fiber.Ctx) error {
	// Mengambil ID dari parameter URL
	id := c.Params("id")

	// Konversi ID dari string ke integer
	intID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	// Parse body request ke struct Pasien
	var pasien models.Pasien
	if err := c.BodyParser(&pasien); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Query untuk memperbarui data pasien
	result, err := config.DB.Exec("UPDATE pasien SET nama = ?, alamat = ?, no_hp = ?, penyakit = ? WHERE id = ?",
		pasien.Nama, pasien.Alamat, pasien.NoHP, pasien.Penyakit, intID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	// Memeriksa apakah ada baris yang terpengaruh
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Pasien not found"})
	}

	return c.JSON(fiber.Map{"message": "Pasien berhasil diperbarui", "pasien": pasien})
}

// DeletePasien menangani request untuk menghapus data pasien
func DeletePasien(c *fiber.Ctx) error {
	// Mengambil ID dari parameter URL
	id := c.Params("id")

	// Query untuk menghapus data pasien
	_, err := config.DB.Exec("DELETE FROM pasien WHERE id = ?", id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Pasien berhasil dihapus"})
}
