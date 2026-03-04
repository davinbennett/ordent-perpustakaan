package controllers

import (
	"github.com/gin-gonic/gin"
	"ordentperpustakaan/services"
	"ordentperpustakaan/utils"
)

func LoginWithEmail(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Email and password must be filled correctly.")
		return
	}

	token, userID, errStr := services.LoginWithEmail(req.Email, req.Password)
	if errStr != "" {
		utils.BadRequestResponse(c, errStr)
		return
	}

	utils.SuccessResponse(c, gin.H{
		"access_token": token,
		"user_id":      userID,
	})
}

func SendOTP(c *gin.Context) {
	var req struct {
		Email  string `json:"email" binding:"required,email"`
		IsFrom string `json:"is_from" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request.")
		return
	}

	errStr := services.SendOTP(req.Email, req.IsFrom)
	if errStr != "" {
		utils.BadRequestResponse(c, errStr)
		return
	}

	utils.SuccessResponse(c, gin.H{"message": "OTP sent successfully."})
}

func VerifyOTP(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
		OTP   string `json:"otp" binding:"required,len=6"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request.")
		return
	}

	success, errStr := services.VerifyOTP(req.Email, req.OTP)
	if !success {
		utils.BadRequestResponse(c, errStr)
		return
	}

	utils.SuccessResponse(c, gin.H{"message": "OTP verified."})
}

func RegisterWithEmail(c *gin.Context) {
	var req struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request.")
		return
	}

	token, userID, errStr := services.RegisterWithEmail(req.Name, req.Email, req.Password)
	if errStr != "" {
		utils.BadRequestResponse(c, errStr)
		return
	}

	utils.SuccessResponse(c, gin.H{
		"access_token": token,
		"user_id":      userID,
	})
}

func RequestOTP(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required,email"`
		IsFrom string `json:"is_from" binding:"omitempty"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid email format.")
		return
	}

	isFrom := "default"
	if req.IsFrom != "" {
		isFrom = req.IsFrom
	}

	if err := services.SendOTP(req.Email, isFrom); err != "" {
		utils.InternalServerErrorResponse(c, err)
		return
	}

	utils.SuccessResponse(c, "OTP sent to your email.")
}