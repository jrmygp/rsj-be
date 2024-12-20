package responses

type InvoiceItemResponse struct {
	ItemName string   `json:"itemName"`
	Currency string   `json:"currency"`
	Price    float64  `json:"price"`
	Kurs     *float64 `json:"kurs"`
	Quantity int      `json:"quantity"`
}

type InvoiceResponse struct {
	ID            int                   `json:"id"`
	InvoiceNumber string                `json:"invoiceNumber"`
	Type          string                `json:"type"`
	CustomerID    int                   `json:"customerId"`
	ConsigneeID   int                   `json:"consigneeId"`
	ShipperID     int                   `json:"shipperId"`
	Service       string                `json:"service"`
	BLAWB         string                `json:"blawb"`
	AJU           string                `json:"aju"`
	POL           string                `json:"pol"`
	POD           string                `json:"pod"`
	ShippingMarks string                `json:"shippingMarks"`
	InvoiceDate   string                `json:"invoiceDate"`
	Status        string                `json:"status"`
	InvoiceItems  []InvoiceItemResponse `json:"invoiceItems"`
}

type DoorToDoorResponse struct {
	ID            int                   `json:"id"`
	InvoiceNumber string                `json:"invoiceNumber"`
	Type          string                `json:"type"`
	CustomerID    int                   `json:"customerId"`
	ConsigneeID   int                   `json:"consigneeId"`
	ShipperID     int                   `json:"shipperId"`
	Service       string                `json:"service"`
	POL           string                `json:"pol"`
	POD           string                `json:"pod"`
	ShippingMarks string                `json:"shippingMarks"`
	InvoiceDate   string                `json:"invoiceDate"`
	Status        string                `json:"status"`
	Quantity      int                   `json:"quantity"`
	Weight        float64               `json:"weight"`
	Volume        float64               `json:"volume"`
	InvoiceItems  []InvoiceItemResponse `json:"invoiceItems"`
}
