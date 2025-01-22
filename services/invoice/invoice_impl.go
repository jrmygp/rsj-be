package services

import (
	"errors"
	"server/models"
	customerRepositories "server/repositories/customer"
	repositories "server/repositories/invoice"
	portRepositories "server/repositories/port"
	shipperRepositories "server/repositories/shipper"
	"server/requests"

	"gorm.io/gorm"
)

type service struct {
	repository         repositories.Repository
	customerRepository customerRepositories.Repository
	shipperRepository  shipperRepositories.Repository
	portRepository     portRepositories.Repository
}

func NewService(repository repositories.Repository, customerRepository customerRepositories.Repository, shipperRepository shipperRepositories.Repository, portRepository portRepositories.Repository) *service {
	return &service{repository, customerRepository, shipperRepository, portRepository}
}

// Export implementations
func (s *service) FindAllExportNoPagination() ([]models.InvoiceExport, error) {
	invoices, err := s.repository.FindAllExportNoPagination()
	return invoices, err
}

func (s *service) CreateExport(invoiceRequest requests.CreateInvoiceRequest) (models.InvoiceExport, error) {

	invoiceItems := make([]models.InvoiceItem, len(invoiceRequest.InvoiceItems))
	for i, item := range invoiceRequest.InvoiceItems {
		invoiceItems[i] = models.InvoiceItem{
			ItemName: item.ItemName,
			Currency: item.Currency,
			Price:    item.Price,
			Kurs:     item.Kurs,
			Quantity: item.Quantity,
			Unit:     item.Unit,
		}
	}

	invoice := models.InvoiceExport{
		InvoiceNumber:     invoiceRequest.InvoiceNumber,
		Type:              invoiceRequest.Type,
		CustomerID:        invoiceRequest.CustomerID,
		ConsigneeID:       invoiceRequest.ConsigneeID,
		ShipperID:         invoiceRequest.ShipperID,
		Service:           invoiceRequest.Service,
		BLAWB:             invoiceRequest.BLAWB,
		AJU:               invoiceRequest.AJU,
		PortOfLoadingID:   invoiceRequest.PortOfLoadingID,
		PortOfDischargeID: invoiceRequest.PortOfDischargeID,
		ShippingMarks:     invoiceRequest.ShippingMarks,
		InvoiceDate:       invoiceRequest.InvoiceDate.Time,
		Status:            invoiceRequest.Status,
		InvoiceItems:      invoiceItems,
	}

	newInvoice, err := s.repository.CreateExport(invoice)
	return newInvoice, err
}

func (s *service) FindExportByID(ID int) (models.InvoiceExport, error) {
	invoice, err := s.repository.FindExportByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.InvoiceExport{}, nil
	}

	return invoice, err
}

func (s *service) EditExport(ID int, invoiceRequest requests.EditInvoiceRequest, userRoleID int) (models.InvoiceExport, error) {
	invoice, err := s.repository.FindExportByID(ID)
	if err != nil {
		return models.InvoiceExport{}, err // Handle not found case
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
	if invoiceRequest.ShippingMarks != "" {
		invoice.ShippingMarks = invoiceRequest.ShippingMarks
	}
	if invoiceRequest.Status != "" && invoiceRequest.Status != invoice.Status {
		if userRoleID != 1 {
			return models.InvoiceExport{}, errors.New("you have no access to change status")
		}
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
				return models.InvoiceExport{}, errors.New("customer not found")
			}
			return models.InvoiceExport{}, err
		}
		invoice.Customer = customer
		invoice.CustomerID = invoiceRequest.CustomerID
	}
	if invoiceRequest.ConsigneeID != 0 {
		consignee, err := s.shipperRepository.FindByID(invoiceRequest.ConsigneeID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.InvoiceExport{}, errors.New("consignee not found")
			}
			return models.InvoiceExport{}, err
		}
		invoice.Consignee = consignee
		invoice.ConsigneeID = invoiceRequest.ConsigneeID
	}
	if invoiceRequest.ShipperID != 0 {
		shipper, err := s.customerRepository.FindByID(invoiceRequest.ShipperID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.InvoiceExport{}, errors.New("shipper not found")
			}
			return models.InvoiceExport{}, err
		}
		invoice.Shipper = shipper
		invoice.ShipperID = invoiceRequest.ShipperID
	}
	if invoiceRequest.PortOfLoadingID != 0 {
		port, err := s.portRepository.FindByID(invoiceRequest.PortOfLoadingID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.InvoiceExport{}, errors.New("port not found")
			}
			return models.InvoiceExport{}, err
		}
		invoice.PortOfLoading = port
		invoice.PortOfLoadingID = invoiceRequest.PortOfLoadingID
	}
	if invoiceRequest.PortOfDischargeID != 0 {
		port, err := s.portRepository.FindByID(invoiceRequest.PortOfDischargeID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.InvoiceExport{}, errors.New("port not found")
			}
			return models.InvoiceExport{}, err
		}
		invoice.PortOfDischarge = port
		invoice.PortOfDischargeID = invoiceRequest.PortOfDischargeID
	}

	if len(invoiceRequest.InvoiceItems) > 0 {
		var jsonItems models.JSONInvoiceItems //
		for _, item := range invoiceRequest.InvoiceItems {
			item := models.InvoiceItem{
				ItemName: item.ItemName,
				Currency: item.Currency,
				Price:    item.Price,
				Kurs:     item.Kurs,
				Quantity: item.Quantity,
				Unit:     item.Unit,
			}
			jsonItems = append(jsonItems, item)
		}
		invoice.InvoiceItems = jsonItems
	}

	updatedInvoice, err := s.repository.EditExport(invoice)
	return updatedInvoice, err
}

