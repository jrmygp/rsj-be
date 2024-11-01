package models

import (
	"time"

	"gorm.io/gorm"
)

// type Quotation struct {
// 	gorm.Model
// 	ID              int
// 	QuotationNumber string
// 	// Customer diambil dari table customer
// 	RateValidity time.Time
// 	ShippingTerm string
// 	// PortOfLoading diambil dari table port
// 	// PortOfDischarge diambil dari table port
// 	Service string
// 	// Sales User
// 	Status    string
// 	Commodity string
// 	Weight    int
// 	Volume    int
// 	Note      string
//  Items     Array of objects
// [{item : diambil dari table cost charges,
// price: int, mata uang: string,
// ratio to idr : int (cuma mandatory kalau mata uang yang dipakai non idr),
// quantity: int, unit: string, note: string}]
// }

type Quotation struct {
	gorm.Model
	ID              int
	QuotationNumber string
	RateValidity    time.Time
	ShippingTerm    string
	Service         string
	Status          string
	Commodity       string
	Weight          int
	Volume          int
	Note            string
	SalesID         int  // Foreign key field to link to the User table
	Sales           User `gorm:"foreignKey:SalesID"` // Reference to the User struct
}
