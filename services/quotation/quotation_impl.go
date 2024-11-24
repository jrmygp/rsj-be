package services

import (
	"errors"
	"server/models"
	customerRepositories "server/repositories/customer"
	portRepositories "server/repositories/port"
	repositories "server/repositories/quotation"
	userRepositories "server/repositories/user"
	"server/requests"

	"gorm.io/gorm"
)

type service struct {
	repository         repositories.Repository
	customerRepository customerRepositories.Repository
	userRepository     userRepositories.Repository
	portRepository     portRepositories.Repository
}

func NewService(repository repositories.Repository, customerRepository customerRepositories.Repository, userRepository userRepositories.Repository, portRepository portRepositories.Repository) *service {
	return &service{repository, customerRepository, userRepository, portRepository}

}

func (s *service) FindAllNoPagination() ([]models.Quotation, error) {
	quotations, err := s.repository.FindAllNoPagination()
	return quotations, err
}

func (s *service) Create(quotationRequest requests.CreateQuotationRequest) (models.Quotation, error) {

	// Convert `quotationRequest.ListCharges` (type []requests.Charge) to []models.Charge
	listCharges := make([]models.Charge, len(quotationRequest.ListCharges))
	for i, reqCharge := range quotationRequest.ListCharges {
		listCharges[i] = models.Charge{
			ItemID:   reqCharge.ItemID,
			ItemName: reqCharge.ItemName,
			Currency: reqCharge.Currency,
			Price:    reqCharge.Price,
			Quantity: reqCharge.Quantity,
			Unit:     reqCharge.Unit,
			Note:     reqCharge.Note,
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
		PaymentTerm:       quotationRequest.PaymentTerm,
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
	if quotationRequest.Note != "" {
		quotation.Note = quotationRequest.Note
	}
	if quotationRequest.PaymentTerm != "" {
		quotation.PaymentTerm = quotationRequest.PaymentTerm
	}
	if quotationRequest.SalesID != 0 {
		sales, err := s.userRepository.FindByID(quotationRequest.SalesID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.Quotation{}, errors.New("user not found")
			}
			return models.Quotation{}, err
		}
		quotation.Sales = sales
		quotation.SalesID = quotationRequest.SalesID
	}
	if quotationRequest.CustomerID != 0 {
		customer, err := s.customerRepository.FindByID(quotationRequest.CustomerID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.Quotation{}, errors.New("customer not found")
			}
			return models.Quotation{}, err
		}
		quotation.Customer = customer
		quotation.CustomerID = quotationRequest.CustomerID

	}
	if quotationRequest.PortOfLoadingID != 0 {
		port, err := s.portRepository.FindByID(quotationRequest.PortOfLoadingID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.Quotation{}, errors.New("port not found")
			}
			return models.Quotation{}, err
		}
		quotation.PortOfLoading = port
		quotation.PortOfLoadingID = quotationRequest.PortOfLoadingID
	}
	if quotationRequest.PortOfDischargeID != 0 {
		port, err := s.portRepository.FindByID(quotationRequest.PortOfDischargeID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.Quotation{}, errors.New("port not found")
			}
			return models.Quotation{}, err
		}
		quotation.PortOfLoading = port
		quotation.PortOfDischargeID = quotationRequest.PortOfDischargeID
	}

	// Convert ListCharges to JSONCharges
	if len(quotationRequest.ListCharges) > 0 {
		var jsonCharges models.JSONCharges //
		for _, charge := range quotationRequest.ListCharges {
			jsonCharge := models.Charge{
				ItemID:   charge.ItemID,
				ItemName: charge.ItemName,
				Price:    charge.Price,
				Currency: charge.Currency,
				Quantity: charge.Quantity,
				Unit:     charge.Unit,
				Note:     charge.Note,
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
