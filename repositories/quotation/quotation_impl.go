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
