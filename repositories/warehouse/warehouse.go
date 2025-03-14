package repositories

import "server/models"

type Repository interface {
	FindAllNoPagination() ([]models.Warehouse, error)
	Create(warehouse models.Warehouse) (models.Warehouse, error)
	FindByID(ID int) (models.Warehouse, error)
	Edit(warehouse models.Warehouse) (models.Warehouse, error)
	Delete(ID int) (models.Warehouse, error)
	FindAll(searchQuery string, offset int, pageSize int) (warehouse []models.Warehouse, totalCount int64)
}
