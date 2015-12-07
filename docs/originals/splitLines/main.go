package main

import (
	"../"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	const (
		fontPtSize = 18.0
		wd         = 100.0
	)
	pdf := gofpdf.New("P", "mm", "A4", "") // A4 210.0 x 297.0
	pdf.SetFont("Times", "", fontPtSize)
	_, lineHt := pdf.GetFontSize()
	pdf.AddPage()
	pdf.SetMargins(10, 10, 10)
	lines := pdf.SplitLines([]byte(lorem()), wd)
	ht := float64(len(lines)) * lineHt
	y := (297.0 - ht) / 2.0
	pdf.SetDrawColor(128, 128, 128)
	pdf.SetFillColor(255, 255, 210)
	x := (210.0 - (wd + 40.0)) / 2.0
	pdf.Rect(x, y-20.0, wd+40.0, ht+40.0, "FD")
	pdf.SetY(y)
	for _, line := range lines {
		pdf.CellFormat(190.0, lineHt, string(line), "", 1, "C", false, 0, "")
	}
	fileStr := example.Filename("Fpdf_Splitlines")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
}

func lorem() string {
	// TODO --
	return ""
}
