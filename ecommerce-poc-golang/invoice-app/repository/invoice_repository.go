package repository

import (
	"invoice-app/models"

	"gorm.io/gorm"
)

type InvoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) *InvoiceRepository {
	return &InvoiceRepository{db: db}
}

func (r *InvoiceRepository) Create(invoice *models.Invoice) error {
	return r.db.Create(invoice).Error
}

func (r *InvoiceRepository) FindByID(id uint, userID string) (*models.Invoice, error) {
	var invoice models.Invoice
	err := r.db.Where("id = ? AND user_id = ?", id, userID).
		Preload("Items").
		First(&invoice).Error
	if err != nil {
		return nil, err
	}
	return &invoice, nil
}

func (r *InvoiceRepository) FindByUserID(userID string) ([]models.Invoice, error) {
	var invoices []models.Invoice
	err := r.db.Where("user_id = ?", userID).
		Preload("Items").
		Find(&invoices).Error
	return invoices, err
}
