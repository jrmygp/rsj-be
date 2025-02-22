package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

type InvoiceItem struct {
	ItemName string   `json:"itemName"`
	Currency string   `json:"currency"`
	Price    float64  `json:"price"`
	Kurs     *float64 `json:"kurs"`
	Quantity int      `json:"quantity"`
	Unit     string   `json:"unit"`
}

type InvoiceD2DItem struct {
	ItemName string   `json:"itemName"`
	Currency string   `json:"currency"`
	Price    float64  `json:"price"`
	Kurs     *float64 `json:"kurs"`
	Quantity int      `json:"quantity"`
}

type JSONInvoiceItems []InvoiceItem
type JSONInvoiceD2DItems []InvoiceD2DItem

func (jc JSONInvoiceItems) Value() (driver.Value, error) {
	return json.Marshal(jc)
}

func (jc *JSONInvoiceItems) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, jc)
}

func (jc JSONInvoiceD2DItems) Value() (driver.Value, error) {
	return json.Marshal(jc)
}

func (jc *JSONInvoiceD2DItems) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, jc)
}

// Export and Import invoice
type InvoiceExport struct {
	gorm.Model
	ID                int
	InvoiceNumber     string
	Type              string
	CustomerID        int `gorm:"foreignKey:CustomerID"`
	Customer          Customer
	ConsigneeID       int `gorm:"foreignKey:ConsigneeID"`
	Consignee         Shipper
	ShipperID         int `gorm:"foreignKey:ShipperID"`
	Shipper           Customer
	Service           string
	BLAWB             string
	AJU               string
	PortOfLoadingID   int `gorm:"foreignKey:PortOfLoadingID"`
	PortOfLoading     Port
	PortOfDischargeID int `gorm:"foreignKey:PortOfDischargeID"`
	PortOfDischarge   Port
	ShippingMarks     string
	InvoiceDate       time.Time
	Status            string
	InvoiceItems      JSONInvoiceItems `gorm:"type:json"`
	ShipmentID        int              `gorm:"foreignKey:ShipmentID"`
}

type InvoiceImport struct {
	gorm.Model
	ID                int
	InvoiceNumber     string
	Type              string
	CustomerID        int `gorm:"foreignKey:CustomerID"`
	Customer          Customer
	ConsigneeID       int `gorm:"foreignKey:ConsigneeID"`
	Consignee         Customer
	ShipperID         int `gorm:"foreignKey:ShipperID"`
	Shipper           Shipper
	Service           string
	BLAWB             string
	AJU               string
	PortOfLoadingID   int `gorm:"foreignKey:PortOfLoadingID"`
	PortOfLoading     Port
	PortOfDischargeID int `gorm:"foreignKey:PortOfDischargeID"`
	PortOfDischarge   Port
	ShippingMarks     string
	InvoiceDate       time.Time
	Status            string
	InvoiceItems      JSONInvoiceItems `gorm:"type:json"`
	ShipmentID        int              `gorm:"foreignKey:ShipmentID"`
}

// Door to Door invoice
type DoorToDoorInvoice struct {
	gorm.Model
	ID                int
	InvoiceNumber     string
	Type              string
	CustomerID        int `gorm:"foreignKey:CustomerID"`
	Customer          Customer
	ConsigneeID       int `gorm:"foreignKey:ConsigneeID"`
	Consignee         Customer
	ShipperID         int `gorm:"foreignKey:ShipperID"`
	Shipper           Shipper
	Service           string
	PortOfLoadingID   int `gorm:"foreignKey:PortOfLoadingID"`
	PortOfLoading     Port
	PortOfDischargeID int `gorm:"foreignKey:PortOfDischargeID"`
	PortOfDischarge   Port
	ShippingMarks     string
	InvoiceDate       time.Time
	Status            string
	Quantity          string
	Weight            float64
	Volume            float64
	InvoiceItems      JSONInvoiceD2DItems `gorm:"type:json"`
	ShipmentID        int                 `gorm:"foreignKey:ShipmentID"`
}
