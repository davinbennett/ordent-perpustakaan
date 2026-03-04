package repositories

import (
	"ordentperpustakaan/config"
	"ordentperpustakaan/models"
)

func CreateLoan(loan *models.Loan) string {
	if err := config.PostgresDB.Create(loan).Error; err != nil {
		return "Failed to create loan."
	}
	return ""
}

func GetActiveLoan(userID, bookID uint) (*models.Loan, string) {
	var loan models.Loan

	err := config.PostgresDB.
		Where("user_id = ? AND book_id = ? AND status = ?", userID, bookID, "borrowed").
		First(&loan).Error

	if err != nil {
		return nil, "Active loan not found."
	}

	return &loan, ""
}

func UpdateLoan(loan *models.Loan) string {
	if err := config.PostgresDB.Save(loan).Error; err != nil {
		return "Failed to update loan."
	}
	return ""
}