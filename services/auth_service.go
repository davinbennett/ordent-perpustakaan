package services

import (
	"strings"
	"time"

	"ordentperpustakaan/config"
	"ordentperpustakaan/models"
	"ordentperpustakaan/repositories"
	"ordentperpustakaan/utils"
)

func LoginWithEmail(email, password string) (string, uint, string) {
	user, errStr := repositories.GetUserByEmail(email)
	if errStr != "" {
		return "", 0, errStr
	}

	if user.PasswordHash == "" || !utils.CheckPasswordHash(password, user.PasswordHash) {
		return "", 0, "Incorrect email or password."
	}

	token, errStr := utils.GenerateJWT(user.ID, email)
	if errStr != "" {
		return "", 0, errStr
	}

	return token, user.ID, ""
}

func SendOTP(email string, isFrom string) string {
	email = strings.ToLower(email)

	if isFrom == "signup" {
		exists, _ := repositories.IsEmailExists(email)
		if exists {
			return "Email already registered."
		}
	}

	otp, err := utils.GenerateOTP()
	if err != nil {
		return "Failed to generate OTP."
	}

	hashedOTP, err := utils.HashPassword(otp)
	if err != nil {
		return "Failed to secure OTP."
	}

	key := "otp:" + email

	err = config.RedisClient.HSet(config.RedisCtx, key, map[string]interface{}{
		"code":     hashedOTP,
		"verified": false,
	}).Err()
	if err != nil {
		return "Failed to save OTP."
	}

	config.RedisClient.Expire(config.RedisCtx, key, 10*time.Minute)

	// kirim email jika ada
	_ = utils.SendOTPEmail(email, otp)

	return ""
}

func VerifyOTP(email, otp string) (bool, string) {
	email = strings.ToLower(email)
	key := "otp:" + email

	data, err := config.RedisClient.HGetAll(config.RedisCtx, key).Result()
	if err != nil || len(data) == 0 {
		return false, "OTP expired or not found."
	}

	if data["verified"] == "1" {
		return false, "OTP already verified."
	}

	if !utils.CheckPasswordHash(otp, data["code"]) {
		return false, "Incorrect OTP."
	}

	config.RedisClient.HSet(config.RedisCtx, key, "verified", true)

	return true, ""
}

func RegisterWithEmail(name, email, password string) (string, uint, string) {
	email = strings.ToLower(email)
	key := "otp:" + email

	verified, err := config.RedisClient.HGet(config.RedisCtx, key, "verified").Result()
	if err != nil || verified != "1" {
		return "", 0, "Email not verified."
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return "", 0, "Failed to secure password."
	}

	user := &models.User{
		Username:     name,
		Email:        email,
		PasswordHash: hashedPassword,
	}

	if errStr := repositories.CreateUser(user); errStr != "" {
		return "", 0, errStr
	}

	token, errStr := utils.GenerateJWT(user.ID, email)
	if errStr != "" {
		return "", 0, errStr
	}

	config.RedisClient.Del(config.RedisCtx, key)

	return token, user.ID, ""
}
