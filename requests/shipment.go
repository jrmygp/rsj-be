package requests

type ShipmentRequest struct {
	ShipmentNumber     string `json:"shipmentNumber" binding:"required"`
	Quotations         []int  `json:"quotations"`         // List of Quotation IDs associated with this Shipment
	InvoiceExports     []int  `json:"invoiceExports"`     // List of InvoiceExport IDs
	InvoiceImports     []int  `json:"invoiceImports"`     // List of InvoiceImport IDs
	InvoiceDoorToDoors []int  `json:"invoiceDoorToDoors"` // List of DoorToDoorInvoice IDs
}
