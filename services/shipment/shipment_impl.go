package service

import (
	"server/models"
	invoiceRepositories "server/repositories/invoice"
	quotationRepositories "server/repositories/quotation"
	repositories "server/repositories/shipment"
	"server/requests"
)

type service struct {
	repository            repositories.Repository
	quotationRepositories quotationRepositories.Repository
	invoiceRepositories   invoiceRepositories.Repository
}

func NewService(repository repositories.Repository, quotationRepository quotationRepositories.Repository, invoiceRepository invoiceRepositories.Repository) *service {
	return &service{repository, quotationRepository, invoiceRepository}
}

func (s *service) FindAll(searchQuery string, page int) ([]models.Shipment, int64, int, int, int) {
	if page < 1 {
		return []models.Shipment{}, 0, 0, 0, 0 // Handle invalid page
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	shipment, totalCount := s.repository.FindAll(searchQuery, offset, pageSize)

	firstRow := offset + 1
	lastRow := offset + len(shipment)
	if len(shipment) == 0 {
		firstRow = 0
		lastRow = 0
	}
	totalPages := (int(totalCount) + pageSize - 1) / pageSize

	return shipment, totalCount, firstRow, lastRow, totalPages
}

func (s *service) Create(shipmentRequest requests.ShipmentRequest) (models.Shipment, error) {
	var quotations []models.Quotation
	var invoiceExports []models.InvoiceExport
	var invoiceImports []models.InvoiceImport
	var invoiceD2Ds []models.DoorToDoorInvoice

	// Fetch Quotations in a single query
	if len(shipmentRequest.Quotations) > 0 {
		var err error
		quotations, err = s.quotationRepositories.FindByIDs(shipmentRequest.Quotations)
		if err != nil {
			return models.Shipment{}, err
		}
	}

	// Fetch Export invoices in a single query
	if len(shipmentRequest.InvoiceExports) > 0 {
		var err error
		invoiceExports, err = s.invoiceRepositories.FindExportByIDs(shipmentRequest.InvoiceExports)
		if err != nil {
			return models.Shipment{}, err
		}
	}

	// Fetch Import invoices in a single query
	if len(shipmentRequest.Quotations) > 0 {
		var err error
		invoiceImports, err = s.invoiceRepositories.FindImportByIDs(shipmentRequest.InvoiceImports)
		if err != nil {
			return models.Shipment{}, err
		}
	}

	// Fetch D2D invoices in a single query
	if len(shipmentRequest.Quotations) > 0 {
		var err error
		invoiceD2Ds, err = s.invoiceRepositories.FindDoorToDoorByIDs(shipmentRequest.InvoiceDoorToDoors)
		if err != nil {
			return models.Shipment{}, err
		}
	}

	shipment := models.Shipment{
		ShipmentNumber:     shipmentRequest.ShipmentNumber,
		Quotations:         quotations,
		InvoiceExports:     invoiceExports,
		InvoiceImports:     invoiceImports,
		InvoiceDoorToDoors: invoiceD2Ds,
	}

	newShipment, err := s.repository.Create(shipment)
	return newShipment, err
}
