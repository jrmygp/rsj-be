package repositories

import "server/models"

type Repository interface {
	Create(customer models.Customer) (models.Customer, error)
	FindByID(ID int) (models.Customer, error)
	Edit(customer models.Customer) (models.Customer, error)
	Delete(ID int) (models.Customer, error)
	FindAll(searchQuery string, offset int, pageSize int) (customers []models.Customer, totalCount int64)
}
