package repositories

import "server/models"

type Repository interface {
	FindAll() ([]models.User, error)
	FindByID(ID int) (models.User, error)
	Create(user models.User) (models.User, error)
}
