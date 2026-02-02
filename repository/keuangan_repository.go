package repository

import (
	"tugasbesar/config"
	"tugasbesar/model"
)

// Ambil semua data transaksi keuangan
func GetAllKeuangan() ([]model.Transaction, error) {
	var data []model.Transaction
	err := config.GetDB().Find(&data).Error
	return data, err
}

// Insert transaksi keuangan baru
func InsertKeuangan(trx *model.Transaction) error {
	return config.GetDB().Create(trx).Error
}

// Ambil satu data transaksi keuangan berdasarkan ID
func GetKeuanganByID(id string) (model.Transaction, error) {
	var trx model.Transaction
	err := config.GetDB().First(&trx, "id = ?", id).Error
	return trx, err
}

// Update data transaksi keuangan berdasarkan ID
func UpdateKeuangan(id string, newData model.Transaction) (model.Transaction, error) {
	db := config.GetDB()
	var trx model.Transaction

	// Ambil data lama
	if err := db.First(&trx, "id = ?", id).Error; err != nil {
		return trx, err
	}

	// Update HANYA field yang dikirim
	if newData.Tanggal != "" {
		trx.Tanggal = newData.Tanggal
	}
	if newData.Jenis != "" {
		trx.Jenis = newData.Jenis
	}
	if newData.Kategori != "" {
		trx.Kategori = newData.Kategori
	}
	if newData.Deskripsi != "" {
		trx.Deskripsi = newData.Deskripsi
	}
	if newData.Jumlah != 0 {
		trx.Jumlah = newData.Jumlah
	}

	// Simpan perubahan
	if err := db.Save(&trx).Error; err != nil {
		return trx, err
	}

	return trx, nil
}

// Hapus data transaksi keuangan berdasarkan ID
func DeleteKeuangan(id string) error {
	return config.GetDB().
		Where("id = ?", id).
		Delete(&model.Transaction{}).
		Error
}
