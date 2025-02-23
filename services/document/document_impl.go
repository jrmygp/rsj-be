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

// Surat Tugas
func (s *documentService) FindAllSuratTugas(searchQuery string, page int) ([]models.SuratTugas, int64, int, int, int) {
	if page < 1 {
		return []models.SuratTugas{}, 0, 0, 0, 0
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	document, totalCount := s.repository.FindAllSuratTugas(searchQuery, offset, pageSize)

	firstRow := offset + 1
	lastRow := offset + len(document)
	if len(document) == 0 {
		firstRow = 0
		lastRow = 0
	}
	totalPages := (int(totalCount) + pageSize - 1) / pageSize

	return document, totalCount, firstRow, lastRow, totalPages
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

// Surat Jalan
func (s *documentService) FindAllSuratJalan(searchQuery string, page int) ([]models.SuratJalan, int64, int, int, int) {
	if page < 1 {
		return []models.SuratJalan{}, 0, 0, 0, 0
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	document, totalCount := s.repository.FindAllSuratJalan(searchQuery, offset, pageSize)

	firstRow := offset + 1
	lastRow := offset + len(document)
	if len(document) == 0 {
		firstRow = 0
		lastRow = 0
	}
	totalPages := (int(totalCount) + pageSize - 1) / pageSize

	return document, totalCount, firstRow, lastRow, totalPages
}

func (s *documentService) CreateSuratJalan(suratJalan requests.CreateSuratJalanRequest) (models.SuratJalan, error) {
	items := make([]models.Item, len(suratJalan.Items))
	for i, item := range suratJalan.Items {
		items[i] = models.Item{
			ItemName: item.ItemName,
			Type:     item.Type,
			Quantity: item.Quantity,
			Colly:    item.Colly,
			Volume:   item.Volume,
			Unit:     item.Unit,
			Note:     item.Note,
		}
	}

	document := models.SuratJalan{
		DocumentNumber: suratJalan.DocumentNumber,
		Recipient:      suratJalan.Recipient,
		Address:        suratJalan.Address,
		Date:           suratJalan.Date.Time,
		Items:          items,
	}

	newDocument, err := s.repository.CreateSuratJalan(document)
	return newDocument, err
}

func (s *documentService) FindSuratJalanByID(ID int) (models.SuratJalan, error) {
	document, err := s.repository.FindSuratJalanByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.SuratJalan{}, nil
	}

	return document, err
}

func (s *documentService) EditSuratJalan(ID int, suratJalan requests.EditSuratJalanRequest) (models.SuratJalan, error) {
	document, err := s.repository.FindSuratJalanByID(ID)
	if err != nil {
		return models.SuratJalan{}, err
	}
	if suratJalan.DocumentNumber != "" {
		document.DocumentNumber = suratJalan.DocumentNumber
	}
	if suratJalan.Address != "" {
		document.Address = suratJalan.Address
	}
	if suratJalan.Recipient != "" {
		document.Recipient = suratJalan.Recipient
	}
	if !suratJalan.Date.IsZero() {
		document.Date = suratJalan.Date.Time
	}
	if len(suratJalan.Items) > 0 {
		var items models.JSONItems //
		for _, item := range suratJalan.Items {
			jsonCharge := models.Item{
				ItemName: item.ItemName,
				Type:     item.Type,
				Colly:    item.Colly,
				Quantity: item.Quantity,
				Volume:   item.Volume,
				Unit:     item.Unit,
				Note:     item.Note,
			}
			items = append(items, jsonCharge)
		}
		document.Items = items
	}

	updatedDocument, err := s.repository.EditSuratJalan(document)
	return updatedDocument, err
}

func (s *documentService) DeleteSuratJalan(ID int) (models.SuratJalan, error) {
	document, err := s.repository.DeleteSuratJalan(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.SuratJalan{}, nil
	}

	return document, err
}
