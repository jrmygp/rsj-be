package repositories

import "server/models"

type Repository interface {
	FindAllExportNoPagination() ([]models.InvoiceExport, error)
	CreateExport(invoice models.InvoiceExport) (models.InvoiceExport, error)
	FindExportByID(ID int) (models.InvoiceExport, error)
	EditExport(invoice models.InvoiceExport) (models.InvoiceExport, error)
	DeleteExport(ID int) (models.InvoiceExport, error)
	FindAllExport(searchQuery string, offset int, pageSize int, customerID int) (invoice []models.InvoiceExport, totalCount int64)
	FindExportByIDs(IDs []int) ([]models.InvoiceExport, error)

	FindAllImportNoPagination() ([]models.InvoiceImport, error)
	CreateImport(invoice models.InvoiceImport) (models.InvoiceImport, error)
	FindImportByID(ID int) (models.InvoiceImport, error)
	EditImport(invoice models.InvoiceImport) (models.InvoiceImport, error)
	DeleteImport(ID int) (models.InvoiceImport, error)
	FindAllImport(searchQuery string, offset int, pageSize int, customerID int) (invoice []models.InvoiceImport, totalCount int64)
	FindImportByIDs(IDs []int) ([]models.InvoiceImport, error)

	FindAllDoorToDoorNoPagination() ([]models.DoorToDoorInvoice, error)
	CreateDoorToDoor(invoice models.DoorToDoorInvoice) (models.DoorToDoorInvoice, error)
	FindDoorToDoorByID(ID int) (models.DoorToDoorInvoice, error)
	EditDoorToDoor(invoice models.DoorToDoorInvoice) (models.DoorToDoorInvoice, error)
	DeleteDoorToDoor(ID int) (models.DoorToDoorInvoice, error)
	FindAllDoorToDoor(searchQuery string, offset int, pageSize int, customerID int) (invoice []models.DoorToDoorInvoice, totalCount int64)
	FindDoorToDoorByIDs(IDs []int) ([]models.DoorToDoorInvoice, error)
}
