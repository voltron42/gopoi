package main

import (
	"../"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetProtection(gofpdf.CnProtectPrint, "123", "abc")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	pdf.Write(10, "Password-protected.")
	fileStr := example.Filename("Fpdf_SetProtection")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
}
