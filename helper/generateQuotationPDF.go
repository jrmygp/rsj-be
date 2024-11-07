package helper

import (
	"fmt"
	"log"
	"regexp"
	"server/models"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/list"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func sanitizeFilename(filename string) string {
	re := regexp.MustCompile(`[\/:*?"<>|]`)
	return re.ReplaceAllString(filename, "_")
}

func GenerateQuotationPDF(quotation models.Quotation) {
	cfg := config.NewBuilder().
		WithOrientation(orientation.Vertical).
		WithPageSize(pagesize.A4).
		WithLeftMargin(15).
		WithTopMargin(5).
		WithRightMargin(15).
		WithBottomMargin(5).
		Build()
	m := maroto.New(cfg)

	// Header
	addHeader(m)
	// Invoice Number
	addInvoiceNumber(m, quotation)
	// Invoice Detail
	addInvoiceDetail(m, quotation)
	// Item List
	addItemList(m, quotation)
	// // Footer
	// addFooter(m)

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	safeQuotationNumber := sanitizeFilename(quotation.QuotationNumber)

	// Use the sanitized QuotationNumber to name the file
	filePath := fmt.Sprintf("pdf/quotation/invoice-%s.pdf", safeQuotationNumber)

	err = document.Save(filePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("PDF saved successfully")
}

func addHeader(m core.Maroto) {
	m.AddRow(50,
		image.NewFromFileCol(12, "assets/logotext.png",
			props.Rect{
				Center:  true,
				Percent: 75,
			}))
}

func addInvoiceNumber(m core.Maroto, quotation models.Quotation) {
	m.AddRow(10,
		text.NewCol(6, "Reff No : "+quotation.QuotationNumber,
			props.Text{
				Align: align.Left,
				Size:  10,
				Style: fontstyle.Bold,
			}))
}

func addInvoiceDetail(m core.Maroto, quotation models.Quotation) {
	m.AddRow(5,
		text.NewCol(6, "To : "+quotation.Customer.Name,
			props.Text{
				Align: align.Left,
				Size:  10,
			}),
		text.NewCol(6, "Service : "+quotation.Service,
			props.Text{
				Align: align.Left,
				Size:  10,
			}),
	)
	m.AddRow(5,
		text.NewCol(6, "Commodity : "+quotation.Commodity,
			props.Text{
				Align: align.Left,
				Size:  10,
			}),
		text.NewCol(6, "Meas/Weight : "+fmt.Sprintf("%d kg / %d CBM", quotation.Weight, quotation.Volume),
			props.Text{
				Align: align.Left,
				Size:  10,
			}),
	)
	m.AddRow(5,
		text.NewCol(6, "Rate Validity : "+quotation.RateValidity.Format("02 January 2006"),
			props.Text{
				Align: align.Left,
				Size:  10,
			}),
		text.NewCol(6, "Sales Name : "+quotation.Sales.Name,
			props.Text{
				Align: align.Left,
				Size:  10,
			}),
	)
	m.AddRow(5,
		text.NewCol(6, "Term of Delivery : "+quotation.ShippingTerm,
			props.Text{
				Align: align.Left,
				Size:  10,
			}),
		text.NewCol(6, "Phone : "+quotation.Sales.PhoneNumber,
			props.Text{
				Align: align.Left,
				Size:  10,
			}),
	)
	m.AddRow(5,
		text.NewCol(6, "Port of Loading : "+quotation.PortOfLoading.PortName,
			props.Text{
				Align: align.Left,
				Size:  10,
			}),
		text.NewCol(6, "Email : "+quotation.Sales.Email,
			props.Text{
				Align: align.Left,
				Size:  10,
			}),
	)
	m.AddRow(5,
		text.NewCol(6, "Port of Discharge : "+quotation.PortOfDischarge.PortName,
			props.Text{
				Align: align.Left,
				Size:  10,
			}),
		text.NewCol(6, "Payment Terms : "+quotation.PaymentTerm,
			props.Text{
				Align: align.Left,
				Size:  10,
			}),
	)
	m.AddRow(10)
}

type QuotationItem struct {
	Item     string
	Qty      string
	Unit     string
	Price    string
	SubTotal string
	Remarks  string
}

func (o QuotationItem) GetHeader() core.Row {
	return row.New(5).Add(
		text.NewCol(3, "Item", props.Text{Style: fontstyle.Bold}),
		text.NewCol(1, "Qty", props.Text{Style: fontstyle.Bold}),
		text.NewCol(1, "Unit", props.Text{Style: fontstyle.Bold}),
		text.NewCol(2, "Price", props.Text{Style: fontstyle.Bold}),
		text.NewCol(3, "Sub Total", props.Text{Style: fontstyle.Bold}),
		text.NewCol(2, "Remarks", props.Text{Style: fontstyle.Bold}),
	)
}

func (o QuotationItem) GetContent(i int) core.Row {
	r := row.New(5).Add(
		text.NewCol(3, o.Item),
		text.NewCol(1, o.Qty),
		text.NewCol(1, o.Unit),
		text.NewCol(2, o.Price),
		text.NewCol(3, o.SubTotal),
		text.NewCol(2, o.Remarks),
	)

	if i%2 == 0 {
		r.WithStyle(&props.Cell{
			BackgroundColor: &props.Color{Red: 240, Green: 240, Blue: 240},
		})
	}

	return r
}

func getObject(quotation models.Quotation) []QuotationItem {
	var items []QuotationItem

	for i := 0; i < len(quotation.ListCharges); i++ {
		charge := quotation.ListCharges[i]

		subTotal := float64(charge.Quantity) * charge.Price

		items = append(items, QuotationItem{
			Item:     charge.ItemName,
			Qty:      FormatThousandSeparatorInt(charge.Quantity),
			Unit:     charge.Unit,
			Price:    "Rp " + FormatThousandSeparatorFloat(charge.Price),
			SubTotal: "Rp " + FormatThousandSeparatorFloat(subTotal),
			Remarks: func() string {
				if charge.Note != nil {
					return *charge.Note
				}
				return "" // Return empty string if Note is nil
			}(),
		})
	}

	return items
}

func addItemList(m core.Maroto, quotation models.Quotation) {
	rows, err := list.Build[QuotationItem](getObject(quotation))
	if err != nil {
		log.Fatal(err.Error())
	}
	m.AddRows(rows...)
}
