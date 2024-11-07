package responses

type ChargeResponse struct {
	ItemName string  `json:"itemName"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
	RatioIDR *int    `json:"ratioIdr"`
	Quantity int     `json:"quantity"`
	Unit     string  `json:"unit"`
	Note     *string `json:"note"`
}

type QuotationResponse struct {
	ID                  int              `json:"id"`
	QuotationNumber     string           `json:"quotationNumber"`
	RateValidity        string           `json:"rateValidity"`
	ShippingTerm        string           `json:"shippingTerm"`
	Service             string           `json:"service"`
	Status              string           `json:"status"`
	Commodity           string           `json:"commodity"`
	Weight              int              `json:"weight"`
	Volume              int              `json:"volume"`
	Note                string           `json:"note"`
	PaymentTerm         string           `json:"paymentTerm"`
	SalesID             int              `json:"salesId"`
	SalesName           string           `json:"salesName"`
	CustomerID          int              `json:"customerId"`
	CustomerName        string           `json:"customerName"`
	PortOfLoadingID     int              `json:"portOfLoadingId"`
	PortOfLoadingName   string           `json:"portOfLoadingName"`
	PortOfDischargeID   int              `json:"portOfDischargeId"`
	PortOfDischargeName string           `json:"portOfDischargeName"`
	ListCharges         []ChargeResponse `json:"listCharges"`
}
