package repositories

import (
	"server/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewDocumentRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllSuratTugas(searchQuery string, offset int, pageSize int) (suratTugas []models.SuratTugas, totalCount int64) {
	result := r.db.Model(&models.SuratTugas{})

	if searchQuery != "" {
		result = result.Where("document_number LIKE ?", "%"+searchQuery+"%")
	}
	result.Count(&totalCount)

	result.Order("created_at DESC").Find(&suratTugas)

	return suratTugas, totalCount
}

func (r *repository) CreateSuratTugas(suratTugas models.SuratTugas) (models.SuratTugas, error) {
	err := r.db.Create(&suratTugas).Error
	if err == nil {
		err = r.db.First(&suratTugas, suratTugas.ID).Error
	}

	return suratTugas, err
}

func (r *repository) FindSuratTugasByID(ID int) (models.SuratTugas, error) {
	var document models.SuratTugas

	err := r.db.First(&document, ID).Error

	return document, err
}

func (r *repository) EditSuratTugas(suratTugas models.SuratTugas) (models.SuratTugas, error) {
	err := r.db.Save(&suratTugas).Error
	if err == nil {
		err = r.db.First(&suratTugas, suratTugas.ID).Error
	}

	return suratTugas, err
}

func (r *repository) DeleteSuratTugas(ID int) (models.SuratTugas, error) {
	var document models.SuratTugas
	if err := r.db.First(&document, ID).Error; err != nil {
		return document, err
	}

	err := r.db.Delete(&document).Error
	return document, err
}
