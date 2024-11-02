package services

import (
	"server/models"
	"server/requests"
)

type Service interface {
	Create(customer requests.CreateCustomerRequest) (models.Customer, error)
	FindByID(ID int) (models.Customer, error)
	Edit(ID int, customer requests.EditCustomerRequest) (models.Customer, error)
	Delete(ID int) (models.Customer, error)
	FindAll(searchQuery string, page int) ([]models.Customer, int64, int, int, int)
}
