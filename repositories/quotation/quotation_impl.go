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

func (r *repository) Create(quotation models.Quotation) (models.Quotation, error) {
	err := r.db.Create(&quotation).Error
	return quotation, err
}

func (r *repository) FindByID(ID int) (models.Quotation, error) {
	var quotation models.Quotation

	err := r.db.First(&quotation, ID).Error
	return quotation, err
}

func (r *repository) Edit(quotation models.Quotation) (models.Quotation, error) {
	err := r.db.Save(&quotation).Error
	return quotation, err
}

func (r *repository) Delete(ID int) (models.Quotation, error) {
	var quotation models.Quotation
	if err := r.db.First(&quotation, ID).Error; err != nil {
		return quotation, err
	}

	err := r.db.Delete(&quotation).Error
	return quotation, err
}

func (r *repository) FindAll(searchQuery string, offset int, pageSize int) (quotation []models.Quotation, totalCount int64) {
	result := r.db.Model(&models.Quotation{})

	if searchQuery != "" {
		result = result.Where("quotation_number LIKE ?", "%"+searchQuery+"%")
	}

	result.Count(&totalCount)

	result = result.Offset(offset).Limit(pageSize)

	result.Find(&quotation)

	return quotation, totalCount
}
