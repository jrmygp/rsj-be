package repositories

import (
	"server/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllNoPagination() ([]models.Invoice, error) {
	var invoices []models.Invoice
	err := r.db.Preload("Customer").Preload("User").Find(&invoices).Error

	return invoices, err
}

func (r *repository) Create(invoice models.Invoice) (models.Invoice, error) {
	err := r.db.Create(&invoice).Error
	if err == nil {
		err = r.db.Preload("Customer").Preload("User").First(&invoice, invoice.ID).Error
	}

	return invoice, err
}

func (r *repository) FindByID(ID int) (models.Invoice, error) {
	var invoice models.Invoice

	err := r.db.Preload("Customer").Preload("User").First(&invoice, ID).Error

	return invoice, err
}
