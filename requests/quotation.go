package requests

import "time"

type ChargeRequest struct {
	ItemName string  `json:"itemName" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
	RatioIDR *int    `json:"ratioIdr"`
	Quantity int     `json:"quantity" binding:"required"`
	Unit     string  `json:"unit" binding:"required"`
	Note     *string `json:"note"`
}

type CreateQuotationRequest struct {
	QuotationNumber   string          `json:"quotationNumber" binding:"required"`
	RateValidity      time.Time       `json:"rateValidity" binding:"required"`
	ShippingTerm      string          `json:"shippingTerm" binding:"required"`
	Service           string          `json:"service" binding:"required"`
	Status            string          `json:"status" binding:"required"`
	Commodity         string          `json:"commodity" binding:"required"`
	Weight            int             `json:"weight" binding:"required"`
	Volume            int             `json:"volume" binding:"required"`
	Note              string          `json:"note" binding:"required"`
	SalesID           int             `json:"salesId" binding:"required"`
	CustomerID        int             `json:"customerId" binding:"required"`
	PortOfLoadingID   int             `json:"portOfLoadingId" binding:"required"`
	PortOfDischargeID int             `json:"portOfDischargeId" binding:"required"`
	ListCharges       []ChargeRequest `json:"listCharges" binding:"required,dive"`
}
