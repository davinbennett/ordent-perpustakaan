package models

import "time"

type Loan struct {
	BaseModel
	UserID    uint       `gorm:"not null" json:"user_id"`
	BookID    uint       `gorm:"not null" json:"book_id"`
	BorrowedAt time.Time `json:"borrowed_at"`
	ReturnedAt *time.Time `json:"returned_at"`
	Status    string     `gorm:"type:varchar(20);not null" json:"status"` // borrowed | returned

	User User `gorm:"foreignKey:UserID"`
	Book Book `gorm:"foreignKey:BookID"`
}