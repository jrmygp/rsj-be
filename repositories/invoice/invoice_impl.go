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
	err := r.db.Preload("Customer").Preload("Shipper").Preload("Consignee").Preload("PortOfLoading").Preload("PortOfDischarge").Find(&invoices).Error

	return invoices, err
}

func (r *repository) Create(invoice models.Invoice) (models.Invoice, error) {
	err := r.db.Create(&invoice).Error
	if err == nil {
		err = r.db.Preload("Customer").Preload("Shipper").Preload("Consignee").Preload("PortOfLoading").Preload("PortOfDischarge").First(&invoice, invoice.ID).Error
	}

	return invoice, err
}

func (r *repository) FindByID(ID int) (models.Invoice, error) {
	var invoice models.Invoice

	err := r.db.Preload("Customer").Preload("Shipper").Preload("Consignee").Preload("PortOfLoading").Preload("PortOfDischarge").First(&invoice, ID).Error

	return invoice, err
}

func (r *repository) Edit(invoice models.Invoice) (models.Invoice, error) {
	err := r.db.Save(&invoice).Error
	if err == nil {
		err = r.db.Preload("Customer").Preload("Shipper").Preload("Consignee").Preload("PortOfLoading").Preload("PortOfDischarge").First(&invoice, invoice.ID).Error
	}

	return invoice, err
}

func (r *repository) Delete(ID int) (models.Invoice, error) {
	var invoice models.Invoice
	if err := r.db.Preload("Customer").Preload("Shipper").Preload("Consignee").Preload("PortOfLoading").Preload("PortOfDischarge").First(&invoice, ID).Error; err != nil {
		return invoice, err
	}

	err := r.db.Delete(&invoice).Error
	return invoice, err
}

func (r *repository) FindAll(searchQuery string, offset int, pageSize int, customerID int, category string) (invoices []models.Invoice, totalCount int64) {
	result := r.db.Model(&models.Invoice{})

	if customerID > 0 {
		result = result.Where("customer_id = ?", customerID)
	}

	if category != "" {
		result = result.Where("category = ?", category)
	}

	if searchQuery != "" {
		result = result.Where("invoice_number LIKE ? OR category LIKE ?", "%"+searchQuery+"%", "%"+searchQuery+"%")
	}

	result.Count(&totalCount)

	result.Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Preload("Customer").
		Preload("Shipper").
		Preload("Consignee").
		Preload("PortOfLoading").
		Preload("PortOfDischarge").
		Find(&invoices)

	return invoices, totalCount
}

// Door to Door implementations
func (r *repository) FindAllDoorToDoorNoPagination() ([]models.DoorToDoorInvoice, error) {
	var invoices []models.DoorToDoorInvoice
	err := r.db.Preload("Customer").Preload("Shipper").Preload("Consignee").Preload("PortOfLoading").Preload("PortOfDischarge").Find(&invoices).Error

	return invoices, err
}

func (r *repository) CreateDoorToDoor(invoice models.DoorToDoorInvoice) (models.DoorToDoorInvoice, error) {
	err := r.db.Create(&invoice).Error
	if err == nil {
		err = r.db.Preload("Customer").Preload("Shipper").Preload("Consignee").Preload("PortOfLoading").Preload("PortOfDischarge").First(&invoice, invoice.ID).Error
	}

	return invoice, err
}

func (r *repository) FindDoorToDoorByID(ID int) (models.DoorToDoorInvoice, error) {
	var invoice models.DoorToDoorInvoice

	err := r.db.Preload("Customer").Preload("Shipper").Preload("Consignee").Preload("PortOfLoading").Preload("PortOfDischarge").First(&invoice, ID).Error

	return invoice, err
}

func (r *repository) EditDoorToDoor(invoice models.DoorToDoorInvoice) (models.DoorToDoorInvoice, error) {
	err := r.db.Save(&invoice).Error
	if err == nil {
		err = r.db.Preload("Customer").Preload("Shipper").Preload("Consignee").Preload("PortOfLoading").Preload("PortOfDischarge").First(&invoice, invoice.ID).Error
	}

	return invoice, err
}

func (r *repository) DeleteDoorToDoor(ID int) (models.DoorToDoorInvoice, error) {
	var invoice models.DoorToDoorInvoice
	if err := r.db.Preload("Customer").Preload("Shipper").Preload("Consignee").Preload("PortOfLoading").Preload("PortOfDischarge").First(&invoice, ID).Error; err != nil {
		return invoice, err
	}

	err := r.db.Delete(&invoice).Error
	return invoice, err
}

func (r *repository) FindAllDoorToDoor(searchQuery string, offset int, pageSize int, customerID int) (invoice []models.DoorToDoorInvoice, totalCount int64) {
	result := r.db.Model(&models.DoorToDoorInvoice{})

	if customerID > 0 {
		result = result.Where("customer_id = ?", customerID)
	}

	if searchQuery != "" {
		result = result.Where("invoice_number LIKE ? OR category LIKE ?", "%"+searchQuery+"%", "%"+searchQuery+"%")
	}

	result.Count(&totalCount)

	result.Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Preload("Customer").
		Preload("Shipper").
		Preload("Consignee").
		Preload("PortOfLoading").
		Preload("PortOfDischarge").
		Find(&invoice)

	return invoice, totalCount
}
