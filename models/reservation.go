package models

type Reservation struct {
	BaseModel
	UserID string `gorm:"type:uuid;index;not null" json:"user_id"`
	BookID string `gorm:"type:uuid;index;not null" json:"book_id"`
	Status string `gorm:"type:varchar(50);default:'waiting'" json:"status"`
}