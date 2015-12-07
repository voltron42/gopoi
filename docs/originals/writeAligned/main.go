package main

import (
	"../"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", example.FontDir())
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", 12)
	pdf.WriteAligned(0, 35, "This text is the default alignment, Left", "")
	pdf.Ln(35)
	pdf.WriteAligned(0, 35, "This text is aligned Left", "L")
	pdf.Ln(35)
	pdf.WriteAligned(0, 35, "This text is aligned Center", "C")
	pdf.Ln(35)
	pdf.WriteAligned(0, 35, "This text is aligned Right", "R")
	fileStr := example.Filename("Fpdf_WriteAligned")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
}
