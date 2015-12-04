package main

import (
	"../"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", example.FontDir())
	pdf.AddFont("Calligrapher", "", "calligra.json")
	pdf.AddPage()
	pdf.SetFont("Calligrapher", "", 35)
	pdf.Cell(0, 10, "Enjoy new fonts with FPDF!")
	fileStr := example.Filename("Fpdf_AddFont")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
}
