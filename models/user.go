package models

import (
	"time"
)

type User struct {
	ID         int
	Name       string
	Username   string `gorm:"unique"`
	Password   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	UserRoleID uint     `gorm:"not null;foreignKey:UserRoleID"`
	UserRole   UserRole `gorm:"foreignKey:UserRoleID"`
}
