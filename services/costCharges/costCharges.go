package services

import (
	"server/models"
	"server/requests"
)

type Service interface {
	Create(costCharges requests.CreateCostChargesRequest) (models.CostCharges, error)
	FindByID(ID int) (models.CostCharges, error)
	Edit(ID int, costCharges requests.EditCostChargesRequest) (models.CostCharges, error)
	Delete(ID int) (models.CostCharges, error)
	FindAll(searchQuery string, page int) ([]models.CostCharges, int64, int, int, int)
}
