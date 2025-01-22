package services

import (
	"server/models"
	"server/requests"
)

type Service interface {
	FindAllExportNoPagination() ([]models.InvoiceExport, error)
	CreateExport(invoice requests.CreateInvoiceRequest) (models.InvoiceExport, error)
	FindExportByID(ID int) (models.InvoiceExport, error)
	EditExport(ID int, invoice requests.EditInvoiceRequest, userRoleID int) (models.InvoiceExport, error)
	DeleteExport(ID int) (models.InvoiceExport, error)
	FindAllExport(searchQuery string, page int, filter requests.InvoiceFilterRequest) ([]models.InvoiceExport, int64, int, int, int)

	FindAllImportNoPagination() ([]models.InvoiceImport, error)
	CreateImport(invoice requests.CreateInvoiceRequest) (models.InvoiceImport, error)
	FindImportByID(ID int) (models.InvoiceImport, error)
	EditImport(ID int, invoice requests.EditInvoiceRequest, userRoleID int) (models.InvoiceImport, error)
	DeleteImport(ID int) (models.InvoiceImport, error)
	FindAllImport(searchQuery string, page int, filter requests.InvoiceFilterRequest) ([]models.InvoiceImport, int64, int, int, int)

	FindAllDoorToDoorNoPagination() ([]models.DoorToDoorInvoice, error)
	CreateDoorToDoor(invoice requests.CreateDoorToDoorRequest) (models.DoorToDoorInvoice, error)
	FindDoorToDoorByID(ID int) (models.DoorToDoorInvoice, error)
	EditDoorToDoor(ID int, invoice requests.EditDoorToDoorRequest) (models.DoorToDoorInvoice, error)
	DeleteDoorToDoor(ID int) (models.DoorToDoorInvoice, error)
	FindAllDoorToDoor(searchQuery string, page int, filter requests.InvoiceFilterRequest) ([]models.DoorToDoorInvoice, int64, int, int, int)
}
