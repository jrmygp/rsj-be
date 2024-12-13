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

func (r *repository) FindAllNoPagination() ([]models.Shipper, error) {
	var shippers []models.Shipper
	err := r.db.Find(&shippers).Error

	return shippers, err
}

func (r *repository) Create(shipper models.Shipper) (models.Shipper, error) {
	err := r.db.Create(&shipper).Error
	return shipper, err
}

func (r *repository) FindByID(ID int) (models.Shipper, error) {
	var shipper models.Shipper

	err := r.db.First(&shipper, ID).Error
	return shipper, err
}

func (r *repository) Edit(shipper models.Shipper) (models.Shipper, error) {
	err := r.db.Save(&shipper).Error
	return shipper, err
}

func (r *repository) Delete(ID int) (models.Shipper, error) {
	var shipper models.Shipper
	if err := r.db.First(&shipper, ID).Error; err != nil {
		return shipper, err
	}

	err := r.db.Delete(&shipper).Error
	return shipper, err
}

func (r *repository) FindAll(searchQuery string, offset int, pageSize int) (shipper []models.Shipper, totalCount int64) {
	result := r.db.Model(&models.Shipper{})

	if searchQuery != "" {
		result = result.Where("name LIKE ? OR address LIKE ?", "%"+searchQuery+"%", "%"+searchQuery+"%")
	}

	result.Count(&totalCount)

	result = result.Offset(offset).Limit(pageSize)

	result.Find(&shipper)

	return shipper, totalCount
}
