package responses

type LinkedDocResponse struct {
	ID        int    `json:"id"`
	DocNumber string `json:"docNumber"`
}

type ShipmentResponse struct {
	ID                 int                 `json:"id"`
	ShipmentNumber     string              `json:"shipmentNumber"`
	Quotations         []LinkedDocResponse `json:"quotations"`
	InvoiceExports     []LinkedDocResponse `json:"invoiceExports"`
	InvoiceImports     []LinkedDocResponse `json:"invoiceImports"`
	InvoiceDoorToDoors []LinkedDocResponse `json:"invoiceDoorToDoors"`
}