func (s *service) DeleteExport(ID int) (models.InvoiceExport, error) {
	invoice, err := s.repository.DeleteExport(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.InvoiceExport{}, nil
	}

	return invoice, err
}

func (s *service) FindAllExport(searchQuery string, page int, filter requests.InvoiceFilterRequest) ([]models.InvoiceExport, int64, int, int, int) {
	if page < 1 {
		return []models.InvoiceExport{}, 0, 0, 0, 0 // Handle invalid page
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	invoice, totalCount := s.repository.FindAllExport(searchQuery, offset, pageSize, filter.CustomerID)

	firstRow := offset + 1
	lastRow := offset + len(invoice)
	if len(invoice) == 0 {
		firstRow = 0
		lastRow = 0
	}
	totalPages := (int(totalCount) + pageSize - 1) / pageSize

	return invoice, totalCount, firstRow, lastRow, totalPages
}

// Import implementations
func (s *service) FindAllImportNoPagination() ([]models.InvoiceImport, error) {
	invoices, err := s.repository.FindAllImportNoPagination()
	return invoices, err
}

func (s *service) CreateImport(invoiceRequest requests.CreateInvoiceRequest) (models.InvoiceImport, error) {

	invoiceItems := make([]models.InvoiceItem, len(invoiceRequest.InvoiceItems))
	for i, item := range invoiceRequest.InvoiceItems {
		invoiceItems[i] = models.InvoiceItem{
			ItemName: item.ItemName,
			Currency: item.Currency,
			Price:    item.Price,
			Kurs:     item.Kurs,
			Quantity: item.Quantity,
			Unit:     item.Unit,
		}
	}

	invoice := models.InvoiceImport{
		InvoiceNumber:     invoiceRequest.InvoiceNumber,
		Type:              invoiceRequest.Type,
		CustomerID:        invoiceRequest.CustomerID,
		ConsigneeID:       invoiceRequest.ConsigneeID,
		ShipperID:         invoiceRequest.ShipperID,
		Service:           invoiceRequest.Service,
		BLAWB:             invoiceRequest.BLAWB,
		AJU:               invoiceRequest.AJU,
		PortOfLoadingID:   invoiceRequest.PortOfLoadingID,
		PortOfDischargeID: invoiceRequest.PortOfDischargeID,
		ShippingMarks:     invoiceRequest.ShippingMarks,
		InvoiceDate:       invoiceRequest.InvoiceDate.Time,
		Status:            invoiceRequest.Status,
		InvoiceItems:      invoiceItems,
	}

	newInvoice, err := s.repository.CreateImport(invoice)
	return newInvoice, err
}

func (s *service) FindImportByID(ID int) (models.InvoiceImport, error) {
	invoice, err := s.repository.FindImportByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.InvoiceImport{}, nil
	}

	return invoice, err
}

