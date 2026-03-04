package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"ordentperpustakaan/services"
	"ordentperpustakaan/utils"
)

func CreateBook(c *gin.Context) {
	var req struct {
		Title       string `json:"title" binding:"required"`
		Author      string `json:"author" binding:"required"`
		ISBN        string `json:"isbn" binding:"required"`
		Stock       int    `json:"stock" binding:"required"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request.")
		return
	}

	errStr := services.CreateBook(req.Title, req.Author, req.ISBN, req.Stock, req.Description)
	if errStr != "" {
		utils.BadRequestResponse(c, errStr)
		return
	}

	utils.SuccessResponse(c, gin.H{"message": "Book created successfully."})
}

func GetAllBooks(c *gin.Context) {
	search := c.Query("search")
	pageStr := c.DefaultQuery("page", "1")

	page, _ := strconv.Atoi(pageStr)

	books, errStr := services.GetAllBooks(search, page)
	if errStr != "" {
		utils.BadRequestResponse(c, errStr)
		return
	}

	utils.SuccessResponse(c, books)
}

func GetBookByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	book, errStr := services.GetBookByID(uint(id))
	if errStr != "" {
		utils.BadRequestResponse(c, errStr)
		return
	}

	utils.SuccessResponse(c, book)
}

func UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req struct {
		Title       string `json:"title"`
		Author      string `json:"author"`
		Stock       int    `json:"stock"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request.")
		return
	}

	errStr := services.UpdateBook(uint(id), req.Title, req.Author, req.Stock, req.Description)
	if errStr != "" {
		utils.BadRequestResponse(c, errStr)
		return
	}

	utils.SuccessResponse(c, gin.H{"message": "Book updated successfully."})
}

func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	errStr := services.DeleteBook(uint(id))
	if errStr != "" {
		utils.BadRequestResponse(c, errStr)
		return
	}

	utils.SuccessResponse(c, gin.H{"message": "Book deleted successfully."})
}