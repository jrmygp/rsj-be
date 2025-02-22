package repositories

import "server/models"

type Repository interface {
	FindAllNoPagination() ([]models.Quotation, error)
	Create(quotation models.Quotation) (models.Quotation, error)
	FindByID(ID int) (models.Quotation, error)
	Edit(quotation models.Quotation) (models.Quotation, error)
	Delete(ID int) (models.Quotation, error)
	FindAll(searchQuery string, offset int, pageSize int) (quotation []models.Quotation, totalCount int64)
	FindByIDs(IDs []int) ([]models.Quotation, error)
}
