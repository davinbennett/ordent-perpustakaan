package routes

import (
	"ordentperpustakaan/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		AuthRoute(api)

		protected := api.Group("/")
		protected.Use(middleware.JWTMiddleware())
		{
			BookRoutes(protected)
			LoanRoutes(protected)
		}
	}
}