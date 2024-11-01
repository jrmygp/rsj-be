package models

import (
	"time"
)

type User struct {
	ID         int
	Name       string
	Username   string `gorm:"unique"`
	Password   string
	UserRoleID uint `gorm:"not null;foreignKey:UserRoleID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	UserRole   UserRole    `gorm:"foreignKey:UserRoleID"`
	Quotations []Quotation `gorm:"foreignKey:SalesID"` // Establish the one-to-many relationship
}
