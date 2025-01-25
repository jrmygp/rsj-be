package services

import (
	"server/models"
	"server/requests"
)

type Service interface {
	FindAllSuratTugas(searchQuery string, page int) ([]models.SuratTugas, int64, int, int, int)
	CreateSuratTugas(suratTugas requests.CreateSuratTugasRequest) (models.SuratTugas, error)
	FindSuratTugasByID(ID int) (models.SuratTugas, error)
	EditSuratTugas(ID int, suratTugas requests.EditSuratTugasRequest) (models.SuratTugas, error)
	DeleteSuratTugas(ID int) (models.SuratTugas, error)
}
