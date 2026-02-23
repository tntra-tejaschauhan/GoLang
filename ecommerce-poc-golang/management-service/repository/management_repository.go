package repository

import (
	"management-service/models"

	"gorm.io/gorm"
)

type ManagementRepository struct {
	db *gorm.DB
}

func NewManagementRepository(db *gorm.DB) *ManagementRepository {
	return &ManagementRepository{db: db}
}

func (r *ManagementRepository) RegisterInvoice(invoice *models.RegisteredInvoice) error {
	return r.db.Create(invoice).Error
}

func (r *ManagementRepository) FindAll() ([]models.RegisteredInvoice, error) {
	var invoices []models.RegisteredInvoice
	err := r.db.Find(&invoices).Error
	return invoices, err
}

func (r *ManagementRepository) FindByID(id uint) (*models.RegisteredInvoice, error) {
	var invoice models.RegisteredInvoice
	err := r.db.First(&invoice, id).Error
	if err != nil {
		return nil, err
	}
	return &invoice, nil
}
