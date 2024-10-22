package models

type UserRole struct {
	ID    int
	Role  string
	Users []User `gorm:"foreignKey:UserRoleID"`
}
