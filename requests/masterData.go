package requests

type CreateCustomerRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}

type EditCustomerRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}
