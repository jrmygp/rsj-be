package services

import (
	"server/models"
	"server/requests"
)

type Service interface {
	FindAllNoPagination() ([]models.Invoice, error)
	Create(quotation requests.CreateInvoiceRequest) (models.Invoice, error)
	FindByID(ID int) (models.Invoice, error)
	Edit(ID int, quotation requests.EditInvoiceRequest) (models.Invoice, error)
	Delete(ID int) (models.Invoice, error)
	FindAll(searchQuery string, page int) ([]models.Invoice, int64, int, int, int)
}
