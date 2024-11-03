package services

import (
	"server/helper"
	"server/models"
	repositories "server/repositories/quotation"
	"server/requests"
)

type service struct {
	repository repositories.Repository
}

func NewService(repository repositories.Repository) *service {
	return &service{repository}
}

func (s *service) Create(quotationRequest requests.CreateQuotationRequest) (models.Quotation, error) {

	// Convert `quotationRequest.ListCharges` (type []requests.Charge) to []models.Charge
	listCharges := make([]models.Charge, len(quotationRequest.ListCharges))
	for i, reqCharge := range quotationRequest.ListCharges {
		listCharges[i] = models.Charge{
			ItemName:   reqCharge.ItemName,
			Price:      reqCharge.Price,
			RatioToIDR: helper.ConvertToNullableFloat64(reqCharge.RatioIDR),
			Quantity:   reqCharge.Quantity,
			Unit:       reqCharge.Unit,
			Note:       helper.DereferenceString(reqCharge.Note),
		}
	}

	quotation := models.Quotation{
		QuotationNumber:   quotationRequest.QuotationNumber,
		RateValidity:      quotationRequest.RateValidity,
		ShippingTerm:      quotationRequest.ShippingTerm,
		Service:           quotationRequest.Service,
		Status:            quotationRequest.Status,
		Commodity:         quotationRequest.Commodity,
		Weight:            quotationRequest.Weight,
		Volume:            quotationRequest.Volume,
		Note:              quotationRequest.Note,
		SalesID:           quotationRequest.SalesID,
		CustomerID:        quotationRequest.CustomerID,
		PortOfLoadingID:   quotationRequest.PortOfLoadingID,
		PortOfDischargeID: quotationRequest.PortOfDischargeID,
		ListCharges:       listCharges,
	}

	newQuotation, err := s.repository.Create(quotation)
	return newQuotation, err
}
