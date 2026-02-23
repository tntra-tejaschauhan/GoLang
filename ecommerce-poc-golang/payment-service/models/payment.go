package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Amount    float64        `gorm:"not null" json:"amount"`
	Status    string         `gorm:"not null" json:"status"` // SUCCESS, FAILED
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type ProcessPaymentRequest struct {
	Amount float64 `json:"amount" binding:"required,gt=0"`
}
