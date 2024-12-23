package helper

import (
	"fmt"
	"log"
	"server/models"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
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
	err := m.RegisterHeader(addHeader())
	if err != nil {
		log.Fatal(err.Error())
	}

	// Footer
	if err = m.RegisterFooter(addFooter()); err != nil {
		log.Fatal(err.Error())
	}

	// Invoice Number
	addInvoiceNumber(m, invoice)
	// Invoice Detail
	// addInvoiceDetail(m, invoice)
	// // Item List
	// addItemList(m, invoice)
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

func addInvoiceNumber(m core.Maroto, invoice models.Invoice) {
	m.AddRow(5)
	m.AddRow(10,
		text.NewCol(6, "Reff No : "+invoice.InvoiceNumber,
			props.Text{
				Align: align.Left,
				Size:  12,
				Style: fontstyle.Bold,
			}))
	m.AddRow(5)
}
