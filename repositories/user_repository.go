package repositories

import (
	"ordentperpustakaan/config"
	"ordentperpustakaan/models"
)

func CreateUser(user *models.User) string {
	if err := config.PostgresDB.Create(user).Error; err != nil {
		return "Failed to create user."
	}
	return ""
}

func GetUserByEmail(email string) (*models.User, string) {
	var user models.User

	err := config.PostgresDB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, "Email not registered."
	}

	return &user, ""
}

func IsEmailExists(email string) (bool, string) {
	var count int64

	config.PostgresDB.Model(&models.User{}).
		Where("email = ?", email).
		Count(&count)

	if count > 0 {
		return true, "Email already registered."
	}

	return false, ""
}