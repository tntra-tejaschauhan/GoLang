package models

import (
	"time"

	"gorm.io/gorm"
)

type RegisteredInvoice struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	InvoiceID  uint           `gorm:"not null;index" json:"invoice_id"`
	UserID     string         `gorm:"not null;index" json:"user_id"`
	Amount     float64        `gorm:"not null" json:"amount"`
	ReceivedAt time.Time      `gorm:"autoCreateTime" json:"received_at"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

type RegisterInvoiceRequest struct {
	InvoiceID uint    `json:"invoice_id" binding:"required"`
	UserID    string  `json:"user_id" binding:"required"`
	Amount    float64 `json:"amount" binding:"required"`
}
