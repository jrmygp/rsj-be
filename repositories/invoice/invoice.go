package repositories

import "server/models"

type Repository interface {
	FindAllNoPagination() ([]models.Invoice, error)
	Create(invoice models.Invoice) (models.Invoice, error)
	FindByID(ID int) (models.Invoice, error)
	Edit(invoice models.Invoice) (models.Invoice, error)
	Delete(ID int) (models.Invoice, error)
	FindAll(searchQuery string, offset int, pageSize int, customerID int, category string) (invoice []models.Invoice, totalCount int64)

	FindAllDoorToDoorNoPagination() ([]models.DoorToDoorInvoice, error)
	CreateDoorToDoor(invoice models.DoorToDoorInvoice) (models.DoorToDoorInvoice, error)
	FindDoorToDoorByID(ID int) (models.DoorToDoorInvoice, error)
	EditDoorToDoor(invoice models.DoorToDoorInvoice) (models.DoorToDoorInvoice, error)
	DeleteDoorToDoor(ID int) (models.DoorToDoorInvoice, error)
	FindAllDoorToDoor(searchQuery string, offset int, pageSize int, customerID int) (invoice []models.DoorToDoorInvoice, totalCount int64)
}
