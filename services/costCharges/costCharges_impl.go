package services

import (
	"errors"
	"server/models"
	repositories "server/repositories/costCharges"
	"server/requests"

	"gorm.io/gorm"
)

type service struct {
	repository repositories.Repository
}

func NewService(repository repositories.Repository) *service {
	return &service{repository}
}

func (s *service) Create(costChargesRequest requests.CreateCostChargesRequest) (models.CostCharges, error) {
	costCharge := models.CostCharges{
		Name:   costChargesRequest.Name,
		Status: costChargesRequest.Status,
	}

	newCostCharge, err := s.repository.Create(costCharge)
	return newCostCharge, err
}

func (s *service) FindByID(ID int) (models.CostCharges, error) {
	costCharge, err := s.repository.FindByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.CostCharges{}, nil
	}

	return costCharge, err
}

func (s *service) Edit(ID int, costChargesRequest requests.EditCostChargesRequest) (models.CostCharges, error) {
	costCharge, _ := s.repository.FindByID(ID)

	if costChargesRequest.Name != "" {
		costCharge.Name = costChargesRequest.Name
	}
	if costChargesRequest.Status != "" {
		costCharge.Status = costChargesRequest.Status
	}

	updatedCostCharges, err := s.repository.Edit(costCharge)
	return updatedCostCharges, err
}

func (s *service) Delete(ID int) (models.CostCharges, error) {
	costCharge, err := s.repository.Delete(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.CostCharges{}, nil
	}

	return costCharge, err
}

func (s *service) FindAll(searchQuery string, page int) ([]models.CostCharges, int64, int, int, int) {
	if page < 1 {
		return []models.CostCharges{}, 0, 0, 0, 0 // Handle invalid page
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	costCharge, totalCount := s.repository.FindAll(searchQuery, offset, pageSize)

	firstRow := offset + 1
	lastRow := offset + len(costCharge)
	if len(costCharge) == 0 {
		firstRow = 0
		lastRow = 0
	}
	totalPages := (int(totalCount) + pageSize - 1) / pageSize

	return costCharge, totalCount, firstRow, lastRow, totalPages
}
