package repositories

import "server/models"

type Repository interface {
	Create(costCharges models.CostCharges) (models.CostCharges, error)
	FindByID(ID int) (models.CostCharges, error)
	Edit(costCharges models.CostCharges) (models.CostCharges, error)
	Delete(ID int) (models.CostCharges, error)
	FindAll(searchQuery string, offset int, pageSize int) (costCharges []models.CostCharges, totalCount int64)
}
