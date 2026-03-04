package models

import "time"

type User struct {
	BaseModel
	Username      string     `gorm:"type:varchar(100);uniqueIndex;not null" json:"username"`
	Email         string     `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	PasswordHash  string     `gorm:"type:text;not null" json:"-"`
	RoleID        string     `gorm:"type:uuid;index" json:"role_id"`
	Role          Role       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	EmailVerified bool       `gorm:"default:false" json:"email_verified"`
	LastLogin     *time.Time `json:"last_login,omitempty"`
}