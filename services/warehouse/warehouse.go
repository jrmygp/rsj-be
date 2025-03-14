package services

import (
	"server/models"
	"server/requests"
)

type Service interface {
	FindAllNoPagination() ([]models.Warehouse, error)
	Create(warehouse requests.WarehouseRequest) (models.Warehouse, error)
	FindByID(ID int) (models.Warehouse, error)
	Edit(ID int, warehouse requests.WarehouseRequest) (models.Warehouse, error)
	Delete(ID int) (models.Warehouse, error)
	FindAll(searchQuery string, page int) ([]models.Warehouse, int64, int, int, int)
}
