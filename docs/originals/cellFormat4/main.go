package main

import (
	"../"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", example.FontDir()) // A4 210.0 x 297.0
	// See documentation for details on how to generate fonts
	pdf.AddFont("Helvetica-1251", "", "helvetica_1251.json")
	pdf.AddFont("Helvetica-1253", "", "helvetica_1253.json")
	fontSize := 16.0
	pdf.SetFont("Helvetica", "", fontSize)
	ht := pdf.PointConvert(fontSize)
	tr := pdf.UnicodeTranslatorFromDescriptor("") // "" defaults to "cp1252"
	write := func(str string) {
		pdf.CellFormat(190, ht, tr(str), "", 1, "C", false, 0, "")
		pdf.Ln(ht)
	}
	pdf.AddPage()
	str := `Gofpdf provides a translator that will convert any UTF-8 code point ` +
		`that is present in the specified code page.`
	pdf.MultiCell(190, ht, str, "", "L", false)
	pdf.Ln(2 * ht)
	write("Voix ambiguë d'un cœur qui au zéphyr préfère les jattes de kiwi.")
	write("Falsches Üben von Xylophonmusik quält jeden größeren Zwerg.")
	write("Heizölrückstoßabdämpfung")
	write("Forårsjævndøgn / Efterårsjævndøgn")

	pdf.SetFont("Helvetica-1251", "", fontSize) // Name matches one specified in AddFont()
	tr = pdf.UnicodeTranslatorFromDescriptor("cp1251")
	write("Съешь же ещё этих мягких французских булок, да выпей чаю.")

	pdf.SetFont("Helvetica-1253", "", fontSize)
	tr = pdf.UnicodeTranslatorFromDescriptor("cp1253")
	write("Θέλει αρετή και τόλμη η ελευθερία. (Ανδρέας Κάλβος)")

	fileStr := example.Filename("Fpdf_CellFormat_4_codepage")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
}
