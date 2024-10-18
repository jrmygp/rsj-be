package models

import "time"

type User struct {
	ID         int
	Name       string
	Username   string
	Password   string
	UserRoleID uint      `gorm:"not null"`
	CreatedAt  time.Time `gorm:"foreignKey:UserRoleID"`
	UpdatedAt  time.Time
}
