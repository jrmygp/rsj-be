package models

import (
	"time"

	"gorm.io/gorm"
)

type SuratTugas struct {
	gorm.Model
	ID             int
	DocumentNumber string
	Assignor       string
	Assignee       string
	Liners         string
	Type           string
	BLAWB          string
	Date           time.Time
}

type SuratJalan struct {
	gorm.Model
}
