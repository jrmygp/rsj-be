package helper

import (
	"fmt"
	"log"
	"server/models"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
)

func GenerateSuratJalanPDF(suratJalan models.SuratJalan) {
	cfg := config.NewBuilder().
		WithOrientation(orientation.Vertical).
		WithPageSize(pagesize.A4).
		WithLeftMargin(15).
		WithTopMargin(0).
		WithRightMargin(15).
		WithBottomMargin(10).
		Build()
	m := maroto.New(cfg)

	err := m.RegisterHeader(addSuratTugasHeader())
	if err != nil {
		log.Println("Error generating PDF Header:", err.Error())
	}

	document, err := m.Generate()
	if err != nil {
		log.Println("Error generating PDF:", err.Error())
	}

	safeSuratJalanNumber := SanitizeFilename(suratJalan.DocumentNumber)

	filePath := fmt.Sprintf("pdf/surat-jalan/%s.pdf", safeSuratJalanNumber)

	err = document.Save(filePath)
	if err != nil {
		log.Println("Error generating PDF:", err.Error())
	}
	log.Println("PDF saved successfully")
}
