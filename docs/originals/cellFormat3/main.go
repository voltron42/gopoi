package main

import (
	"../"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "") // A4 210.0 x 297.0
	fontSize := 16.0
	pdf.SetFont("Helvetica", "", fontSize)
	ht := pdf.PointConvert(fontSize)
	write := func(str string) {
		pdf.CellFormat(190, ht, str, "", 1, "C", false, 0, "")
		pdf.Ln(ht)
	}
	pdf.AddPage()
	htmlStr := `Until gofpdf supports UTF-8 encoded source text, source text needs ` +
		`to be specified with all special characters escaped to match the code page ` +
		`layout of the currently selected font. By default, gofdpf uses code page 1252.` +
		` See <a href="http://en.wikipedia.org/wiki/Windows-1252">Wikipedia</a> for ` +
		`a table of this layout.`
	html := pdf.HTMLBasicNew()
	html.Write(ht, htmlStr)
	pdf.Ln(2 * ht)
	write("Voix ambigu\xeb d'un c\x9cur qui au z\xe9phyr pr\xe9f\xe8re les jattes de kiwi.")
	write("Falsches \xdcben von Xylophonmusik qu\xe4lt jeden gr\xf6\xdferen Zwerg.")
	write("Heiz\xf6lr\xfccksto\xdfabd\xe4mpfung")
	write("For\xe5rsj\xe6vnd\xf8gn / Efter\xe5rsj\xe6vnd\xf8gn")
	fileStr := example.Filename("Fpdf_CellFormat_3_codepageescape")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
}
