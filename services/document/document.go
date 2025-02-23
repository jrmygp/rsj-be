package services

import (
	"server/models"
	"server/requests"
)

type Service interface {
	// Surat Tugas
	FindAllSuratTugas(searchQuery string, page int) ([]models.SuratTugas, int64, int, int, int)
	CreateSuratTugas(suratTugas requests.CreateSuratTugasRequest) (models.SuratTugas, error)
	FindSuratTugasByID(ID int) (models.SuratTugas, error)
	EditSuratTugas(ID int, suratTugas requests.EditSuratTugasRequest) (models.SuratTugas, error)
	DeleteSuratTugas(ID int) (models.SuratTugas, error)

	// Surat Jalan
	FindAllSuratJalan(searchQuery string, page int) ([]models.SuratJalan, int64, int, int, int)
	CreateSuratJalan(suratJalan requests.CreateSuratJalanRequest) (models.SuratJalan, error)
	FindSuratJalanByID(ID int) (models.SuratJalan, error)
	EditSuratJalan(ID int, suratJalan requests.EditSuratJalanRequest) (models.SuratJalan, error)
	DeleteSuratJalan(ID int) (models.SuratJalan, error)
}
