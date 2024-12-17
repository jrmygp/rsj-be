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
	err := r.db.Preload("Customer").Preload("Shipper").Find(&invoices).Error

	return invoices, err
}

func (r *repository) Create(invoice models.Invoice) (models.Invoice, error) {
	err := r.db.Create(&invoice).Error
	if err == nil {
		err = r.db.Preload("Customer").Preload("Shipper").First(&invoice, invoice.ID).Error
	}

	return invoice, err
}

func (r *repository) FindByID(ID int) (models.Invoice, error) {
	var invoice models.Invoice

	err := r.db.Preload("Customer").Preload("Shipper").First(&invoice, ID).Error

	return invoice, err
}

func (r *repository) Edit(invoice models.Invoice) (models.Invoice, error) {
	err := r.db.Save(&invoice).Error
	if err == nil {
		err = r.db.Preload("Customer").Preload("Shipper").First(&invoice, invoice.ID).Error
	}

	return invoice, err
}

func (r *repository) Delete(ID int) (models.Invoice, error) {
	var invoice models.Invoice
	if err := r.db.Preload("Customer").Preload("Shipper").First(&invoice, ID).Error; err != nil {
		return invoice, err
	}

	err := r.db.Delete(&invoice).Error
	return invoice, err
}

func (r *repository) FindAll(searchQuery string, offset int, pageSize int) (invoice []models.Invoice, totalCount int64) {
	result := r.db.Model(&models.Invoice{})

	if searchQuery != "" {
		result = result.Where("invoice_number LIKE ?", "%"+searchQuery+"%")
	}

	result.Count(&totalCount)

	result = result.Order("created_at DESC").Offset(offset).Limit(pageSize)

	result.Preload("Customer").Preload("Shipper").Find(&invoice)

	return invoice, totalCount
}
