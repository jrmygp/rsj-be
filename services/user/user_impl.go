package services

import (
	"errors"
	"os"
	"server/models"
	repositories "server/repositories/user"
	"server/requests"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

func (s *service) Login(userRequest requests.LoginUserRequest) (models.User, string, error) {
	user, err := s.repository.Login(userRequest.Username)
	if err != nil {
		return models.User{}, "", errors.New("invalid username or password")
	}

	if user.ID == 0 { // Check if user is found
		return models.User{}, "", errors.New("user not found")
	}

	// Compare inputted password and hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password))
	if err != nil {
		return models.User{}, "", errors.New("invalid username or password")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		return models.User{}, "", err
	}

	// Send it back

	return user, tokenString, err
}
