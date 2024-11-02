package models

import (
	"time"

	"gorm.io/gorm"
)

type Charge struct {
	ItemID     int    `json:"itemId"`     // Foreign key to CostCharges ID
	Item       string `json:"item"`       // Name of the item from CostCharges
	Price      int    `json:"price"`      // Price of the item
	RatioToIDR *int   `json:"ratioToIdr"` // Nullable integer for ratio to IDR
	Quantity   int    `json:"quantity"`   // Quantity of the item
	Unit       string `json:"unit"`       // Unit of measurement
	Note       string `json:"note"`       // Additional notes
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
	CustomerID        uint     `gorm:"not null;foreignKey:CustomerID"`
	Customer          Customer `gorm:"foreignKey:CustomerID"`
	PortOfLoadingID   int
	PortOfLoading     Port `gorm:"foreignKey:PortOfLoadingID"`
	PortOfDischargeID int
	PortOfDischarge   Port     `gorm:"foreignKey:PortOfDischargeID"`
	ListCharges       []Charge `gorm:"type:json"`
}
