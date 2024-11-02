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

func (r *repository) Create(customer models.Customer) (models.Customer, error) {
	err := r.db.Create(&customer).Error
	return customer, err
}

func (r *repository) FindByID(ID int) (models.Customer, error) {
	var customer models.Customer

	err := r.db.First(&customer, ID).Error
	return customer, err
}

func (r *repository) Edit(customer models.Customer) (models.Customer, error) {
	err := r.db.Save(&customer).Error
	return customer, err
}

func (r *repository) Delete(ID int) (models.Customer, error) {
	var customer models.Customer
	if err := r.db.First(&customer, ID).Error; err != nil {
		return customer, err
	}

	err := r.db.Delete(&customer).Error
	return customer, err
}

func (r *repository) FindAll(searchQuery string, offset int, pageSize int) (customers []models.Customer, totalCount int64) {
	result := r.db.Model(&models.Customer{})

	if searchQuery != "" {
		result = result.Where("name LIKE ? OR address LIKE ?", "%"+searchQuery+"%", "%"+searchQuery+"%")
	}

	result.Count(&totalCount)

	result = result.Offset(offset).Limit(pageSize)

	result.Find(&customers)

	return customers, totalCount
}
