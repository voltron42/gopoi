package main

import (
	"../"
	"github.com/jung-kurt/gofpdf"
	"io"
)

type recType struct {
	align, txt string
}

func main() {
	recList := []recType{
		recType{"TL", "top left"},
		recType{"TC", "top center"},
		recType{"TR", "top right"},
		recType{"LM", "middle left"},
		recType{"CM", "middle center"},
		recType{"RM", "middle right"},
		recType{"BL", "bottom left"},
		recType{"BC", "bottom center"},
		recType{"BR", "bottom right"},
	}
	recListBaseline := []recType{
		recType{"AL", "baseline left"},
		recType{"AC", "baseline center"},
		recType{"AR", "baseline right"},
	}
	var formatRect = func(pdf *gofpdf.Fpdf, recList []recType) {
		linkStr := ""
		for pageJ := 0; pageJ < 2; pageJ++ {
			pdf.AddPage()
			pdf.SetMargins(10, 10, 10)
			pdf.SetAutoPageBreak(false, 0)
			borderStr := "1"
			for _, rec := range recList {
				pdf.SetXY(20, 20)
				pdf.CellFormat(170, 257, rec.txt, borderStr, 0, rec.align, false, 0, linkStr)
				borderStr = ""
			}
			linkStr = "https://github.com/jung-kurt/gofpdf"
		}
	}
	pdf := gofpdf.New("P", "mm", "A4", "") // A4 210.0 x 297.0
	pdf.SetFont("Helvetica", "", 16)
	formatRect(pdf, recList)
	formatRect(pdf, recListBaseline)
	var fr fontResourceType
	pdf.SetFontLoader(fr)
	pdf.AddFont("Calligrapher", "", "calligra.json")
	pdf.SetFont("Calligrapher", "", 16)
	formatRect(pdf, recListBaseline)
	fileStr := example.Filename("Fpdf_CellFormat_2_align")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
}

type fontResourceType struct {
}

func (f fontResourceType) Open(name string) (io.Reader, error) {
	// TODO --
	return nil, nil
}
