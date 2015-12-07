package main

import (
	"../"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	var err error
	fileStr := example.Filename("Fpdf_SetKeywords")
	err = gofpdf.MakeFont(example.FontFile("CalligrapherRegular.pfb"),
		example.FontFile("cp1252.map"), example.FontDir(), nil, true)
	if err == nil {
		err = gofpdf.MakeFont(example.FontFile("calligra.ttf"),
			example.FontFile("cp1252.map"), example.FontDir(), nil, true)
		if err == nil {
			pdf := gofpdf.New("", "", "", "")
			pdf.SetFontLocation(example.FontDir())
			pdf.SetTitle("世界", true)
			pdf.SetAuthor("世界", true)
			pdf.SetSubject("世界", true)
			pdf.SetCreator("世界", true)
			pdf.SetKeywords("世界", true)
			pdf.AddFont("Calligrapher", "", "CalligrapherRegular.json")
			pdf.AddPage()
			pdf.SetFont("Calligrapher", "", 16)
			pdf.Writef(5, "\x95 %s \x95", pdf)
			err = pdf.OutputFileAndClose(fileStr)
		}
	}
	example.Summary(err, fileStr)
}
