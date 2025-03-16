package responses

type LinkedDocResponse struct {
	ID        int    `json:"id"`
	DocNumber string `json:"docNumber"`
}

type ShippingDetailResponse struct {
	ShippingMark    string  `json:"shippingMark"`
	ContainerNumber string  `json:"containerNumber"`
	SEAL            string  `json:"seal"`
	SaidOfContain   string  `json:"saidOfContainer"`
	NettWeight      float64 `json:"nettWeight"`
	GrossWeight     float64 `json:"grossWeight"`
}

type ShipmentResponse struct {
	ID                 int                      `json:"id"`
	ShipmentNumber     string                   `json:"shipmentNumber"`
	WarehouseID        int                      `json:"warehouseId"`
	WarehouseName      string                   `json:"warehouseName"`
	Quotations         []LinkedDocResponse      `json:"quotations"`
	InvoiceExports     []LinkedDocResponse      `json:"invoiceExports"`
	InvoiceImports     []LinkedDocResponse      `json:"invoiceImports"`
	InvoiceDoorToDoors []LinkedDocResponse      `json:"invoiceDoorToDoors"`
	ShippingDetails    []ShippingDetailResponse `json:"shippingDetails"`
}
