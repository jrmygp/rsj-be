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

func (r *repository) FindAllNoPagination() ([]models.Warehouse, error) {
	var warehouses []models.Warehouse
	err := r.db.Find(&warehouses).Error

	return warehouses, err
}

func (r *repository) Create(warehouse models.Warehouse) (models.Warehouse, error) {
	err := r.db.Create(&warehouse).Error
	return warehouse, err
}

func (r *repository) FindByID(ID int) (models.Warehouse, error) {
	var warehouse models.Warehouse

	err := r.db.First(&warehouse, ID).Error
	return warehouse, err
}

func (r *repository) Edit(warehouse models.Warehouse) (models.Warehouse, error) {
	err := r.db.Save(&warehouse).Error
	return warehouse, err
}

func (r *repository) Delete(ID int) (models.Warehouse, error) {
	var warehouse models.Warehouse
	if err := r.db.First(&warehouse, ID).Error; err != nil {
		return warehouse, err
	}

	err := r.db.Delete(&warehouse).Error
	return warehouse, err
}

func (r *repository) FindAll(searchQuery string, offset int, pageSize int) (warehouse []models.Warehouse, totalCount int64) {
	result := r.db.Model(&models.Warehouse{})

	if searchQuery != "" {
		result = result.Where("name LIKE ? OR code LIKE ? OR flight_name LIKE ? OR flight_code LIKE ?", "%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%", "%"+searchQuery+"%")
	}

	result.Count(&totalCount)

	result = result.Offset(offset).Limit(pageSize)

	result.Find(&warehouse)

	return warehouse, totalCount
}
