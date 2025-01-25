package services

import (
	"errors"
	"server/models"
	repositories "server/repositories/document"
	"server/requests"

	"gorm.io/gorm"
)

type documentService struct {
	repository repositories.Repository
}

func NewDocumentRepository(repository repositories.Repository) *documentService {
	return &documentService{repository}
}

func (s *documentService) FindAllSuratTugas(searchQuery string, page int) ([]models.SuratTugas, int64, int, int, int) {
	if page < 1 {
		return []models.SuratTugas{}, 0, 0, 0, 0
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	invoice, totalCount := s.repository.FindAllSuratTugas(searchQuery, offset, pageSize)

	firstRow := offset + 1
	lastRow := offset + len(invoice)
	if len(invoice) == 0 {
		firstRow = 0
		lastRow = 0
	}
	totalPages := (int(totalCount) + pageSize - 1) / pageSize

	return invoice, totalCount, firstRow, lastRow, totalPages
}

func (s *documentService) CreateSuratTugas(suratTugas requests.CreateSuratTugasRequest) (models.SuratTugas, error) {
	document := models.SuratTugas{
		DocumentNumber: suratTugas.DocumentNumber,
		Assignor:       suratTugas.Assignor,
		Assignee:       suratTugas.Assignee,
		Liners:         suratTugas.Liners,
		Type:           suratTugas.Type,
		BLAWB:          suratTugas.BLAWB,
		Date:           suratTugas.Date.Time,
	}

	newDocument, err := s.repository.CreateSuratTugas(document)
	return newDocument, err
}

func (s *documentService) FindSuratTugasByID(ID int) (models.SuratTugas, error) {
	document, err := s.repository.FindSuratTugasByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.SuratTugas{}, nil
	}

	return document, err
}

func (s *documentService) EditSuratTugas(ID int, suratTugas requests.EditSuratTugasRequest) (models.SuratTugas, error) {
	document, err := s.repository.FindSuratTugasByID(ID)
	if err != nil {
		return models.SuratTugas{}, err
	}
	if suratTugas.DocumentNumber != "" {
		document.DocumentNumber = suratTugas.DocumentNumber
	}
	if suratTugas.Assignor != "" {
		document.Assignor = suratTugas.Assignor
	}
	if suratTugas.Assignee != "" {
		document.Assignee = suratTugas.Assignee
	}
	if suratTugas.Liners != "" {
		document.Liners = suratTugas.Liners
	}
	if suratTugas.Type != "" {
		document.Type = suratTugas.Type
	}
	if suratTugas.BLAWB != "" {
		document.BLAWB = suratTugas.BLAWB
	}
	if !suratTugas.Date.IsZero() {
		document.Date = suratTugas.Date.Time
	}

	updatedDocument, err := s.repository.EditSuratTugas(document)
	return updatedDocument, err
}

func (s *documentService) DeleteSuratTugas(ID int) (models.SuratTugas, error) {
	document, err := s.repository.DeleteSuratTugas(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.SuratTugas{}, nil
	}

	return document, err
}