func (s *service) EditImport(ID int, invoiceRequest requests.EditInvoiceRequest, userRoleID int) (models.InvoiceImport, error) {
	invoice, err := s.repository.FindImportByID(ID)
	if err != nil {
		return models.InvoiceImport{}, err // Handle not found case
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
	if invoiceRequest.ShippingMarks != "" {
		invoice.ShippingMarks = invoiceRequest.ShippingMarks
	}
	if invoiceRequest.Status != "" && invoiceRequest.Status != invoice.Status {
		if userRoleID != 1 {
			return models.InvoiceImport{}, errors.New("you have no access to change status")
		}
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
				return models.InvoiceImport{}, errors.New("customer not found")
			}
			return models.InvoiceImport{}, err
		}
		invoice.Customer = customer
		invoice.CustomerID = invoiceRequest.CustomerID
	}
	if invoiceRequest.ConsigneeID != 0 {
		consignee, err := s.customerRepository.FindByID(invoiceRequest.ConsigneeID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.InvoiceImport{}, errors.New("consignee not found")
			}
			return models.InvoiceImport{}, err
		}
		invoice.Consignee = consignee
		invoice.ConsigneeID = invoiceRequest.ConsigneeID
	}
	if invoiceRequest.ShipperID != 0 {
		shipper, err := s.shipperRepository.FindByID(invoiceRequest.ShipperID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.InvoiceImport{}, errors.New("shipper not found")
			}
			return models.InvoiceImport{}, err
		}
		invoice.Shipper = shipper
		invoice.ShipperID = invoiceRequest.ShipperID
	}
	if invoiceRequest.PortOfLoadingID != 0 {
		port, err := s.portRepository.FindByID(invoiceRequest.PortOfLoadingID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.InvoiceImport{}, errors.New("port not found")
			}
			return models.InvoiceImport{}, err
		}
		invoice.PortOfLoading = port
		invoice.PortOfLoadingID = invoiceRequest.PortOfLoadingID
	}
	if invoiceRequest.PortOfDischargeID != 0 {
		port, err := s.portRepository.FindByID(invoiceRequest.PortOfDischargeID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.InvoiceImport{}, errors.New("port not found")
			}
			return models.InvoiceImport{}, err
		}
		invoice.PortOfDischarge = port
		invoice.PortOfDischargeID = invoiceRequest.PortOfDischargeID
	}

	if len(invoiceRequest.InvoiceItems) > 0 {
		var jsonItems models.JSONInvoiceItems //
		for _, item := range invoiceRequest.InvoiceItems {
			item := models.InvoiceItem{
				ItemName: item.ItemName,
				Currency: item.Currency,
				Price:    item.Price,
				Kurs:     item.Kurs,
				Quantity: item.Quantity,
				Unit:     item.Unit,
			}
			jsonItems = append(jsonItems, item)
		}
		invoice.InvoiceItems = jsonItems
	}

	updatedInvoice, err := s.repository.EditImport(invoice)
	return updatedInvoice, err
}

func (s *service) DeleteImport(ID int) (models.InvoiceImport, error) {
	invoice, err := s.repository.DeleteImport(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.InvoiceImport{}, nil
	}

	return invoice, err
}

func (s *service) FindAllImport(searchQuery string, page int, filter requests.InvoiceFilterRequest) ([]models.InvoiceImport, int64, int, int, int) {
	if page < 1 {
		return []models.InvoiceImport{}, 0, 0, 0, 0 // Handle invalid page
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	invoice, totalCount := s.repository.FindAllImport(searchQuery, offset, pageSize, filter.CustomerID)

	firstRow := offset + 1
	lastRow := offset + len(invoice)
	if len(invoice) == 0 {
		firstRow = 0
		lastRow = 0
	}
	totalPages := (int(totalCount) + pageSize - 1) / pageSize

	return invoice, totalCount, firstRow, lastRow, totalPages
}

// Door to Door implementations
func (s *service) FindAllDoorToDoorNoPagination() ([]models.DoorToDoorInvoice, error) {
	invoices, err := s.repository.FindAllDoorToDoorNoPagination()
	return invoices, err
}

func (s *service) CreateDoorToDoor(invoiceRequest requests.CreateDoorToDoorRequest) (models.DoorToDoorInvoice, error) {
	invoiceItems := make([]models.InvoiceD2DItem, len(invoiceRequest.InvoiceItems))
	for i, item := range invoiceRequest.InvoiceItems {
		invoiceItems[i] = models.InvoiceD2DItem{
			ItemName: item.ItemName,
			Currency: item.Currency,
			Price:    item.Price,
			Kurs:     item.Kurs,
			Quantity: item.Quantity,
		}
	}

	invoice := models.DoorToDoorInvoice{
		InvoiceNumber:     invoiceRequest.InvoiceNumber,
		Type:              invoiceRequest.Type,
		CustomerID:        invoiceRequest.CustomerID,
		ConsigneeID:       invoiceRequest.ConsigneeID,
		ShipperID:         invoiceRequest.ShipperID,
		Service:           invoiceRequest.Service,
		PortOfLoadingID:   invoiceRequest.PortOfLoadingID,
		PortOfDischargeID: invoiceRequest.PortOfDischargeID,
		ShippingMarks:     invoiceRequest.ShippingMarks,
		InvoiceDate:       invoiceRequest.InvoiceDate.Time,
		Status:            invoiceRequest.Status,
		Quantity:          invoiceRequest.Quantity,
		Weight:            invoiceRequest.Weight,
		Volume:            invoiceRequest.Volume,
		InvoiceItems:      invoiceItems,
	}

	newInvoice, err := s.repository.CreateDoorToDoor(invoice)
	return newInvoice, err
}

func (s *service) FindDoorToDoorByID(ID int) (models.DoorToDoorInvoice, error) {
	invoice, err := s.repository.FindDoorToDoorByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.DoorToDoorInvoice{}, nil
	}

	return invoice, err
}

