package services

import (
	"errors"
	"server/models"
	repositories "server/repositories/customer"
	"server/requests"

	"gorm.io/gorm"
)

type service struct {
	repository repositories.Repository
}

func NewService(repository repositories.Repository) *service {
	return &service{repository}
}

func (s *service) Create(customerRequest requests.CreateCustomerRequest) (models.Customer, error) {
	customer := models.Customer{
		Name:    customerRequest.Name,
		Address: customerRequest.Address,
	}

	newCustomer, err := s.repository.Create(customer)
	return newCustomer, err
}

func (s *service) FindByID(ID int) (models.Customer, error) {
	customer, err := s.repository.FindByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Customer{}, nil
	}

	return customer, err
}

func (s *service) Edit(ID int, customerRequest requests.EditCustomerRequest) (models.Customer, error) {
	customer, _ := s.repository.FindByID(ID)

	if customerRequest.Name != "" {
		customer.Name = customerRequest.Name
	}
	if customerRequest.Address != "" {
		customer.Address = customerRequest.Address
	}

	updatedCustomer, err := s.repository.Edit(customer)
	return updatedCustomer, err
}

func (s *service) Delete(ID int) (models.Customer, error) {
	customer, err := s.repository.Delete(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Customer{}, nil
	}

	return customer, err
}

func (s *service) FindAll(searchQuery string, page int) ([]models.Customer, int64, int, int, int) {
	if page < 1 {
		return []models.Customer{}, 0, 0, 0, 0 // Handle invalid page
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	customers, totalCount := s.repository.FindAll(searchQuery, offset, pageSize)

	firstRow := offset + 1
	lastRow := offset + len(customers)
	if len(customers) == 0 {
		firstRow = 0
		lastRow = 0
	}
	totalPages := (int(totalCount) + pageSize - 1) / pageSize

	return customers, totalCount, firstRow, lastRow, totalPages
}
