package models

type Customer struct {
	ID      int
	Name    string
	Address string `gorm:"type:text"`
}

type Port struct {
	ID       int
	PortName string
	Note     string `gorm:"type:text"`
	Status   string
}

type CostCharges struct {
	ID     int
	Name   string
	Status string
}