func (s *service) EditDoorToDoor(ID int, invoiceRequest requests.EditDoorToDoorRequest) (models.DoorToDoorInvoice, error) {
	invoice, err := s.repository.FindDoorToDoorByID(ID)
	if err != nil {
		return models.DoorToDoorInvoice{}, err // Handle not found case
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
	if invoiceRequest.ShippingMarks != "" {
		invoice.ShippingMarks = invoiceRequest.ShippingMarks
	}
	if invoiceRequest.Status != "" {
		invoice.Status = invoiceRequest.Status
	}
	if invoiceRequest.Quantity != "" {
		invoice.Quantity = invoiceRequest.Quantity
	}
	if invoiceRequest.Weight != 0 {
		invoice.Weight = invoiceRequest.Weight
	}
	if invoiceRequest.Volume != 0 {
		invoice.Volume = invoiceRequest.Volume
	}
	// Check if InvoiceDate is not zero (not the zero value for time.Time)
	if !invoiceRequest.InvoiceDate.IsZero() {
		invoice.InvoiceDate = invoiceRequest.InvoiceDate.Time
	}
	if invoiceRequest.CustomerID != 0 {
		customer, err := s.customerRepository.FindByID(invoiceRequest.CustomerID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.DoorToDoorInvoice{}, errors.New("customer not found")
			}
			return models.DoorToDoorInvoice{}, err
		}
		invoice.Customer = customer
		invoice.CustomerID = invoiceRequest.CustomerID
	}
	if invoiceRequest.ConsigneeID != 0 {
		consignee, err := s.customerRepository.FindByID(invoiceRequest.ConsigneeID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.DoorToDoorInvoice{}, errors.New("consignee not found")
			}
			return models.DoorToDoorInvoice{}, err
		}
		invoice.Consignee = consignee
		invoice.ConsigneeID = invoiceRequest.ConsigneeID
	}
	if invoiceRequest.ShipperID != 0 {
		shipper, err := s.shipperRepository.FindByID(invoiceRequest.ShipperID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.DoorToDoorInvoice{}, errors.New("shipper not found")
			}
			return models.DoorToDoorInvoice{}, err
		}
		invoice.Shipper = shipper
		invoice.ShipperID = invoiceRequest.ShipperID
	}
	if invoiceRequest.PortOfLoadingID != 0 {
		port, err := s.portRepository.FindByID(invoiceRequest.PortOfLoadingID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.DoorToDoorInvoice{}, errors.New("port not found")
			}
			return models.DoorToDoorInvoice{}, err
		}
		invoice.PortOfLoading = port
		invoice.PortOfLoadingID = invoiceRequest.PortOfLoadingID
	}
	if invoiceRequest.PortOfDischargeID != 0 {
		port, err := s.portRepository.FindByID(invoiceRequest.PortOfDischargeID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.DoorToDoorInvoice{}, errors.New("port not found")
			}
			return models.DoorToDoorInvoice{}, err
		}
		invoice.PortOfDischarge = port
		invoice.PortOfDischargeID = invoiceRequest.PortOfDischargeID
	}

	if len(invoiceRequest.InvoiceItems) > 0 {
		var jsonItems models.JSONInvoiceD2DItems //
		for _, item := range invoiceRequest.InvoiceItems {
			item := models.InvoiceD2DItem{
				ItemName: item.ItemName,
				Currency: item.Currency,
				Price:    item.Price,
				Kurs:     item.Kurs,
				Quantity: item.Quantity,
			}
			jsonItems = append(jsonItems, item)
		}
		invoice.InvoiceItems = jsonItems
	}

	updatedInvoice, err := s.repository.EditDoorToDoor(invoice)
	return updatedInvoice, err
}

func (s *service) DeleteDoorToDoor(ID int) (models.DoorToDoorInvoice, error) {
	invoice, err := s.repository.DeleteDoorToDoor(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.DoorToDoorInvoice{}, nil
	}

	return invoice, err
}

func (s *service) FindAllDoorToDoor(searchQuery string, page int, filter requests.InvoiceFilterRequest) ([]models.DoorToDoorInvoice, int64, int, int, int) {
	if page < 1 {
		return []models.DoorToDoorInvoice{}, 0, 0, 0, 0 // Handle invalid page
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	invoice, totalCount := s.repository.FindAllDoorToDoor(searchQuery, offset, pageSize, filter.CustomerID)

	firstRow := offset + 1
	lastRow := offset + len(invoice)
	if len(invoice) == 0 {
		firstRow = 0
		lastRow = 0
	}
	totalPages := (int(totalCount) + pageSize - 1) / pageSize

	return invoice, totalCount, firstRow, lastRow, totalPages
}
