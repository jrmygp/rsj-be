package services

import (
	"errors"
	"server/models"
	repositories "server/repositories/warehouse"
	"server/requests"

	"gorm.io/gorm"
)

type service struct {
	repository repositories.Repository
}

func NewService(repository repositories.Repository) *service {
	return &service{repository}
}

func (s *service) FindAllNoPagination() ([]models.Warehouse, error) {
	warehouses, err := s.repository.FindAllNoPagination()
	return warehouses, err
}

func (s *service) Create(warehouseRequest requests.WarehouseRequest) (models.Warehouse, error) {
	warehouse := models.Warehouse{
		Category:   warehouseRequest.Category,
		Name:       warehouseRequest.Name,
		Code:       warehouseRequest.Code,
		FlightName: warehouseRequest.FlightName,
		FlightCode: warehouseRequest.FlightCode,
	}

	newWarehouse, err := s.repository.Create(warehouse)
	return newWarehouse, err
}

func (s *service) FindByID(ID int) (models.Warehouse, error) {
	warehouse, err := s.repository.FindByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Warehouse{}, nil
	}

	return warehouse, err
}

func (s *service) Edit(ID int, warehouseRequest requests.WarehouseRequest) (models.Warehouse, error) {
	warehouse, _ := s.repository.FindByID(ID)

	warehouse.Category = warehouseRequest.Category
	warehouse.Name = warehouseRequest.Name
	warehouse.Code = warehouseRequest.Code
	warehouse.FlightName = warehouseRequest.FlightName
	warehouse.FlightCode = warehouseRequest.FlightCode

	updatedWarehouse, err := s.repository.Edit(warehouse)
	return updatedWarehouse, err
}

func (s *service) Delete(ID int) (models.Warehouse, error) {
	warehouse, err := s.repository.Delete(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Warehouse{}, nil
	}

	return warehouse, err
}

func (s *service) FindAll(searchQuery string, page int) ([]models.Warehouse, int64, int, int, int) {
	if page < 1 {
		return []models.Warehouse{}, 0, 0, 0, 0 // Handle invalid page
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	warehouse, totalCount := s.repository.FindAll(searchQuery, offset, pageSize)

	firstRow := offset + 1
	lastRow := offset + len(warehouse)
	if len(warehouse) == 0 {
		firstRow = 0
		lastRow = 0
	}
	totalPages := (int(totalCount) + pageSize - 1) / pageSize

	return warehouse, totalCount, firstRow, lastRow, totalPages
}
