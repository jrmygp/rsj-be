package requests

import "server/helper"

type InvoiceItemRequest struct {
	ItemName string   `json:"itemName"`
	Currency string   `json:"currency"`
	Price    float64  `json:"price"`
	Kurs     *float64 `json:"kurs"`
	Quantity int      `json:"quantity"`
}

type CreateInvoiceRequest struct {
	Category      string               `json:"category" binding:"required"`
	InvoiceNumber string               `json:"invoiceNumber"`
	Type          string               `json:"type"`
	CustomerID    int                  `json:"customerId"`
	ConsigneeID   int                  `json:"consigneeId"`
	ShipperID     int                  `json:"shipperId"`
	Service       string               `json:"service"`
	BLAWB         string               `json:"blawb"`
	AJU           string               `json:"aju"`
	POL           string               `json:"pol"`
	POD           string               `json:"pod"`
	ShippingMarks string               `json:"shippingMarks"`
	InvoiceDate   helper.CustomDate    `json:"invoiceDate"`
	Status        string               `json:"status"`
	InvoiceItems  []InvoiceItemRequest `json:"invoiceItems"`
}

type EditInvoiceRequest struct {
	Category      string               `json:"category"`
	InvoiceNumber string               `json:"invoiceNumber"`
	Type          string               `json:"type"`
	CustomerID    int                  `json:"customerId"`
	ConsigneeID   int                  `json:"consigneeId"`
	ShipperID     int                  `json:"shipperId"`
	Service       string               `json:"service"`
	BLAWB         string               `json:"blawb"`
	AJU           string               `json:"aju"`
	POL           string               `json:"pol"`
	POD           string               `json:"pod"`
	ShippingMarks string               `json:"shippingMarks"`
	InvoiceDate   helper.CustomDate    `json:"invoiceDate"`
	Status        string               `json:"status"`
	InvoiceItems  []InvoiceItemRequest `json:"invoiceItems"`
}
