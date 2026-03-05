package routes

import (
	"ordentperpustakaan/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoute(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	{
		auth.POST("/login", controllers.LoginWithEmail)
		auth.POST("/send-otp", controllers.RequestOTP)
		auth.POST("/verify-otp", controllers.VerifyOTP)
		auth.POST("/register", controllers.RegisterWithEmail)
	}
}