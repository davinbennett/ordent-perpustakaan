package models

import (
	"time"
)

type BaseModel struct {
	ID           uint       `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}