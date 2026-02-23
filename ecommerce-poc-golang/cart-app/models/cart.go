package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	UserID    string         `gorm:"not null;index" json:"user_id"`
	Status    string         `gorm:"not null;default:'active'" json:"status"` // active, checked_out
	Items     []CartItem     `gorm:"foreignKey:CartID" json:"items"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type CartItem struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CartID    uint           `gorm:"not null;index" json:"cart_id"`
	ProductID uint           `gorm:"not null" json:"product_id"`
	Quantity  int            `gorm:"not null" json:"quantity"`
	Price     float64        `gorm:"not null" json:"price"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type AddItemRequest struct {
	ProductID uint    `json:"product_id" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required,gt=0"`
	Price     float64 `json:"price" binding:"required,gt=0"`
}

type CheckoutResponse struct {
	Message   string      `json:"message"`
	CartID    uint        `json:"cart_id"`
	Total     float64     `json:"total"`
	PaymentID uint        `json:"payment_id"`
	InvoiceID uint        `json:"invoice_id"`
}
