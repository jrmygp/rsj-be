package services

import (
	"server/models"
	"server/requests"
)

type Service interface {
	FindAllNoPagination() ([]models.Shipper, error)
	Create(port requests.CreateShipperRequest) (models.Shipper, error)
	FindByID(ID int) (models.Shipper, error)
	Edit(ID int, port requests.EditShipperRequest) (models.Shipper, error)
	Delete(ID int) (models.Shipper, error)
	FindAll(searchQuery string, page int) ([]models.Shipper, int64, int, int, int)
}
