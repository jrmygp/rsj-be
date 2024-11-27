package helper

import (
	"fmt"
	"log"
	"regexp"
	"server/models"
	"sort"
	"strings"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/list"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
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
		WithTopMargin(0).
		WithRightMargin(15).
		WithBottomMargin(5).
		Build()
	m := maroto.New(cfg)

	// Header
	err := m.RegisterHeader(addHeader())
	if err != nil {
		log.Fatal(err.Error())
	}

	// Footer
	if err = m.RegisterFooter(addFooter()); err != nil {
		log.Fatal(err.Error())
	}

	// Invoice Number
	addInvoiceNumber(m, quotation)
	// Invoice Detail
	addInvoiceDetail(m, quotation)
	// Item List
	addItemList(m, quotation)
	// Note and Signature
	addNoteAndSignature(m, quotation)

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

func addHeader() core.Row {
	return row.New(50).Add(
		image.NewFromFileCol(12, "assets/logotext.png",
			props.Rect{
				Center:  true,
				Percent: 75,
			}),
	)
}

func addInvoiceNumber(m core.Maroto, quotation models.Quotation) {
	m.AddRow(5)
	m.AddRow(10,
		text.NewCol(6, "Reff No : "+quotation.QuotationNumber,
			props.Text{
				Align: align.Left,
				Size:  12,
				Style: fontstyle.Bold,
			}))
	m.AddRow(5)
}

func addInvoiceDetail(m core.Maroto, quotation models.Quotation) {
	m.AddRow(5,
		text.NewCol(6, "To : "+quotation.Customer.Name,
			props.Text{
				Align: align.Left,
				Size:  12,
			}),
		text.NewCol(6, "Service : "+quotation.Service,
			props.Text{
				Align: align.Left,
				Size:  12,
			}),
	)
	m.AddRow(5,
		text.NewCol(6, "Commodity : "+quotation.Commodity,
			props.Text{
				Align: align.Left,
				Size:  12,
			}),
		text.NewCol(6, "Weight/Volume : "+fmt.Sprintf("%.2f kg / %.2f CBM", quotation.Weight, quotation.Volume),
			props.Text{
				Align: align.Left,
				Size:  12,
			}),
	)
	m.AddRow(5,
		text.NewCol(6, "Rate Validity : "+quotation.RateValidity.Format("02 January 2006"),
			props.Text{
				Align: align.Left,
				Size:  12,
			}),
		text.NewCol(6, "Sales Name : "+quotation.Sales.Name,
			props.Text{
				Align: align.Left,
				Size:  12,
			}),
	)
	m.AddRow(5,
		text.NewCol(6, "Term of Delivery : "+quotation.ShippingTerm,
			props.Text{
				Align: align.Left,
				Size:  12,
			}),
		text.NewCol(6, "Phone : "+quotation.Sales.PhoneNumber,
			props.Text{
				Align: align.Left,
				Size:  12,
			}),
	)
	m.AddRow(5,
		text.NewCol(6, "Port of Loading : "+quotation.PortOfLoading.PortName,
			props.Text{
				Align: align.Left,
				Size:  12,
			}),
		text.NewCol(6, "Email : "+quotation.Sales.Email,
			props.Text{
				Align: align.Left,
				Size:  12,
			}),
	)
	m.AddRow(5,
		text.NewCol(6, "Port of Discharge : "+quotation.PortOfDischarge.PortName,
			props.Text{
				Align: align.Left,
				Size:  12,
			}),
		text.NewCol(6, "Payment Terms : "+quotation.PaymentTerm,
			props.Text{
				Align: align.Left,
				Size:  12,
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
	rowStyle := &props.Cell{
		BorderColor:     &props.BlackColor,
		BorderType:      border.Full,
		LineStyle:       linestyle.Solid,
		BackgroundColor: &props.Color{220, 220, 220},
	}

	return row.New().Add(
		text.NewCol(2, "Item", props.Text{Style: fontstyle.Bold, Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2}).WithStyle(rowStyle),
		text.NewCol(1, "Qty", props.Text{Style: fontstyle.Bold, Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2}).WithStyle(rowStyle),
		text.NewCol(1, "Unit", props.Text{Style: fontstyle.Bold, Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2}).WithStyle(rowStyle),
		text.NewCol(3, "Price", props.Text{Style: fontstyle.Bold, Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2}).WithStyle(rowStyle),
		text.NewCol(3, "Sub Total", props.Text{Style: fontstyle.Bold, Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2}).WithStyle(rowStyle),
		text.NewCol(2, "Remarks", props.Text{Style: fontstyle.Bold, Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2}).WithStyle(rowStyle),
	)
}

func (o QuotationItem) GetContent(i int) core.Row {
	rowStyle := &props.Cell{
		BorderColor: &props.BlackColor,
		BorderType:  border.Full,
		LineStyle:   linestyle.Solid,
	}

	r := row.New().Add(
		text.NewCol(2, o.Item, props.Text{Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2}).WithStyle(rowStyle),
		text.NewCol(1, o.Qty, props.Text{Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2}).WithStyle(rowStyle),
		text.NewCol(1, o.Unit, props.Text{Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2}).WithStyle(rowStyle),
		text.NewCol(3, o.Price, props.Text{Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2}).WithStyle(rowStyle),
		text.NewCol(3, o.SubTotal, props.Text{Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2}).WithStyle(rowStyle),
		text.NewCol(2, o.Remarks, props.Text{Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2}).WithStyle(rowStyle),
	)

	return r
}

func getObject(quotation models.Quotation) []QuotationItem {
	var items []QuotationItem

	for _, charge := range quotation.ListCharges {
		subTotal := float64(charge.Quantity) * charge.Price

		items = append(items, QuotationItem{
			Item: charge.ItemName,
			Qty:  FormatThousandSeparatorInt(charge.Quantity),
			Unit: func() string {
				if charge.Unit != nil {
					return *charge.Unit
				}
				return ""
			}(),
			Price:    charge.Currency + " " + FormatThousandSeparatorFloat(charge.Price),
			SubTotal: charge.Currency + " " + FormatThousandSeparatorFloat(subTotal),
			Remarks: func() string {
				if charge.Note != nil {
					return *charge.Note
				}
				return ""
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

	// Step 1: Initialize a map to hold totals for each currency
	currencyTotals := make(map[string]float64)

	// Step 2: Calculate subtotal for each item and add to the corresponding currency total
	for _, charge := range quotation.ListCharges {
		subTotal := float64(charge.Quantity) * charge.Price
		currencyTotals[charge.Currency] += subTotal
	}

	// Step 3: Separate IDR from other currencies, then sort the rest alphabetically
	var totalAmountStrings []string
	if idrTotal, exists := currencyTotals["IDR"]; exists {
		totalAmountStrings = append(totalAmountStrings, fmt.Sprintf("Rp %s", FormatThousandSeparatorFloat(idrTotal)))
		delete(currencyTotals, "IDR") // Remove IDR so we can sort other currencies separately
	}

	// Sort the remaining currencies alphabetically
	var otherCurrencies []string
	for currency := range currencyTotals {
		otherCurrencies = append(otherCurrencies, currency)
	}
	sort.Strings(otherCurrencies)

	// Format the totals for other currencies and add to the result
	for _, currency := range otherCurrencies {
		formattedTotal := fmt.Sprintf("%s %s", currency, FormatThousandSeparatorFloat(currencyTotals[currency]))
		totalAmountStrings = append(totalAmountStrings, formattedTotal)
	}

	// Join all totals into the final string
	finalTotalAmountString := strings.Join(totalAmountStrings, " + ")

	m.AddRow(5)

	m.AddRow(10,
		text.NewCol(6, "Total Amount :", props.Text{
			Align: align.Left,
			Size:  14,
			Style: fontstyle.Bold,
		}),
		text.NewCol(6, finalTotalAmountString, props.Text{
			Align: align.Right,
			Size:  14,
			Style: fontstyle.Bold,
		}),
	)

}

func addNoteAndSignature(m core.Maroto, quotation models.Quotation) {
	noteStyle := &props.Cell{
		BackgroundColor: &props.Color{255, 243, 128},
	}

	m.AddRows(
		row.New().Add(
			text.NewCol(12, "Note :",
				props.Text{
					Align:  align.Left,
					Size:   10,
					Top:    2,
					Left:   2,
					Right:  2,
					Bottom: 1,
				}).WithStyle(noteStyle),
		),
		row.New().Add(
			text.NewCol(12, quotation.Note,
				props.Text{
					Align:  align.Left,
					Size:   10,
					Left:   2,
					Right:  2,
					Bottom: 2,
				}).WithStyle(noteStyle),
		),
	)

	m.AddRow(5)

	m.AddRow(100,
		image.NewFromFileCol(12, "assets/signpdf.jpg",
			props.Rect{
				Center:  true,
				Percent: 100,
			}),
	)

	// m.AddRow(5,
	// 	text.NewCol(12, "1.Semua biaya dan tarif yang tercantum harus dibayar dalam IDR (Rupiah Indonesia) sesuai tanggal Faktur.",
	// 		props.Text{
	// 			Align: align.Left,
	// 			Size:  10,
	// 		}),
	// )

	// m.AddRow(5,
	// 	text.NewCol(12, "2.Importir ataupun Exportir wajib menyiapkan NPWP/NIB dan document pendukung lainnya.",
	// 		props.Text{
	// 			Align: align.Left,
	// 			Size:  10,
	// 		}),
	// )

	// m.AddRow(5,
	// 	text.NewCol(12, "3.Pengiriman tidak diasuransikan kecuali atas permintaan khusus dari customer.",
	// 		props.Text{
	// 			Align: align.Left,
	// 			Size:  10,
	// 		}),
	// )

	// m.AddRow(5,
	// 	text.NewCol(12, "4.Setiap Cancellation shipment wajib membayar Cancellation Fee jika ada.",
	// 		props.Text{
	// 			Align: align.Left,
	// 			Size:  10,
	// 		}),
	// )

	// m.AddRow(5,
	// 	text.NewCol(12, "5.Semua harga diatas berlaku sesuai tanggal valid yang tersedia dan tidak mengikat jika sudah melebihi.",
	// 		props.Text{
	// 			Align: align.Left,
	// 			Size:  10,
	// 		}),
	// )
	// m.AddRow(5)

	// m.AddRow(25,
	// 	text.NewCol(8, ""),
	// 	text.NewCol(4, "Approved By,",
	// 		props.Text{
	// 			Align: align.Center,
	// 			Size:  10,
	// 		}),
	// )

	// m.AddRow(5,
	// 	text.NewCol(8, ""),
	// 	signature.NewCol(4, "Martin Joman",
	// 		props.Signature{}),
	// )

	// m.AddRow(10)

}

func addFooter() core.Row {
	rowStyle := &props.Cell{
		BorderColor:     &props.Color{138, 45, 38},
		BorderType:      border.Top,
		LineStyle:       linestyle.Solid,
		BorderThickness: 1,
	}

	return row.New().Add(
		col.New(3).Add(
			text.New("021-6927181", props.Text{
				Align: align.Left,
				Top:   10,
				Size:  10,
			}),
		).WithStyle(rowStyle),
		col.New(5).Add(
			text.New("martin@radixsagarajomantara.com", props.Text{
				Align: align.Left,
				Top:   10,
				Size:  10,
			}),
		).WithStyle(rowStyle),
		col.New(4).Add(
			text.New("Jl.Teh no 3.C Tamansari Pinangsia, Jakarta Barat 11110 Indonesia", props.Text{
				Align: align.Right,
				Top:   10,
				Size:  10,
			}),
		).WithStyle(rowStyle),
	)
}
