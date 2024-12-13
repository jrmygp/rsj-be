package services

import (
	"errors"
	"server/models"
	repositories "server/repositories/shipper"
	"server/requests"

	"gorm.io/gorm"
)

type service struct {
	repository repositories.Repository
}

func NewService(repository repositories.Repository) *service {
	return &service{repository}
}

func (s *service) FindAllNoPagination() ([]models.Shipper, error) {
	shippers, err := s.repository.FindAllNoPagination()
	return shippers, err
}

func (s *service) Create(shipperRequest requests.CreateShipperRequest) (models.Shipper, error) {
	shipper := models.Shipper{
		Name:    shipperRequest.Name,
		Address: shipperRequest.Address,
	}

	newShipper, err := s.repository.Create(shipper)
	return newShipper, err
}

func (s *service) FindByID(ID int) (models.Shipper, error) {
	shipper, err := s.repository.FindByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Shipper{}, nil
	}

	return shipper, err
}

func (s *service) Edit(ID int, shipperRequest requests.EditShipperRequest) (models.Shipper, error) {
	shipper, _ := s.repository.FindByID(ID)

	if shipperRequest.Name != "" {
		shipper.Name = shipperRequest.Name
	}
	if shipperRequest.Address != "" {
		shipper.Address = shipperRequest.Address
	}

	updatedPort, err := s.repository.Edit(shipper)
	return updatedPort, err
}

func (s *service) Delete(ID int) (models.Shipper, error) {
	shipper, err := s.repository.Delete(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Shipper{}, nil
	}

	return shipper, err
}

func (s *service) FindAll(searchQuery string, page int) ([]models.Shipper, int64, int, int, int) {
	if page < 1 {
		return []models.Shipper{}, 0, 0, 0, 0 // Handle invalid page
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	shipper, totalCount := s.repository.FindAll(searchQuery, offset, pageSize)

	firstRow := offset + 1
	lastRow := offset + len(shipper)
	if len(shipper) == 0 {
		firstRow = 0
		lastRow = 0
	}
	totalPages := (int(totalCount) + pageSize - 1) / pageSize

	return shipper, totalCount, firstRow, lastRow, totalPages
}
