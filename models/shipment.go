package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type ShippingDetail struct {
	ShippingMark    string  `json:"shippingMark"`
	ContainerNumber string  `json:"containerNumber"`
	SEAL            string  `json:"seal"`
	SaidOfContain   string  `json:"saidOfContainer"`
	NettWeight      float64 `json:"nettWeight"`
	GrossWeight     float64 `json:"grossWeight"`
}

type JSONShippingDetail []ShippingDetail

func (jc JSONShippingDetail) Value() (driver.Value, error) {
	return json.Marshal(jc)
}

func (jc *JSONShippingDetail) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, jc)
}

type Shipment struct {
	gorm.Model
	ID                 int
	ShipmentNumber     string
	Quotations         []Quotation         `gorm:"foreignKey:ShipmentID"`
	InvoiceExports     []InvoiceExport     `gorm:"foreignKey:ShipmentID"`
	InvoiceImports     []InvoiceImport     `gorm:"foreignKey:ShipmentID"`
	InvoiceDoorToDoors []DoorToDoorInvoice `gorm:"foreignKey:ShipmentID"`
}
