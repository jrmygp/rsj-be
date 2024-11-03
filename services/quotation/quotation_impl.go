package services

import (
	"errors"
	"server/models"
	repositories "server/repositories/quotation"
	"server/requests"

	"gorm.io/gorm"
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
			RatioToIDR: reqCharge.RatioIDR,
			Quantity:   reqCharge.Quantity,
			Unit:       reqCharge.Unit,
			Note:       reqCharge.Note,
		}
	}

	quotation := models.Quotation{
		QuotationNumber:   quotationRequest.QuotationNumber,
		RateValidity:      quotationRequest.RateValidity.Time,
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

func (s *service) FindByID(ID int) (models.Quotation, error) {
	quotation, err := s.repository.FindByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Quotation{}, nil
	}

	return quotation, err
}

func (s *service) Edit(ID int, quotationRequest requests.EditQuotationRequest) (models.Quotation, error) {
	quotation, err := s.repository.FindByID(ID)
	if err != nil {
		return models.Quotation{}, err // Handle not found case
	}

	if quotationRequest.QuotationNumber != "" {
		quotation.QuotationNumber = quotationRequest.QuotationNumber
	}

	// Check if RateValidity is not zero (not the zero value for time.Time)
	if !quotationRequest.RateValidity.IsZero() {
		quotation.RateValidity = quotationRequest.RateValidity
	}

	if quotationRequest.ShippingTerm != "" {
		quotation.ShippingTerm = quotationRequest.ShippingTerm
	}
	if quotationRequest.Service != "" {
		quotation.Service = quotationRequest.Service
	}
	if quotationRequest.Status != "" {
		quotation.Status = quotationRequest.Status
	}
	if quotationRequest.Commodity != "" {
		quotation.Commodity = quotationRequest.Commodity
	}
	if quotationRequest.Weight != 0 {
		quotation.Weight = quotationRequest.Weight
	}
	if quotationRequest.Volume != 0 {
		quotation.Volume = quotationRequest.Volume
	}
	if quotationRequest.SalesID != 0 {
		quotation.SalesID = quotationRequest.SalesID
	}
	if quotationRequest.CustomerID != 0 {
		quotation.CustomerID = quotationRequest.CustomerID
	}
	if quotationRequest.PortOfLoadingID != 0 {
		quotation.PortOfLoadingID = quotationRequest.PortOfLoadingID
	}
	if quotationRequest.PortOfDischargeID != 0 {
		quotation.PortOfDischargeID = quotationRequest.PortOfDischargeID
	}

	// Convert ListCharges to JSONCharges
	if len(quotationRequest.ListCharges) > 0 {
		var jsonCharges models.JSONCharges //
		for _, charge := range quotationRequest.ListCharges {
			jsonCharge := models.Charge{
				ItemName:   charge.ItemName,
				Price:      charge.Price,
				Currency:   charge.Currency,
				RatioToIDR: charge.RatioIDR,
				Quantity:   charge.Quantity,
				Unit:       charge.Unit,
				Note:       charge.Note,
			}
			jsonCharges = append(jsonCharges, jsonCharge)
		}
		quotation.ListCharges = jsonCharges
	}

	updatedQuotation, err := s.repository.Edit(quotation)
	return updatedQuotation, err
}

func (s *service) Delete(ID int) (models.Quotation, error) {
	quotation, err := s.repository.Delete(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Quotation{}, nil
	}

	return quotation, err
}

func (s *service) FindAll(searchQuery string, page int) ([]models.Quotation, int64, int, int, int) {
	if page < 1 {
		return []models.Quotation{}, 0, 0, 0, 0 // Handle invalid page
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	quotation, totalCount := s.repository.FindAll(searchQuery, offset, pageSize)

	firstRow := offset + 1
	lastRow := offset + len(quotation)
	if len(quotation) == 0 {
		firstRow = 0
		lastRow = 0
	}
	totalPages := (int(totalCount) + pageSize - 1) / pageSize

	return quotation, totalCount, firstRow, lastRow, totalPages
}
