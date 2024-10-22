package services

import (
	"server/models"
	"server/requests"
)

type Service interface {
	FindAll() ([]models.User, error)
	FindByID(ID int) (models.User, error)
	Create(user requests.CreateUserRequest) (models.User, error)
	Login(user requests.LoginUserRequest) (models.User, string, error)
}
