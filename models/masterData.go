package models

type Customer struct {
	ID      int
	Name    string
	Address string
}

type Port struct {
	ID       int
	PortName string
	Note     string
	Status   string
}

type CostCharges struct {
	ID     int
	Name   string
	Status string
}
