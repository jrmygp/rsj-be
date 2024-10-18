package services

import (
	"errors"
	"server/models"
	repositories "server/repositories/user"
	"server/requests"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type service struct {
	repository repositories.Repository
}

func NewService(repository repositories.Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]models.User, error) {
	users, err := s.repository.FindAll()
	return users, err
}

func (s *service) FindByID(ID int) (models.User, error) {
	user, err := s.repository.FindByID(ID)

	// Check if the error is record not found
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Return an empty user and nil error to indicate no user found, but no error
		return models.User{}, nil
	}

	// Return the user and any other errors (e.g., DB connection issues)
	return user, err
}

func (s *service) Create(userRequest requests.CreateUserRequest) (models.User, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Name:       userRequest.Name,
		Username:   userRequest.Username,
		Password:   string(hashedPassword),
		UserRoleID: userRequest.UserRoleID,
	}

	newUser, err := s.repository.Create(user)
	return newUser, err
}
