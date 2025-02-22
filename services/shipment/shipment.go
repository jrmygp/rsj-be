package service

import (
	"server/models"
	"server/requests"
)

type Service interface {
	FindAll(searchQuery string, page int) ([]models.Shipment, int64, int, int, int)
	Create(invoice requests.ShipmentRequest) (models.Shipment, error)
}
