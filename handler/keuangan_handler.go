package handler

import (
	"tugasbesar/model"
	"tugasbesar/repository"

	"github.com/gofiber/fiber/v2"
)

// Get All Transaksi Keuangan
func GetAllKeuangan(c *fiber.Ctx) error {
	data, err := repository.GetAllKeuangan()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal mengambil data transaksi",
			"error":   err.Error(),
		})
	}

	// Rapihin tanggal
	for i := range data {
		if len(data[i].Tanggal) >= 10 {
			data[i].Tanggal = data[i].Tanggal[:10]
		}
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil mengambil data transaksi keuangan",
		"data":    data,
	})
}

// Insert Transaksi Keuangan
func InsertKeuangan(c *fiber.Ctx) error {
	var trx model.Transaction

	if err := c.BodyParser(&trx); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "BodyParser error",
			"error":   err.Error(),
			"raw":     string(c.Body()),
		})
	}

	if err := repository.InsertKeuangan(&trx); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal insert",
			"error":   err.Error(),
		})
	}

	return c.JSON(trx)
}

// Get satu transaksi keuangan berdasarkan ID
func GetKeuanganByID(c *fiber.Ctx) error {
	id := c.Params("id")

	data, err := repository.GetKeuanganByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Transaksi tidak ditemukan",
			"error":   err.Error(),
		})
	}

	// Rapihin tanggal (YYYY-MM-DD)
	if len(data.Tanggal) >= 10 {
		data.Tanggal = data.Tanggal[:10]
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil mengambil data transaksi keuangan",
		"data":    data,
	})
}

// Update data transaksi keuangan berdasarkan ID
func UpdateKeuangan(c *fiber.Ctx) error {
	id := c.Params("id")

	var input model.Transaction
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Format data salah",
		})
	}

	updated, err := repository.UpdateKeuangan(id, input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal mengupdate data transaksi",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data transaksi berhasil diupdate",
		"data":    updated,
	})
}

// Delete transaksi keuangan berdasarkan ID
func DeleteKeuangan(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := repository.DeleteKeuangan(id); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal menghapus data transaksi",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Transaksi berhasil dihapus",
	})
}
