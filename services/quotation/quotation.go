package services

import (
	"server/models"
	"server/requests"
)

type Service interface {
	Create(quotation requests.CreateQuotationRequest) (models.Quotation, error)
	FindByID(ID int) (models.Quotation, error)
	Edit(ID int, quotation requests.EditQuotationRequest) (models.Quotation, error)
	Delete(ID int) (models.Quotation, error)
	FindAll(searchQuery string, page int) ([]models.Quotation, int64, int, int, int)
}
