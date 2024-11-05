package services

import (
	"errors"
	"server/models"
	repositories "server/repositories/port"
	"server/requests"

	"gorm.io/gorm"
)

type service struct {
	repository repositories.Repository
}

func NewService(repository repositories.Repository) *service {
	return &service{repository}
}

func (s *service) FindAllNoPagination() ([]models.Port, error) {
	ports, err := s.repository.FindAllNoPagination()
	return ports, err
}

func (s *service) Create(portRequest requests.CreatePortRequest) (models.Port, error) {
	port := models.Port{
		PortName: portRequest.PortName,
		Note:     portRequest.Note,
		Status:   portRequest.Status,
	}

	newPort, err := s.repository.Create(port)
	return newPort, err
}

func (s *service) FindByID(ID int) (models.Port, error) {
	port, err := s.repository.FindByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Port{}, nil
	}

	return port, err
}

func (s *service) Edit(ID int, portRequest requests.EditPortRequest) (models.Port, error) {
	port, _ := s.repository.FindByID(ID)

	if portRequest.PortName != "" {
		port.PortName = portRequest.PortName
	}
	if portRequest.Note != "" {
		port.Note = portRequest.Note
	}
	if portRequest.Status != "" {
		port.Status = portRequest.Status
	}

	updatedPort, err := s.repository.Edit(port)
	return updatedPort, err
}

func (s *service) Delete(ID int) (models.Port, error) {
	port, err := s.repository.Delete(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Port{}, nil
	}

	return port, err
}

func (s *service) FindAll(searchQuery string, page int) ([]models.Port, int64, int, int, int) {
	if page < 1 {
		return []models.Port{}, 0, 0, 0, 0 // Handle invalid page
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	port, totalCount := s.repository.FindAll(searchQuery, offset, pageSize)

	firstRow := offset + 1
	lastRow := offset + len(port)
	if len(port) == 0 {
		firstRow = 0
		lastRow = 0
	}
	totalPages := (int(totalCount) + pageSize - 1) / pageSize

	return port, totalCount, firstRow, lastRow, totalPages
}
