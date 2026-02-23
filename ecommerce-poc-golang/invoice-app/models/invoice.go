package models

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	UserID      string         `gorm:"not null;index" json:"user_id"`
	TotalAmount float64        `gorm:"not null" json:"total_amount"`
	PaymentID   uint           `gorm:"not null" json:"payment_id"`
	Items       []InvoiceItem  `gorm:"foreignKey:InvoiceID" json:"items"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type InvoiceItem struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	InvoiceID uint           `gorm:"not null;index" json:"invoice_id"`
	ProductID uint           `gorm:"not null" json:"product_id"`
	Quantity  int            `gorm:"not null" json:"quantity"`
	Price     float64        `gorm:"not null" json:"price"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type CreateInvoiceRequest struct {
	UserID      string               `json:"user_id" binding:"required"`
	TotalAmount float64              `json:"total_amount" binding:"required"`
	PaymentID   uint                 `json:"payment_id" binding:"required"`
	Items       []CreateInvoiceItem  `json:"items" binding:"required"`
}

type CreateInvoiceItem struct {
	ProductID uint    `json:"product_id" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
}
