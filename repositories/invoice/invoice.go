package repositories

import "server/models"

type Repository interface {
	FindAllNoPagination() ([]models.Invoice, error)
	Create(invoice models.Invoice) (models.Invoice, error)
	FindByID(ID int) (models.Invoice, error)
	Edit(invoice models.Invoice) (models.Invoice, error)
	Delete(ID int) (models.Invoice, error)
	FindAll(searchQuery string, offset int, pageSize int) (invoice []models.Invoice, totalCount int64)
}
