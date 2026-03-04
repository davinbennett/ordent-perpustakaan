package models

type Book struct {
	BaseModel
	Title       string  `gorm:"type:varchar(255);not null" json:"title"`
	Author      string  `gorm:"type:varchar(255);not null" json:"author"`
	ISBN        string  `gorm:"type:varchar(50);uniqueIndex;not null" json:"isbn"`
	Description string  `gorm:"type:text" json:"description,omitempty"`
	Stock       int     `gorm:"default:0" json:"stock"`
}