package repositories

import "server/models"

type Repository interface {
	FindAllSuratTugas(searchQuery string, offset int, pageSize int) (suratTugas []models.SuratTugas, totalCount int64)
	CreateSuratTugas(suratTugas models.SuratTugas) (models.SuratTugas, error)
	FindSuratTugasByID(ID int) (models.SuratTugas, error)
	EditSuratTugas(suratTugas models.SuratTugas) (models.SuratTugas, error)
	DeleteSuratTugas(ID int) (models.SuratTugas, error)
}
