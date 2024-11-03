package responses

import "time"

type ChargeResponse struct {
	ItemName string   `json:"itemName"`
	Price    float64  `json:"price"`
	Currency string   `json:"currency"`
	RatioIDR *float64 `json:"ratioIdr"`
	Quantity int      `json:"quantity"`
	Unit     string   `json:"unit"`
	Note     *string  `json:"note"`
}

type QuotationResponse struct {
	QuotationNumber   string           `json:"quotationNumber"`
	RateValidity      time.Time        `json:"rateValidity"`
	ShippingTerm      string           `json:"shippingTerm"`
	Service           string           `json:"service"`
	Status            string           `json:"status"`
	Commodity         string           `json:"commodity"`
	Weight            int              `json:"weight"`
	Volume            int              `json:"volume"`
	Note              string           `json:"note"`
	SalesID           int              `json:"salesId"`
	CustomerID        int              `json:"customerId"`
	PortOfLoadingID   int              `json:"portOfLoadingId"`
	PortOfDischargeID int              `json:"portOfDischargeId"`
	ListCharges       []ChargeResponse `json:"listCharges"`
}
