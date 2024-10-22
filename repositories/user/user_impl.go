package repositories

import (
	"errors"
	"server/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]models.User, error) {
	var users []models.User

	err := r.db.Find(&users).Error

	return users, err
}

func (r *repository) FindByID(ID int) (models.User, error) {
	var user models.User

	err := r.db.Preload("UserRole").First(&user, ID).Error
	return user, err
}

func (r *repository) Create(user models.User) (models.User, error) {
	var existingUser models.User

	// Check if the username already exists in the database
	err := r.db.Where("username = ?", user.Username).First(&existingUser).Error
	if err == nil {
		// Username already exists, return a custom error
		return models.User{}, errors.New("username already taken")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// If there's another unexpected error (not "record not found"), return the error
		return models.User{}, err
	}

	// If no errors and the username doesn't exist, create the new user
	err = r.db.Create(&user).Error
	return user, err
}

func (r *repository) Login(username string) (models.User, error) {
	var user models.User

	err := r.db.Preload("UserRole").First(&user, "username = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.User{}, nil // Return an empty user and nil error to indicate "not found"
	}
	return user, err
}
