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

// Export implementations
func (r *repository) FindAllExportNoPagination() ([]models.InvoiceExport, error) {
	var invoices []models.InvoiceExport
	err := r.db.Preload("Customer").Preload("Shipper").Preload("Consignee").Preload("PortOfLoading").Preload("PortOfDischarge").Find(&invoices).Error

	return invoices, err
}

func (r *repository) CreateExport(invoice models.InvoiceExport) (models.InvoiceExport, error) {
	err := r.db.Create(&invoice).Error
	if err == nil {
		err = r.db.Preload("Customer").Preload("Shipper").Preload("Consignee").Preload("PortOfLoading").Preload("PortOfDischarge").First(&invoice, invoice.ID).Error
	}

	return invoice, err
}

func (r *repository) FindExportByID(ID int) (models.InvoiceExport, error) {
	var invoice models.InvoiceExport

	err := r.db.Preload("Customer").Preload("Shipper").Preload("Consignee").Preload("PortOfLoading").Preload("PortOfDischarge").First(&invoice, ID).Error

	return invoice, err
}

func (r *repository) EditExport(invoice models.InvoiceExport) (models.InvoiceExport, error) {
	err := r.db.Save(&invoice).Error
	if err == nil {
		err = r.db.Preload("Customer").Preload("Shipper").Preload("Consignee").Preload("PortOfLoading").Preload("PortOfDischarge").First(&invoice, invoice.ID).Error
	}

	return invoice, err
}

func (r *repository) DeleteExport(ID int) (models.InvoiceExport, error) {
	var invoice models.InvoiceExport
	if err := r.db.Preload("Customer").Preload("Shipper").Preload("Consignee").Preload("PortOfLoading").Preload("PortOfDischarge").First(&invoice, ID).Error; err != nil {
		return invoice, err
	}

	err := r.db.Delete(&invoice).Error
	return invoice, err
}

func (r *repository) FindAllExport(searchQuery string, offset int, pageSize int, customerID int) (invoices []models.InvoiceExport, totalCount int64) {
	result := r.db.Model(&models.InvoiceExport{})

	if customerID > 0 {
		result = result.Where("customer_id = ?", customerID)
	}

	if searchQuery != "" {
		result = result.Where("invoice_number LIKE ?", "%"+searchQuery+"%")
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

func (r *repository) FindExportByIDs(IDs []int) ([]models.InvoiceExport, error) {
	var invoices []models.InvoiceExport

	err := r.db.Preload("Customer").
		Preload("Shipper").
		Preload("Consignee").
		Preload("PortOfLoading").
		Preload("PortOfDischarge").
		Where("id IN (?)", IDs).
		Find(&invoices).Error

	return invoices, err
}

// Import implementations
func (r *repository) FindAllImportNoPagination() ([]models.InvoiceImport, error) {
	var invoices []models.InvoiceImport
	err := r.db.Preload("Customer").Preload("Shipper").Preload("Consignee").Preload("PortOfLoading").Preload("PortOfDischarge").Find(&invoices).Error

	return invoices, err
}

func (r *repository) CreateImport(invoice models.InvoiceImport) (models.InvoiceImport, error) {
	err := r.db.Create(&invoice).Error
	if err == nil {
		err = r.db.Preload("Customer").Preload("Shipper").Preload("Consignee").Preload("PortOfLoading").Preload("PortOfDischarge").First(&invoice, invoice.ID).Error
	}

	return invoice, err
}

func (r *repository) FindImportByID(ID int) (models.InvoiceImport, error) {
	var invoice models.InvoiceImport

	err := r.db.Preload("Customer").Preload("Shipper").Preload("Consignee").Preload("PortOfLoading").Preload("PortOfDischarge").First(&invoice, ID).Error

	return invoice, err
}

func (r *repository) EditImport(invoice models.InvoiceImport) (models.InvoiceImport, error) {
	err := r.db.Save(&invoice).Error
	if err == nil {
		err = r.db.Preload("Customer").Preload("Shipper").Preload("Consignee").Preload("PortOfLoading").Preload("PortOfDischarge").First(&invoice, invoice.ID).Error
	}

	return invoice, err
}

func (r *repository) DeleteImport(ID int) (models.InvoiceImport, error) {
	var invoice models.InvoiceImport
	if err := r.db.Preload("Customer").Preload("Shipper").Preload("Consignee").Preload("PortOfLoading").Preload("PortOfDischarge").First(&invoice, ID).Error; err != nil {
		return invoice, err
	}

	err := r.db.Delete(&invoice).Error
	return invoice, err
}

func (r *repository) FindAllImport(searchQuery string, offset int, pageSize int, customerID int) (invoices []models.InvoiceImport, totalCount int64) {
	result := r.db.Model(&models.InvoiceImport{})

	if customerID > 0 {
		result = result.Where("customer_id = ?", customerID)
	}

	if searchQuery != "" {
		result = result.Where("invoice_number LIKE ?", "%"+searchQuery+"%")
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

func (r *repository) FindImportByIDs(IDs []int) ([]models.InvoiceImport, error) {
	var invoices []models.InvoiceImport

	err := r.db.Preload("Customer").
		Preload("Shipper").
		Preload("Consignee").
		Preload("PortOfLoading").
		Preload("PortOfDischarge").
		Where("id IN (?)", IDs).
		Find(&invoices).Error

	return invoices, err
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

func (r *repository) FindDoorToDoorByIDs(IDs []int) ([]models.DoorToDoorInvoice, error) {
	var invoices []models.DoorToDoorInvoice

	err := r.db.Preload("Customer").
		Preload("Shipper").
		Preload("Consignee").
		Preload("PortOfLoading").
		Preload("PortOfDischarge").
		Where("id IN (?)", IDs).
		Find(&invoices).Error

	return invoices, err
}
