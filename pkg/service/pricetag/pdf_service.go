package pricetag

import (
	"bytes"
	"fmt"
	"io"
	"strconv"

	"github.com/signintech/gopdf"
)

// pdfService implements Service generating PDF price tags.
type pdfService struct {
}

// NewPDFService constructs new pdfService.
func NewPDFService() Service {
	return &pdfService{}
}

// Generate generates new PDF of a price tag using given barcode, product name and cost.
func (s *pdfService) Generate(barcode, name string, cost int64) (io.Reader, error) {
	const (
		width  = 595
		height = 420

		barcodeX = 60
		barcodeY = 88
		nameX    = 60
		nameY    = 199
		costX    = 460
		costY    = 313
	)

	var pdf gopdf.GoPdf
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: width, H: height}}) //595.28, 841.89 = A4

	pdf.AddPage()

	if err := pdf.AddTTFFont("wts11", "assets/fonts/wts11.ttf"); err != nil {
		return nil, fmt.Errorf("pdf.AddTTFFont: %w", err)
	}

	if err := pdf.SetFont("wts11", "", 20); err != nil {
		return nil, fmt.Errorf("pdf.SetFont: %w", err)
	}

	pdf.SetLineWidth(0.1)
	pdf.SetFillColor(0, 0, 0)

	pdf.SetXY(1, 1)

	// Import page 1
	tpl1 := pdf.ImportPage("assets/pdf/price_tag_template.pdf", 1, "/MediaBox")

	// Draw pdf onto page
	pdf.UseImportedTemplate(tpl1, 1, 1, width, height)
	pdf.SetXY(barcodeX, barcodeY)
	if err := pdf.Cell(nil, barcode); err != nil {
		return nil, fmt.Errorf("pdf.Cell(barcode): %w", err)
	}
	pdf.SetXY(nameX, nameY)
	if err := pdf.Cell(nil, name); err != nil {
		return nil, fmt.Errorf("pdf.Cell(name): %w", err)
	}
	pdf.SetXY(costX, costY)
	if err := pdf.Cell(nil, strconv.FormatInt(cost, 10)); err != nil {
		return nil, fmt.Errorf("pdf.Cell(cost): %w", err)
	}

	buf := bytes.NewBuffer(nil)
	if _, err := pdf.WriteTo(buf); err != nil {
		return nil, fmt.Errorf("pdf.WriteTo: %w", err)
	}

	return buf, nil
}
