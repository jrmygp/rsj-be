package models

type Customer struct {
	ID          int
	Name        string
	Address     string `gorm:"type:text"`
	CompanyCode string
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

type Shipper struct {
	ID      int
	Name    string
	Address string `gorm:"type:text"`
}

type Warehouse struct {
	ID         int
	Category   string
	Name       string
	Code       string
	FlightName *string
	FlightCode *int
}
