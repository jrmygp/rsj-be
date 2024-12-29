package responses

type InvoiceItemResponse struct {
	ItemName string   `json:"itemName"`
	Currency string   `json:"currency"`
	Price    float64  `json:"price"`
	Kurs     *float64 `json:"kurs"`
	Quantity int      `json:"quantity"`
}

type InvoiceResponse struct {
	ID                  int                   `json:"id"`
	InvoiceNumber       string                `json:"invoiceNumber"`
	Type                string                `json:"type"`
	CustomerID          int                   `json:"customerId"`
	CustomerName        string                `json:"customerName"`
	ConsigneeID         int                   `json:"consigneeId"`
	CosgineeName        string                `json:"consigneeName"`
	ShipperID           int                   `json:"shipperId"`
	ShipperName         string                `json:"shipperName"`
	Service             string                `json:"service"`
	BLAWB               string                `json:"blawb"`
	AJU                 string                `json:"aju"`
	PortOfLoadingID     int                   `json:"portOfLoadingId"`
	PortOfLoadingName   string                `json:"portOfLoadingName"`
	PortOfDischargeID   int                   `json:"portOfDischargeId"`
	PortOfDischargeName string                `json:"portOfDischargeName"`
	ShippingMarks       string                `json:"shippingMarks"`
	InvoiceDate         string                `json:"invoiceDate"`
	Status              string                `json:"status"`
	InvoiceItems        []InvoiceItemResponse `json:"invoiceItems"`
}

type DoorToDoorResponse struct {
	ID                  int                   `json:"id"`
	InvoiceNumber       string                `json:"invoiceNumber"`
	Type                string                `json:"type"`
	CustomerID          int                   `json:"customerId"`
	CustomerName        string                `json:"customerName"`
	ConsigneeID         int                   `json:"consigneeId"`
	CosgineeName        string                `json:"consigneeName"`
	ShipperID           int                   `json:"shipperId"`
	ShipperName         string                `json:"shipperName"`
	Service             string                `json:"service"`
	PortOfLoadingID     int                   `json:"portOfLoadingId"`
	PortOfLoadingName   string                `json:"portOfLoadingName"`
	PortOfDischargeID   int                   `json:"portOfDischargeId"`
	PortOfDischargeName string                `json:"portOfDischargeName"`
	ShippingMarks       string                `json:"shippingMarks"`
	InvoiceDate         string                `json:"invoiceDate"`
	Status              string                `json:"status"`
	Quantity            string                `json:"quantity"`
	Weight              float64               `json:"weight"`
	Volume              float64               `json:"volume"`
	InvoiceItems        []InvoiceItemResponse `json:"invoiceItems"`
}
