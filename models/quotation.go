package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Charge struct {
	ItemID     int     `json:"itemId"`     // Foreign key to CostCharges ID
	ItemName   string  `json:"itemName"`   // Name of the item from CostCharges
	Price      float64 `json:"price"`      // Price of the item
	Currency   string  `json:"currency"`   // Currency of item price
	RatioToIDR *int    `json:"ratioToIdr"` // Nullable integer for ratio to IDR
	Quantity   int     `json:"quantity"`   // Quantity of the item
	Unit       string  `json:"unit"`       // Unit of measurement
	Note       *string `json:"note"`       // Additional notes
}

// JSONCharges is a custom type that wraps []Charge
type JSONCharges []Charge

// Value marshals JSONCharges to JSON for storing in the database
func (jc JSONCharges) Value() (driver.Value, error) {
	return json.Marshal(jc)
}

// Scan unmarshals JSON data from the database into JSONCharges
func (jc *JSONCharges) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, jc)
}

type Quotation struct {
	gorm.Model
	ID                int
	QuotationNumber   string
	RateValidity      time.Time
	ShippingTerm      string
	Service           string
	Status            string
	Commodity         string
	Weight            int
	Volume            int
	Note              string
	SalesID           int
	Sales             User     `gorm:"foreignKey:SalesID"`
	CustomerID        int      `gorm:"not null;foreignKey:CustomerID"`
	Customer          Customer `gorm:"foreignKey:CustomerID"`
	PortOfLoadingID   int
	PortOfLoading     Port `gorm:"foreignKey:PortOfLoadingID"`
	PortOfDischargeID int
	PortOfDischarge   Port        `gorm:"foreignKey:PortOfDischargeID"`
	ListCharges       JSONCharges `gorm:"type:json"` // Define as JSON type in MySQL
	// ShipmentID
}
