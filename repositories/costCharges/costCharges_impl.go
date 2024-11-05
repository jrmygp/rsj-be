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

func (r *repository) FindAllNoPagination() ([]models.CostCharges, error) {
	var costCharges []models.CostCharges
	err := r.db.Find(&costCharges).Error

	return costCharges, err
}

func (r *repository) Create(costCharges models.CostCharges) (models.CostCharges, error) {
	err := r.db.Create(&costCharges).Error
	return costCharges, err
}

func (r *repository) FindByID(ID int) (models.CostCharges, error) {
	var costCharges models.CostCharges

	err := r.db.First(&costCharges, ID).Error
	return costCharges, err
}

func (r *repository) Edit(costCharges models.CostCharges) (models.CostCharges, error) {
	err := r.db.Save(&costCharges).Error
	return costCharges, err
}

func (r *repository) Delete(ID int) (models.CostCharges, error) {
	var costCharges models.CostCharges
	if err := r.db.First(&costCharges, ID).Error; err != nil {
		return costCharges, err
	}

	err := r.db.Delete(&costCharges).Error
	return costCharges, err
}

func (r *repository) FindAll(searchQuery string, offset int, pageSize int) (costCharges []models.CostCharges, totalCount int64) {
	result := r.db.Model(&models.CostCharges{})

	if searchQuery != "" {
		result = result.Where("name LIKE ?", "%"+searchQuery+"%")
	}

	result.Count(&totalCount)

	result = result.Offset(offset).Limit(pageSize)

	result.Find(&costCharges)

	return costCharges, totalCount
}
