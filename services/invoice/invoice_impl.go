package services

import (
	"errors"
	"server/models"
	customerRepositories "server/repositories/customer"
	repositories "server/repositories/invoice"
	shipperRepositories "server/repositories/shipper"
	"server/requests"

	"gorm.io/gorm"
)

type service struct {
	repository         repositories.Repository
	customerRepository customerRepositories.Repository
	shipperRepository  shipperRepositories.Repository
}

func NewService(repository repositories.Repository, customerRepository customerRepositories.Repository, shipperRepository shipperRepositories.Repository) *service {
	return &service{repository, customerRepository, shipperRepository}
}

func (s *service) FindAllNoPagination() ([]models.Invoice, error) {
	invoices, err := s.repository.FindAllNoPagination()
	return invoices, err
}

func (s *service) Create(invoiceRequest requests.CreateInvoiceRequest) (models.Invoice, error) {

	invoiceItems := make([]models.InvoiceItem, len(invoiceRequest.InvoiceItems))
	for i, item := range invoiceRequest.InvoiceItems {
		invoiceItems[i] = models.InvoiceItem{
			ItemName: item.ItemName,
			Currency: item.Currency,
			Price:    item.Price,
			Kurs:     *item.Kurs,
			Quantity: item.Quantity,
		}
	}

	invoice := models.Invoice{
		Category:      invoiceRequest.Category,
		InvoiceNumber: invoiceRequest.InvoiceNumber,
		Type:          invoiceRequest.Type,
		CustomerID:    invoiceRequest.CustomerID,
		ConsigneeID:   invoiceRequest.ConsigneeID,
		ShipperID:     invoiceRequest.ShipperID,
		Service:       invoiceRequest.Service,
		BLAWB:         invoiceRequest.BLAWB,
		AJU:           invoiceRequest.AJU,
		POL:           invoiceRequest.POL,
		POD:           invoiceRequest.POD,
		ShippingMarks: invoiceRequest.ShippingMarks,
		InvoiceDate:   invoiceRequest.InvoiceDate.Time,
		Status:        invoiceRequest.Status,
		InvoiceItems:  invoiceItems,
	}

	newInvoice, err := s.repository.Create(invoice)
	return newInvoice, err
}

func (s *service) FindByID(ID int) (models.Invoice, error) {
	invoice, err := s.repository.FindByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Invoice{}, nil
	}

	return invoice, err
}

func (s *service) Edit(ID int, invoiceRequest requests.EditInvoiceRequest) (models.Invoice, error) {
	invoice, err := s.repository.FindByID(ID)
	if err != nil {
		return models.Invoice{}, err // Handle not found case
	}
	if invoiceRequest.Category != "" {
		invoice.Category = invoiceRequest.Category
	}
	if invoiceRequest.InvoiceNumber != "" {
		invoice.InvoiceNumber = invoiceRequest.InvoiceNumber
	}
	if invoiceRequest.Type != "" {
		invoice.Type = invoiceRequest.Type
	}
	if invoiceRequest.Service != "" {
		invoice.Service = invoiceRequest.Service
	}
	if invoiceRequest.BLAWB != "" {
		invoice.BLAWB = invoiceRequest.BLAWB
	}
	if invoiceRequest.AJU != "" {
		invoice.AJU = invoiceRequest.AJU
	}
	if invoiceRequest.POL != "" {
		invoice.POL = invoiceRequest.POL
	}
	if invoiceRequest.POD != "" {
		invoice.POD = invoiceRequest.POD
	}
	if invoiceRequest.ShippingMarks != "" {
		invoice.ShippingMarks = invoiceRequest.ShippingMarks
	}
	if invoiceRequest.Status != "" {
		invoice.Status = invoiceRequest.Status
	}
	// Check if InvoiceDate is not zero (not the zero value for time.Time)
	if !invoiceRequest.InvoiceDate.IsZero() {
		invoice.InvoiceDate = invoiceRequest.InvoiceDate.Time
	}
	if invoiceRequest.CustomerID != 0 {
		customer, err := s.customerRepository.FindByID(invoiceRequest.CustomerID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.Invoice{}, errors.New("customer not found")
			}
			return models.Invoice{}, err
		}
		invoice.Customer = customer
		invoice.CustomerID = invoiceRequest.CustomerID
	}
	if invoiceRequest.ConsigneeID != 0 {
		consignee, err := s.shipperRepository.FindByID(invoiceRequest.ConsigneeID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.Invoice{}, errors.New("consignee not found")
			}
			return models.Invoice{}, err
		}
		invoice.Consignee = consignee
		invoice.ConsigneeID = invoiceRequest.ConsigneeID
	}
	if invoiceRequest.ShipperID != 0 {
		shipper, err := s.customerRepository.FindByID(invoiceRequest.ShipperID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.Invoice{}, errors.New("shipper not found")
			}
			return models.Invoice{}, err
		}
		invoice.Shipper = shipper
		invoice.ShipperID = invoiceRequest.ShipperID
	}

	if len(invoiceRequest.InvoiceItems) > 0 {
		var jsonItems models.JSONInvoiceItems //
		for _, item := range invoiceRequest.InvoiceItems {
			item := models.InvoiceItem{
				ItemName: item.ItemName,
				Currency: item.Currency,
				Price:    item.Price,
				Kurs:     *item.Kurs,
				Quantity: item.Quantity,
			}
			jsonItems = append(jsonItems, item)
		}
		invoice.InvoiceItems = jsonItems
	}

	updatedInvoice, err := s.repository.Edit(invoice)
	return updatedInvoice, err
}

func (s *service) Delete(ID int) (models.Invoice, error) {
	invoice, err := s.repository.Delete(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Invoice{}, nil
	}

	return invoice, err
}

func (s *service) FindAll(searchQuery string, page int) ([]models.Invoice, int64, int, int, int) {
	if page < 1 {
		return []models.Invoice{}, 0, 0, 0, 0 // Handle invalid page
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	invoice, totalCount := s.repository.FindAll(searchQuery, offset, pageSize)

	firstRow := offset + 1
	lastRow := offset + len(invoice)
	if len(invoice) == 0 {
		firstRow = 0
		lastRow = 0
	}
	totalPages := (int(totalCount) + pageSize - 1) / pageSize

	return invoice, totalCount, firstRow, lastRow, totalPages
}
