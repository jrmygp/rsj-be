package responses

type CustomerResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	CompanyCode string `json:"companyCode"`
}

type PortResponse struct {
	ID       int    `json:"id"`
	PortName string `json:"portName"`
	Note     string `json:"note"`
	Status   string `json:"status"`
}

type CostChargesResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type ShipperResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type WarehouseResponse struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Code       string  `json:"code"`
	FlightName *string `json:"flightName"`
	FlightCode *int    `json:"flightCode"`
}
