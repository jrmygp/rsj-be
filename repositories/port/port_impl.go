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

func (r *repository) FindAllNoPagination() ([]models.Port, error) {
	var ports []models.Port
	err := r.db.Find(&ports).Error

	return ports, err
}

func (r *repository) Create(port models.Port) (models.Port, error) {
	err := r.db.Create(&port).Error
	return port, err
}

func (r *repository) FindByID(ID int) (models.Port, error) {
	var port models.Port

	err := r.db.First(&port, ID).Error
	return port, err
}

func (r *repository) Edit(port models.Port) (models.Port, error) {
	err := r.db.Save(&port).Error
	return port, err
}

func (r *repository) Delete(ID int) (models.Port, error) {
	var port models.Port
	if err := r.db.First(&port, ID).Error; err != nil {
		return port, err
	}

	err := r.db.Delete(&port).Error
	return port, err
}

func (r *repository) FindAll(searchQuery string, offset int, pageSize int) (port []models.Port, totalCount int64) {
	result := r.db.Model(&models.Port{})

	if searchQuery != "" {
		result = result.Where("port_name LIKE ?", "%"+searchQuery+"%")
	}

	result.Count(&totalCount)

	result = result.Offset(offset).Limit(pageSize)

	result.Find(&port)

	return port, totalCount
}
