package repositories

import "server/models"

type Repository interface {
	FindAll(searchQuery string, offset int, pageSize int) (shipment []models.Shipment, totalCount int64)
	Create(shipment models.Shipment) (models.Shipment, error)
}
