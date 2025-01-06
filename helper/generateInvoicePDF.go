package helper

import (
	"fmt"
	"log"
	"server/models"
	"strings"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/list"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/signature"
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
	ntw "moul.io/number-to-words"
)

func GenerateInvoicePDF(invoice models.Invoice) {
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
	err := m.RegisterHeader(addInvoiceHeader())
	if err != nil {
		log.Fatal(err.Error())
	}

	// Footer
	if err = m.RegisterFooter(addInvoiceFooter(invoice)); err != nil {
		log.Fatal(err.Error())
	}

	// addInvoiceHeader(m)
	addInvoiceNumber(m, invoice)
	customerSection(m, invoice)
	priceSpellingSection(m, invoice)
	addInvoiceDetail(m, invoice)
	addInvoiceItemList(m, invoice)
	// // Item List
	// // Note and Signature
	// addNoteAndSignature(m, invoice)

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	safeInvoiceNumber := SanitizeFilename(invoice.InvoiceNumber)

	// Use the sanitized InvoiceNumber to name the file
	filePath := fmt.Sprintf("pdf/invoice/%s.pdf", safeInvoiceNumber)

	err = document.Save(filePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("PDF saved successfully")
}

// func addInvoiceHeader(m core.Maroto) {
// 	m.AddRow(40, image.NewFromFileCol(12, "assets/logotext.png", props.Rect{
// 		Percent: 100,
// 		Center:  true,
// 	}))
// 	// m.AddRow(5,
// 	// 	text.NewCol(12, "INTERNATIONAL SEA/AIR FREIGHT FORWARDERS",
// 	// 		props.Text{
// 	// 			Align: align.Center,
// 	// 			Size:  12,
// 	// 			Style: fontstyle.Bold,
// 	// 		}))
// 	// m.AddRow(5,
// 	// 	text.NewCol(12, "(SEA-AIR, PROJECT CARGO SHIP & AIRCRAFT CHARTER)",
// 	// 		props.Text{
// 	// 			Align: align.Center,
// 	// 			Size:  12,
// 	// 			Style: fontstyle.Bold,
// 	// 		}))
// }

func addInvoiceHeader() core.Row {
	return row.New(40).Add(
		image.NewFromFileCol(12, "assets/logotext.png",
			props.Rect{
				Center:  true,
				Percent: 100,
			}),
	)
}

func addInvoiceNumber(m core.Maroto, invoice models.Invoice) {
	m.AddRow(5)

	m.AddRow(20,
		col.New(4).Add(
			text.New("No Invoice", props.Text{
				Align: align.Left,
				Style: fontstyle.Bold,
			}),
			text.New(invoice.InvoiceNumber, props.Text{
				Top:   4,
				Align: align.Left,
				Style: fontstyle.Bold,
			}),
		),
		col.New(4).Add(
			text.New("FAKTUR", props.Text{
				Align: align.Center,
				Style: fontstyle.Bold,
			}),
			text.New("INVOICE", props.Text{
				Top:   4,
				Align: align.Center,
				Style: fontstyle.BoldItalic,
			}),
		),
		col.New(4).Add(
			text.New("JL. Teh No. 3 - C", props.Text{
				Align: align.Right,
				Style: fontstyle.Bold,
			}),
			text.New("TELP. : 021-6927181", props.Text{
				Top:   4,
				Align: align.Right,
				Style: fontstyle.Bold,
			}),
			text.New("FAX : (62-21) 6916336", props.Text{
				Top:   8,
				Align: align.Right,
				Style: fontstyle.Bold,
			}),
			text.New("JAKARTA 11110 - INDONESIA", props.Text{
				Top:   12,
				Align: align.Right,
				Style: fontstyle.Bold,
			}),
		),
	)

	m.AddRow(5, line.NewCol(12, props.Line{SizePercent: 100}))
}

func customerSection(m core.Maroto, invoice models.Invoice) {
	m.AddRow(15,
		col.New(6).Add(
			text.New("Sudah diterima dari :"),
			text.New("Received from", props.Text{
				Top:   5,
				Style: fontstyle.Italic,
			}),
			text.New(invoice.Customer.Name, props.Text{
				Left: 40,
			}),
		),
		col.New(6).Add(
			text.New("Alamat :"),
			text.New("Address", props.Text{
				Top:   5,
				Style: fontstyle.Italic,
			}),
			text.New(invoice.Customer.Address, props.Text{
				Left: 20,
			}),
		))
	m.AddRow(5, line.NewCol(12, props.Line{SizePercent: 100}))

}

// Function to calculate total in IDR and spell it
func CalculateTotalAndSpell(invoiceItems []models.InvoiceItem) string {
	var totalIDR float64

	// Loop through invoice items to calculate total in IDR
	for _, item := range invoiceItems {
		itemTotal := item.Price * float64(item.Quantity)
		if strings.ToUpper(item.Currency) != "IDR" {
			// Convert to IDR
			itemTotal *= *item.Kurs
		}
		totalIDR += itemTotal
	}

	// Convert total to words in Indonesian
	spelledTotal := ntw.IntegerToIDID(int(totalIDR)) + " rupiah"
	return strings.ToUpper(spelledTotal)
}

func priceSpellingSection(m core.Maroto, invoice models.Invoice) {
	m.AddRow(15,
		col.New(12).Add(
			text.New("Uang sejumlah :"),
			text.New("The sum of", props.Text{
				Top:   5,
				Style: fontstyle.Italic,
			}),
			// Place the spelling here
			text.New("# "+CalculateTotalAndSpell(invoice.InvoiceItems)+" #", props.Text{
				Left:  40,
				Style: fontstyle.Bold,
				Size:  12,
			}),
		))
	m.AddRow(5, line.NewCol(12, props.Line{SizePercent: 100}))

}

func addInvoiceDetail(m core.Maroto, invoice models.Invoice) {
	m.AddRow(40,
		col.New(3).Add(
			text.New("Untuk pembayaran :"),
			text.New("In payment of", props.Text{
				Top:   5,
				Style: fontstyle.Italic,
			}),
		),
		col.New(6).Add(
			text.New(invoice.Type),
			text.New("BL/AWB : "+invoice.BLAWB, props.Text{
				Top: 5,
			}),
			text.New("No Aju : "+invoice.AJU, props.Text{
				Top: 10,
			}),
			text.New("Shipper : "+invoice.Shipper.Name, props.Text{
				Top: 15,
			}),
			text.New("Remarks : "+invoice.ShippingMarks, props.Text{
				Top: 20,
			}),
			text.New("POL : "+invoice.PortOfLoading.PortName, props.Text{
				Top: 25,
			}),
			text.New("POD : "+invoice.PortOfDischarge.PortName, props.Text{
				Top: 30,
			}),
		))

}

type InvoiceItem struct {
	Item     string
	Qty      string
	Kurs     string
	Price    string
	SubTotal string
}

func (o InvoiceItem) GetHeader() core.Row {
	rowStyle := &props.Cell{
		BorderColor:     &props.BlackColor,
		BorderType:      border.Full,
		LineStyle:       linestyle.Solid,
		BackgroundColor: &props.Color{220, 220, 220},
	}

	return row.New().Add(
		text.NewCol(2, "Item", props.Text{Style: fontstyle.Bold, Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2}).WithStyle(rowStyle),
		text.NewCol(2, "Qty", props.Text{Style: fontstyle.Bold, Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2}).WithStyle(rowStyle),
		text.NewCol(2, "Kurs", props.Text{Style: fontstyle.Bold, Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2}).WithStyle(rowStyle),
		text.NewCol(3, "Price", props.Text{Style: fontstyle.Bold, Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2}).WithStyle(rowStyle),
		text.NewCol(3, "Sub Total", props.Text{Style: fontstyle.Bold, Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2}).WithStyle(rowStyle),
	)
}

func (o InvoiceItem) GetContent(i int) core.Row {
	rowStyle := &props.Cell{
		BorderColor: &props.BlackColor,
		BorderType:  border.Full,
		LineStyle:   linestyle.Solid,
	}

	r := row.New().Add(
		text.NewCol(2, o.Item, props.Text{Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2}).WithStyle(rowStyle),
		text.NewCol(2, o.Qty, props.Text{Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2}).WithStyle(rowStyle),
		text.NewCol(2, o.Kurs, props.Text{Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2}).WithStyle(rowStyle),
		text.NewCol(3, o.Price, props.Text{Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2}).WithStyle(rowStyle),
		text.NewCol(3, o.SubTotal, props.Text{Size: 12, Top: 2, Left: 2, Bottom: 2, Right: 2}).WithStyle(rowStyle),
	)

	return r
}

func getInvoiceObject(invoice models.Invoice) []InvoiceItem {
	var items []InvoiceItem

	for _, item := range invoice.InvoiceItems {
		var subTotal float64
		if item.Currency != "IDR" {
			subTotal = float64(item.Quantity) * item.Price * *item.Kurs
		} else {
			subTotal = float64(item.Quantity) * item.Price
		}

		items = append(items, InvoiceItem{
			Item: item.ItemName,
			Qty:  FormatThousandSeparatorInt(item.Quantity) + " " + item.Unit,
			Kurs: func() string {
				if item.Kurs != nil {
					return "Rp " + string(FormatThousandSeparatorFloat(*item.Kurs))
				}
				return "-"
			}(),
			Price:    item.Currency + " " + FormatThousandSeparatorFloat(item.Price),
			SubTotal: "Rp " + FormatThousandSeparatorFloat(subTotal),
		})
	}

	return items
}

func addInvoiceItemList(m core.Maroto, invoice models.Invoice) {
	rows, err := list.Build[InvoiceItem](getInvoiceObject(invoice))
	if err != nil {
		log.Fatal(err.Error())
	}
	m.AddRows(rows...)

	var totalIDR float64

	for _, item := range invoice.InvoiceItems {
		var subTotal float64

		if item.Currency != "IDR" && item.Kurs != nil {
			subTotal = float64(item.Quantity) * item.Price * *item.Kurs
		} else {
			subTotal = float64(item.Quantity) * item.Price
		}
		totalIDR += subTotal
	}

	m.AddRow(5)

	m.AddRow(6,
		text.NewCol(9, "SUBTOTAL", props.Text{
			Align: align.Left,
			Style: fontstyle.Bold,
			Size:  14,
		}),
		text.NewCol(3, "Rp "+FormatThousandSeparatorFloat(totalIDR), props.Text{
			Align: align.Left,
			Style: fontstyle.Bold,
			Size:  14,
		}),
	)

	m.AddRow(6,
		text.NewCol(9, "VAT", props.Text{
			Align: align.Left,
			Size:  14,
		}),
		text.NewCol(3, "Rp 0", props.Text{
			Align: align.Left,
			Size:  14,
		}),
	)

	m.AddRow(1)

	m.AddRow(1, line.NewCol(12, props.Line{SizePercent: 100}))

	m.AddRow(6,
		text.NewCol(9, "TOTAL", props.Text{
			Align: align.Left,
			Style: fontstyle.Bold,
			Size:  14,
		}),
		text.NewCol(3, "Rp "+FormatThousandSeparatorFloat(totalIDR), props.Text{
			Align: align.Left,
			Style: fontstyle.Bold,
			Size:  14,
		}),
	)

}

func addInvoiceFooter(invoice models.Invoice) core.Row {
	return row.New().Add(
		col.New(8).Add(
			text.New("INFORMASI PEMBAYARAN / PAYMENT DETAILS :"),
			text.New("NAMA BANK / BANK DETAILS : BANK CENTRAL ASIA - CENGKEH", props.Text{
				Top:   5,
				Style: fontstyle.Bold,
			}),
			text.New("NAMA AKUN / ACCOUNT NAME : PT RADIX SAGARA JOMANTARA", props.Text{
				Top:   10,
				Style: fontstyle.Bold,
			}),
			text.New("1. REKENING RUPIAH / INDONESIAN RUPIAH ACCOUNT", props.Text{
				Top: 15,
			}),
			text.New("NOMOR REKENING / ACCOUNT NUMBER : 240-303-0023", props.Text{
				Top:   20,
				Style: fontstyle.Bold,
			}),
			text.New("2. REKENING MATA UANG ASING / FOREIGN CURRENCY ACCOUNT", props.Text{
				Top: 25,
			}),
			text.New("NOMOR REKENING /ACCOUNT NUMBER : 240-036-3121", props.Text{
				Top:   30,
				Style: fontstyle.Bold,
			}),
			text.New("SWIFT CODE : CENAIDJA", props.Text{
				Top:   35,
				Style: fontstyle.Bold,
			}),
		),
		col.New(4).Add(
			text.New("Jakarta, "+invoice.InvoiceDate.Format("02 January 2006"), props.Text{
				Align: align.Center,
			}),
			signature.New("HARTO JOMAN"),
		),
	)
}
