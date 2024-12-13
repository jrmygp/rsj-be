package repositories

import "server/models"

type Repository interface {
	FindAllNoPagination() ([]models.Shipper, error)
	Create(port models.Shipper) (models.Shipper, error)
	FindByID(ID int) (models.Shipper, error)
	Edit(port models.Shipper) (models.Shipper, error)
	Delete(ID int) (models.Shipper, error)
	FindAll(searchQuery string, offset int, pageSize int) (port []models.Shipper, totalCount int64)
}
