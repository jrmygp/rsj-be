package repositories

import "server/models"

type Repository interface {
	Create(port models.Port) (models.Port, error)
	FindByID(ID int) (models.Port, error)
	Edit(port models.Port) (models.Port, error)
	Delete(ID int) (models.Port, error)
	FindAll(searchQuery string, offset int, pageSize int) (port []models.Port, totalCount int64)
}
