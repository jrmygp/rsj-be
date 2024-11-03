package services

import (
	"server/models"
	"server/requests"
)

type Service interface {
	Create(quotation requests.CreateQuotationRequest) (models.Quotation, error)
}
