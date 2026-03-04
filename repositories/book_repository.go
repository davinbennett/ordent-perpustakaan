package repositories

import (
	"ordentperpustakaan/config"
	"ordentperpustakaan/models"
)

func CreateBook(book *models.Book) string {
	if err := config.PostgresDB.Create(book).Error; err != nil {
		return "Failed to create book."
	}
	return ""
}

func GetAllBooks(books *[]models.Book, search string, limit, offset int) string {
	query := config.PostgresDB.Model(&models.Book{})

	if search != "" {
		query = query.Where("title ILIKE ? OR author ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Limit(limit).Offset(offset).Find(books).Error; err != nil {
		return "Failed to fetch books."
	}

	return ""
}

func GetBookByID(id uint) (*models.Book, string) {
	var book models.Book

	if err := config.PostgresDB.First(&book, id).Error; err != nil {
		return nil, "Book not found."
	}

	return &book, ""
}

func UpdateBook(book *models.Book) string {
	if err := config.PostgresDB.Save(book).Error; err != nil {
		return "Failed to update book."
	}
	return ""
}

func DeleteBook(id uint) string {
	if err := config.PostgresDB.Delete(&models.Book{}, id).Error; err != nil {
		return "Failed to delete book."
	}
	return ""
}