package routes

import (
	"ordentperpustakaan/controllers"

	"github.com/gin-gonic/gin"
)

func LoanRoutes(r *gin.RouterGroup) {
	loans := r.Group("/loans")
	{
		loans.POST("/borrow/:book_id", controllers.BorrowBook)
		loans.POST("/return/:book_id", controllers.ReturnBook)
	}
}