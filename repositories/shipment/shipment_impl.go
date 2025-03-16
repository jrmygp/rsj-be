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

func (r *repository) FindAll(searchQuery string, offset int, pageSize int) (shipment []models.Shipment, totalCount int64) {
	result := r.db.Model(&models.Shipment{})

	if searchQuery != "" {
		result = result.Where("shipment_number LIKE", "%"+searchQuery+"%")
	}

	result.Count(&totalCount)

	result = result.Order("created_at DESC").Offset(offset).Limit(pageSize)

	result.Preload("Quotations").
		Preload("InvoiceExports").
		Preload("InvoiceImports").
		Preload("InvoiceDoorToDoors").
		Preload("Warehouse").
		Find(&shipment)

	return shipment, totalCount
}

func (r *repository) Create(shipment models.Shipment) (models.Shipment, error) {
	err := r.db.Create(&shipment).Error
	if err == nil {
		err = r.db.Preload("Quotations").
			Preload("InvoiceExports").
			Preload("InvoiceImports").
			Preload("InvoiceDoorToDoors").
			Preload("Warehouse").
			First(&shipment, shipment.ID).Error
	}

	return shipment, err
}
