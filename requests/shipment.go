package requests

type ShippingDetailRequest struct {
	ShippingMark    string  `json:"shippingMark"`
	ContainerNumber string  `json:"containerNumber"`
	SEAL            string  `json:"seal"`
	SaidOfContain   string  `json:"saidOfContainer"`
	NettWeight      float64 `json:"nettWeight"`
	GrossWeight     float64 `json:"grossWeight"`
}

type ShipmentRequest struct {
	ShipmentNumber     string                  `json:"shipmentNumber" binding:"required"`
	WarehouseID        int                     `json:"warehouseId" binding:"required"`
	Quotations         []int                   `json:"quotations"`         // List of Quotation IDs associated with this Shipment
	InvoiceExports     []int                   `json:"invoiceExports"`     // List of InvoiceExport IDs
	InvoiceImports     []int                   `json:"invoiceImports"`     // List of InvoiceImport IDs
	InvoiceDoorToDoors []int                   `json:"invoiceDoorToDoors"` // List of DoorToDoorInvoice IDs
	ShippingDetails    []ShippingDetailRequest `json:"shippingDetails"`
}
