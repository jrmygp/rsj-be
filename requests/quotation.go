package requests

import (
	"server/helper"
)

type ChargeRequest struct {
	ItemID   int      `json:"itemId" binding:"required"`
	ItemName string   `json:"itemName" binding:"required"`
	Price    float64  `json:"price" binding:"required"`
	Currency string   `json:"currency" binding:"required"`
	RatioIDR *float64 `json:"ratioIdr"`
	Quantity int      `json:"quantity" binding:"required"`
	Unit     *string  `json:"unit" binding:"required"`
	Note     *string  `json:"note"`
}

type CreateQuotationRequest struct {
	QuotationNumber   string            `json:"quotationNumber" binding:"required"`
	RateValidity      helper.CustomDate `json:"rateValidity" binding:"required"`
	ShippingTerm      string            `json:"shippingTerm" binding:"required"`
	Service           string            `json:"service" binding:"required"`
	Status            string            `json:"status" binding:"required"`
	Commodity         string            `json:"commodity" binding:"required"`
	Weight            float64           `json:"weight" binding:"required"`
	Volume            float64           `json:"volume" binding:"required"`
	Note              string            `json:"note" binding:"required"`
	PaymentTerm       string            `json:"paymentTerm" binding:"required"`
	SalesID           int               `json:"salesId" binding:"required"`
	CustomerID        int               `json:"customerId" binding:"required"`
	PortOfLoadingID   int               `json:"portOfLoadingId" binding:"required"`
	PortOfDischargeID int               `json:"portOfDischargeId" binding:"required"`
	ListCharges       []ChargeRequest   `json:"listCharges" binding:"required,dive"`
}

type EditQuotationRequest struct {
	QuotationNumber   string            `json:"quotationNumber"`
	RateValidity      helper.CustomDate `json:"rateValidity"`
	ShippingTerm      string            `json:"shippingTerm"`
	Service           string            `json:"service"`
	Status            string            `json:"status"`
	Commodity         string            `json:"commodity"`
	Weight            float64           `json:"weight"`
	Volume            float64           `json:"volume"`
	Note              string            `json:"note"`
	PaymentTerm       string            `json:"paymentTerm"`
	SalesID           int               `json:"salesId"`
	CustomerID        int               `json:"customerId"`
	PortOfLoadingID   int               `json:"portOfLoadingId"`
	PortOfDischargeID int               `json:"portOfDischargeId"`
	ListCharges       []ChargeRequest   `json:"listCharges"`
}
