package models

import "time"

type UserRole struct {
	ID        int
	Role      string
	Users     []User `gorm:"foreignKey:UserRoleID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
