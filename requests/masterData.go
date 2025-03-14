package requests

type CreateCustomerRequest struct {
	Name        string `json:"name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	CompanyCode string `json:"companyCode" binding:"required"`
}

type EditCustomerRequest struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	CompanyCode string `json:"companyCode"`
}

type CreatePortRequest struct {
	PortName string `json:"portName" binding:"required"`
	Note     string `json:"note"`
	Status   string `json:"status" binding:"required"`
}

type EditPortRequest struct {
	PortName string `json:"portName"`
	Note     string `json:"note"`
	Status   string `json:"status"`
}

type CreateCostChargesRequest struct {
	Name   string `json:"name" binding:"required"`
	Status string `json:"status" binding:"required"`
}

type EditCostChargesRequest struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type CreateShipperRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}

type EditShipperRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type WarehouseRequest struct {
	Category   string  `json:"category" binding:"required"`
	Name       string  `json:"name" binding:"required"`
	Code       string  `json:"code" binding:"required"`
	FlightName *string `json:"flightName"`
	FlightCode *int    `json:"flightCode"`
}
