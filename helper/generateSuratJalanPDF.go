package helper

import (
	"fmt"
	"log"
	"server/models"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/list"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/signature"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
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

	if err = m.RegisterFooter(addSuratJalanFooter(suratJalan)); err != nil {
		log.Println("Error generating PDF Footer:", err.Error())
	}

	addSuratJalanNumber(m, suratJalan)
	addSuratJalanBody(m, suratJalan)
	addSuratJalanItemList(m, suratJalan)

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

func addSuratJalanNumber(m core.Maroto, suratJalan models.SuratJalan) {
	m.AddRow(10)

	m.AddRow(15,
		col.New(12).Add(
			text.New("SURAT JALAN", props.Text{
				Align:  align.Center,
				Style:  fontstyle.Bold,
				Family: fontfamily.Courier,
				Size:   14,
			}),
			text.New(suratJalan.DocumentNumber, props.Text{
				Align:  align.Center,
				Family: fontfamily.Courier,
				Size:   12,
				Top:    5,
			}),
		),
	)

}

func addSuratJalanBody(m core.Maroto, suratJalan models.SuratJalan) {
	m.AddRow(5, line.NewCol(12, props.Line{SizePercent: 100}))

	m.AddRow(20,
		col.New(12).Add(
			text.New("Nama Penerima : "+suratJalan.Recipient, props.Text{
				Align:  align.Left,
				Family: fontfamily.Courier,
				Size:   11,
			}),
			text.New("Alamat        : "+suratJalan.Address, props.Text{
				Align:  align.Left,
				Family: fontfamily.Courier,
				Size:   11,
				Top:    5,
			}),
			text.New("Bersama dengan ini mohon diterima barang sebagai berikut :", props.Text{
				Align:  align.Left,
				Family: fontfamily.Courier,
				Size:   11,
				Top:    15,
			}),
		),
	)

	m.AddRow(5)
}

type SuratJalanItem struct {
	ItemName string
	Type     string
	Colly    string
	Volume   string
	Note     string
}

func (o SuratJalanItem) GetHeader() core.Row {
	rowStyle := &props.Cell{
		BorderColor:     &props.BlackColor,
		BorderType:      border.Full,
		LineStyle:       linestyle.Solid,
		BackgroundColor: &props.Color{220, 220, 220},
	}

	return row.New().Add(
		text.NewCol(2, "Type", props.Text{Style: fontstyle.Bold, Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2, Family: fontfamily.Courier}).WithStyle(rowStyle),
		text.NewCol(2, "Item", props.Text{Style: fontstyle.Bold, Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2, Family: fontfamily.Courier}).WithStyle(rowStyle),
		text.NewCol(2, "Colly/Pkgs", props.Text{Style: fontstyle.Bold, Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2, Family: fontfamily.Courier}).WithStyle(rowStyle),
		text.NewCol(3, "Volume", props.Text{Style: fontstyle.Bold, Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2, Family: fontfamily.Courier}).WithStyle(rowStyle),
		text.NewCol(3, "Note", props.Text{Style: fontstyle.Bold, Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2, Family: fontfamily.Courier}).WithStyle(rowStyle),
	)
}

func (o SuratJalanItem) GetContent(i int) core.Row {
	rowStyle := &props.Cell{
		BorderColor: &props.BlackColor,
		BorderType:  border.Full,
		LineStyle:   linestyle.Solid,
	}

	r := row.New().Add(
		text.NewCol(2, o.Type, props.Text{Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2, Family: fontfamily.Courier}).WithStyle(rowStyle),
		text.NewCol(2, o.ItemName, props.Text{Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2, Family: fontfamily.Courier}).WithStyle(rowStyle),
		text.NewCol(2, o.Colly, props.Text{Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2, Family: fontfamily.Courier}).WithStyle(rowStyle),
		text.NewCol(3, o.Volume, props.Text{Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2, Family: fontfamily.Courier}).WithStyle(rowStyle),
		text.NewCol(3, o.Note, props.Text{Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2, Family: fontfamily.Courier}).WithStyle(rowStyle),
	)

	return r
}

func getSuratJalanObject(suratJalan models.SuratJalan) []SuratJalanItem {
	var items []SuratJalanItem

	for _, item := range suratJalan.Items {
		items = append(items, SuratJalanItem{
			ItemName: item.ItemName,
			Type:     item.Type,
			Colly: func() string {
				if item.Colly != nil {
					return string(FormatThousandSeparatorFloat(*item.Colly))
				}
				return "-"
			}(),
			Volume: func() string {
				if item.Volume != nil {
					return string(FormatThousandSeparatorFloat(*item.Volume)) + " " + item.Unit
				}
				return "-"
			}(),
			Note: func() string {
				if item.Note != nil {
					return *item.Note
				}
				return ""
			}(),
		})
	}

	return items
}

func addSuratJalanItemList(m core.Maroto, suratJalan models.SuratJalan) {
	rows, err := list.Build[SuratJalanItem](getSuratJalanObject(suratJalan))
	if err != nil {
		log.Println("Error generating PDF:", err.Error())
	}
	m.AddRows(rows...)

	m.AddRow(10)

	m.AddRow(6,
		text.NewCol(12, "Pastikan cargo yang disebutkan diatas telah diterima dalam keadaan baik. Kami tidak menerima komplain setelah pengiriman", props.Text{
			Align:  align.Left,
			Family: fontfamily.Courier,
			Size:   11,
		}),
	)
}

func addSuratJalanFooter(suratJalan models.SuratJalan) core.Row {
	return row.New().Add(
		col.New(6).Add(
			text.New("", props.Text{
				Top: 50,
			}),
			signature.New("PENGIRIM", props.Signature{
				FontSize:   12,
				FontFamily: fontfamily.Courier,
			}),
		),
		col.New(6).Add(
			text.New("Jakarta, "+suratJalan.Date.Format("02 January 2006"), props.Text{
				Align:  align.Center,
				Size:   12,
				Family: fontfamily.Courier,
				Top:    0,
			}),
			signature.New("PENERIMA", props.Signature{
				FontSize:   12,
				FontFamily: fontfamily.Courier,
			}),
		),
	)
}
