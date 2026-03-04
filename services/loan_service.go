package services

import (
	"time"

	"ordentperpustakaan/models"
	"ordentperpustakaan/repositories"
)

func BorrowBook(userID, bookID uint) string {
	// cek book
	book, errStr := repositories.GetBookByID(bookID)
	if errStr != "" {
		return errStr
	}

	if book.Stock <= 0 {
		return "Book is out of stock."
	}

	// kurangi stock
	book.Stock -= 1
	if err := repositories.UpdateBook(book); err != "" {
		return err
	}

	loan := &models.Loan{
		UserID:     userID,
		BookID:     bookID,
		BorrowedAt: time.Now(),
		Status:     "borrowed",
	}

	return repositories.CreateLoan(loan)
}

func ReturnBook(userID, bookID uint) string {
	loan, errStr := repositories.GetActiveLoan(userID, bookID)
	if errStr != "" {
		return "No active borrowing found."
	}

	now := time.Now()
	loan.Status = "returned"
	loan.ReturnedAt = &now

	if err := repositories.UpdateLoan(loan); err != "" {
		return err
	}

	// tambah stock
	book, errStr := repositories.GetBookByID(bookID)
	if errStr != "" {
		return errStr
	}

	book.Stock += 1
	return repositories.UpdateBook(book)
}