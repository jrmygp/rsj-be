package services

import (
	"server/models"
	"server/requests"
)

type Service interface {
	Create(port requests.CreatePortRequest) (models.Port, error)
	FindByID(ID int) (models.Port, error)
	Edit(ID int, port requests.EditPortRequest) (models.Port, error)
	Delete(ID int) (models.Port, error)
	FindAll(searchQuery string, page int) ([]models.Port, int64, int, int, int)
}
