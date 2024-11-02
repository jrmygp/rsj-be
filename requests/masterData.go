package requests

type CreateCustomerRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}

type EditCustomerRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type CreatePortRequest struct {
	PortName string `json:"portName" binding:"required"`
	Note     string `json:"note" binding:"required"`
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
