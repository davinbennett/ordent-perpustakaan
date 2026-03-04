package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"ordentperpustakaan/services"
	"ordentperpustakaan/utils"
)

func BorrowBook(c *gin.Context) {
	userID := c.GetUint("user_id")

	bookID, _ := strconv.Atoi(c.Param("book_id"))

	errStr := services.BorrowBook(userID, uint(bookID))
	if errStr != "" {
		utils.BadRequestResponse(c, errStr)
		return
	}

	utils.SuccessResponse(c, gin.H{
		"message": "Book borrowed successfully.",
	})
}

func ReturnBook(c *gin.Context) {
	userID := c.GetUint("user_id")

	bookID, _ := strconv.Atoi(c.Param("book_id"))

	errStr := services.ReturnBook(userID, uint(bookID))
	if errStr != "" {
		utils.BadRequestResponse(c, errStr)
		return
	}

	utils.SuccessResponse(c, gin.H{
		"message": "Book returned successfully.",
	})
}