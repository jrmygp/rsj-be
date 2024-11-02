package responses

type CustomerResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type PortResponse struct {
	ID       int    `json:"id"`
	PortName string `json:"portName"`
	Note     string `json:"note"`
	Status   string `json:"status"`
}
