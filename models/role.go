package models

type Role struct {
	BaseModel
	Name string `gorm:"type:varchar(50);uniqueIndex;not null" json:"name"`
}