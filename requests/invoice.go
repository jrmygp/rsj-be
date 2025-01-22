package requests

import "server/helper"

type InvoiceFilterRequest struct {
	CustomerID int `json:"customerId"`
}

type InvoiceItemRequest struct {
	ItemName string   `json:"itemName"`
	Currency string   `json:"currency"`
	Price    float64  `json:"price"`
	Kurs     *float64 `json:"kurs"`
	Quantity int      `json:"quantity"`
	Unit     string   `json:"unit"`
}

type InvoiceD2DItemRequest struct {
	ItemName string   `json:"itemName"`
	Currency string   `json:"currency"`
	Price    float64  `json:"price"`
	Kurs     *float64 `json:"kurs"`
	Quantity int      `json:"quantity"`
}

type CreateInvoiceRequest struct {
	InvoiceNumber     string               `json:"invoiceNumber"`
	Type              string               `json:"type"`
	CustomerID        int                  `json:"customerId"`
	ConsigneeID       int                  `json:"consigneeId"`
	ShipperID         int                  `json:"shipperId"`
	Service           string               `json:"service"`
	BLAWB             string               `json:"blawb"`
	AJU               string               `json:"aju"`
	PortOfLoadingID   int                  `json:"portOfLoadingId"`
	PortOfDischargeID int                  `json:"portOfDischargeId"`
	ShippingMarks     string               `json:"shippingMarks"`
	InvoiceDate       helper.CustomDate    `json:"invoiceDate"`
	Status            string               `json:"status"`
	InvoiceItems      []InvoiceItemRequest `json:"invoiceItems"`
}

type EditInvoiceRequest struct {
	InvoiceNumber     string               `json:"invoiceNumber"`
	Type              string               `json:"type"`
	CustomerID        int                  `json:"customerId"`
	ConsigneeID       int                  `json:"consigneeId"`
	ShipperID         int                  `json:"shipperId"`
	Service           string               `json:"service"`
	BLAWB             string               `json:"blawb"`
	AJU               string               `json:"aju"`
	PortOfLoadingID   int                  `json:"portOfLoadingId"`
	PortOfDischargeID int                  `json:"portOfDischargeId"`
	ShippingMarks     string               `json:"shippingMarks"`
	InvoiceDate       helper.CustomDate    `json:"invoiceDate"`
	Status            string               `json:"status"`
	InvoiceItems      []InvoiceItemRequest `json:"invoiceItems"`
}

type CreateDoorToDoorRequest struct {
	InvoiceNumber     string                  `json:"invoiceNumber"`
	Type              string                  `json:"type"`
	CustomerID        int                     `json:"customerId"`
	ConsigneeID       int                     `json:"consigneeId"`
	ShipperID         int                     `json:"shipperId"`
	Service           string                  `json:"service"`
	PortOfLoadingID   int                     `json:"portOfLoadingId"`
	PortOfDischargeID int                     `json:"portOfDischargeId"`
	ShippingMarks     string                  `json:"shippingMarks"`
	InvoiceDate       helper.CustomDate       `json:"invoiceDate"`
	Status            string                  `json:"status"`
	Quantity          string                  `json:"quantity"`
	Weight            float64                 `json:"weight"`
	Volume            float64                 `json:"volume"`
	InvoiceItems      []InvoiceD2DItemRequest `json:"invoiceItems"`
}

type EditDoorToDoorRequest struct {
	InvoiceNumber     string                  `json:"invoiceNumber"`
	Type              string                  `json:"type"`
	CustomerID        int                     `json:"customerId"`
	ConsigneeID       int                     `json:"consigneeId"`
	ShipperID         int                     `json:"shipperId"`
	Service           string                  `json:"service"`
	PortOfLoadingID   int                     `json:"portOfLoadingId"`
	PortOfDischargeID int                     `json:"portOfDischargeId"`
	ShippingMarks     string                  `json:"shippingMarks"`
	InvoiceDate       helper.CustomDate       `json:"invoiceDate"`
	Status            string                  `json:"status"`
	Quantity          string                  `json:"quantity"`
	Weight            float64                 `json:"weight"`
	Volume            float64                 `json:"volume"`
	InvoiceItems      []InvoiceD2DItemRequest `json:"invoiceItems"`
}
