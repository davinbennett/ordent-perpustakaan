package services

import (
	"encoding/json"
	"fmt"
	"time"

	"ordentperpustakaan/config"
	"ordentperpustakaan/models"
	"ordentperpustakaan/repositories"
)

func CreateBook(title, author, isbn string, stock int, description string) string {
	book := &models.Book{
		Title:       title,
		Author:      author,
		ISBN:        isbn,
		Stock:       stock,
		Description: description,
	}

	return repositories.CreateBook(book)
}

func GetAllBooks(search string, page int) ([]models.Book, string) {
	var books []models.Book

	limit := 10
	offset := (page - 1) * limit

	cacheKey := fmt.Sprintf("books:%s:%d", search, page)

	// Try Redis
	cached, err := config.RedisClient.Get(config.RedisCtx, cacheKey).Result()
	if err == nil {
		json.Unmarshal([]byte(cached), &books)
		return books, ""
	}

	// Fetch from DB
	errStr := repositories.GetAllBooks(&books, search, limit, offset)
	if errStr != "" {
		return nil, errStr
	}

	// Save to Redis
	data, _ := json.Marshal(books)
	config.RedisClient.Set(config.RedisCtx, cacheKey, data, 5*time.Minute)

	return books, ""
}

func GetBookByID(id uint) (*models.Book, string) {
	return repositories.GetBookByID(id)
}

func UpdateBook(id uint, title, author string, stock int, description string) string {
	book, errStr := repositories.GetBookByID(id)
	if errStr != "" {
		return errStr
	}

	book.Title = title
	book.Author = author
	book.Stock = stock
	book.Description = description

	return repositories.UpdateBook(book)
}

func DeleteBook(id uint) string {
	return repositories.DeleteBook(id)
}