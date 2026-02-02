package repository

import (
	"tugasbesar/config"
	"tugasbesar/model"
)

// Cari user berdasarkan username
func FindUserByUsername(username string) (model.User, error) {
	var user model.User
	err := config.GetDB().
		Where("username = ?", username).
		First(&user).Error
	return user, err
}

// Simpan user baru
func CreateUser(user model.User) (model.User, error) {
	err := config.GetDB().Create(&user).Error
	return user, err
}
