package helper

import (
	"fmt"
	"log"
	"server/models"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/signature"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func GenerateSuratTugasPDF(suratTugas models.SuratTugas) {
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

	addSuratTugasNumber(m, suratTugas)
	addSuratTugasBody(m, suratTugas)
	addSuratTugasFooter(m, suratTugas)

	document, err := m.Generate()
	if err != nil {
		log.Println("Error generating PDF:", err.Error())
	}

	safeSuratTugasNumber := SanitizeFilename(suratTugas.DocumentNumber)

	filePath := fmt.Sprintf("pdf/surat-tugas/%s.pdf", safeSuratTugasNumber)

	err = document.Save(filePath)
	if err != nil {
		log.Println("Error generating PDF:", err.Error())
	}
	log.Println("PDF saved successfully")
}

func addSuratTugasHeader() core.Row {
	return row.New(40).Add(
		image.NewFromFileCol(12, "/app/assets/logotext.png",
			props.Rect{
				Center:  true,
				Percent: 100,
			}),
	)
}

func addSuratTugasNumber(m core.Maroto, suratTugas models.SuratTugas) {
	m.AddRow(15)

	m.AddRow(5,
		col.New(12).Add(
			text.New("SURAT TUGAS", props.Text{
				Align:  align.Center,
				Style:  fontstyle.Bold,
				Family: fontfamily.Courier,
				Size:   14,
			}),
			text.New(suratTugas.DocumentNumber, props.Text{
				Align:  align.Center,
				Family: fontfamily.Courier,
				Size:   12,
				Top:    5,
			}),
		),
	)
}

func addSuratTugasBody(m core.Maroto, suratTugas models.SuratTugas) {
	m.AddRow(15)

	// Variables for Assignor

	alamatAssignor := ""
	jabatanAssignor := ""

	// Conditional logic for Assignor (Pemberi Tugas)
	switch suratTugas.Assignor {
	case "Harto Joman":
		alamatAssignor = "Jl. Teh No. 3C"
		jabatanAssignor = "Manager Import"
	case "Martin Joman":
		alamatAssignor = "Jl. Teh No. 3C"
		jabatanAssignor = "Manager Import"
	}

	m.AddRow(20,
		col.New(12).Add(
			text.New("Jakarta, "+suratTugas.Date.Format("02 January 2006"), props.Text{
				Align:  align.Left,
				Family: fontfamily.Courier,
				Size:   11,
			}),
			text.New("Saya yang bertanda tangan di bawah ini :", props.Text{
				Align:  align.Left,
				Family: fontfamily.Courier,
				Size:   11,
				Top:    5,
			}),
			text.New("Nama    : "+suratTugas.Assignor, props.Text{
				Align:  align.Left,
				Family: fontfamily.Courier,
				Size:   11,
				Top:    10,
				Left:   10,
			}),
			text.New("Alamat  : "+alamatAssignor, props.Text{
				Align:  align.Left,
				Family: fontfamily.Courier,
				Size:   11,
				Top:    15,
				Left:   10,
			}),
			text.New("Jabatan : "+jabatanAssignor, props.Text{
				Align:  align.Left,
				Family: fontfamily.Courier,
				Size:   11,
				Top:    20,
				Left:   10,
			}),
		),
	)

	m.AddRow(10)

	// Variables for Assignee
	nikAssignee := ""
	alamatAssignee := ""
	jabatanAssignee := ""

	// Conditional logic for Assignee (Penerima Tugas)
	switch suratTugas.Assignee {
	case "Indra Abadi":
		nikAssignee = "3212190802950002"
		alamatAssignee = "Ds Arahan Lor, Indramayu"
		jabatanAssignee = "Staff Operasional"
	case "Sigit Alifiyah":
		nikAssignee = "3212190405060001"
		alamatAssignee = "Desa Arahan Kidul, Indramayu"
		jabatanAssignee = "Staff Operasional"
	}
	m.AddRow(30,
		col.New(12).Add(
			text.New("Selanjutnya disebut sebagai pemberi tugas.", props.Text{
				Align:  align.Left,
				Family: fontfamily.Courier,
				Size:   11,
			}),
			text.New("Dengan ini menyatakan memberikan tugas kepada :", props.Text{
				Align:  align.Left,
				Family: fontfamily.Courier,
				Size:   11,
				Top:    5,
			}),
			text.New("NIK     : "+nikAssignee, props.Text{
				Align:  align.Left,
				Family: fontfamily.Courier,
				Size:   11,
				Top:    10,
				Left:   10,
			}),
			text.New("Nama    : "+suratTugas.Assignee, props.Text{
				Align:  align.Left,
				Family: fontfamily.Courier,
				Size:   11,
				Top:    15,
				Left:   10,
			}),
			text.New("Alamat  : "+alamatAssignee, props.Text{
				Align:  align.Left,
				Family: fontfamily.Courier,
				Size:   11,
				Top:    20,
				Left:   10,
			}),
			text.New("Jabatan : "+jabatanAssignee, props.Text{
				Align:  align.Left,
				Family: fontfamily.Courier,
				Size:   11,
				Top:    25,
				Left:   10,
			}),
		),
	)

	m.AddRow(5)

	m.AddRow(20,
		col.New(12).Add(
			text.New("Selanjutnya disebut sebagai penerima tugas.", props.Text{
				Align:  align.Left,
				Family: fontfamily.Courier,
				Size:   11,
			}),
			text.New("Adapun surat ini akan dipergunakan untuk proses DO", props.Text{
				Align:  align.Left,
				Family: fontfamily.Courier,
				Size:   11,
				Top:    5,
			}),
			text.New("Demikian surat ini dibuat agar dipergunakan sebagaimana perlunya.", props.Text{
				Align:  align.Left,
				Family: fontfamily.Courier,
				Size:   11,
				Top:    10,
			}),
		),
	)
}

func addSuratTugasFooter(m core.Maroto, suratTugas models.SuratTugas) {
	m.AddRow(20)

	m.AddRow(50,
		col.New(6).Add(
			text.New("Penerima Tugas", props.Text{
				Align:  align.Center,
				Size:   12,
				Family: fontfamily.Courier,
			}),
			signature.New(suratTugas.Assignee, props.Signature{
				FontSize:   12,
				FontFamily: fontfamily.Courier,
			}),
		),
		col.New(6).Add(
			text.New("Pemberi Tugas", props.Text{
				Align:  align.Center,
				Size:   12,
				Family: fontfamily.Courier,
			}),
			signature.New(suratTugas.Assignor, props.Signature{
				FontSize:   12,
				FontFamily: fontfamily.Courier,
			}),
		),
	)

}
