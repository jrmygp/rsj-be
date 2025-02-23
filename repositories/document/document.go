package repositories

import "server/models"

type Repository interface {
	// Surat Tugas
	FindAllSuratTugas(searchQuery string, offset int, pageSize int) (suratTugas []models.SuratTugas, totalCount int64)
	CreateSuratTugas(suratTugas models.SuratTugas) (models.SuratTugas, error)
	FindSuratTugasByID(ID int) (models.SuratTugas, error)
	EditSuratTugas(suratTugas models.SuratTugas) (models.SuratTugas, error)
	DeleteSuratTugas(ID int) (models.SuratTugas, error)

	// Surat Jalan
	FindAllSuratJalan(searchQuery string, offset int, pageSize int) (suratJalan []models.SuratJalan, totalCount int64)
	CreateSuratJalan(suratJalan models.SuratJalan) (models.SuratJalan, error)
	FindSuratJalanByID(ID int) (models.SuratJalan, error)
	EditSuratJalan(suratJalan models.SuratJalan) (models.SuratJalan, error)
	DeleteSuratJalan(ID int) (models.SuratJalan, error)
}
