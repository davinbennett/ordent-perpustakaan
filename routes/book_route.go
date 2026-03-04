package routes

import (
	"ordentperpustakaan/controllers"
	"github.com/gin-gonic/gin"
)

func BookRoutes(r *gin.RouterGroup) {
	books := r.Group("/books")
	{
		books.POST("", controllers.CreateBook)
		books.GET("", controllers.GetAllBooks)
		books.GET("/:id", controllers.GetBookByID)
		books.PUT("/:id", controllers.UpdateBook)
		books.DELETE("/:id", controllers.DeleteBook)
	}
}