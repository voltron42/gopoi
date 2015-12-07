package main

import (
	"../"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	const offset = 75.0
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()
	var draw = func(cap, join string, x0, y0, x1, y1 float64) {
		// transform begin & end needed to isolate caps and joins
		pdf.SetLineCapStyle(cap)
		pdf.SetLineJoinStyle(join)

		// Draw thick line
		pdf.SetDrawColor(0x33, 0x33, 0x33)
		pdf.SetLineWidth(30.0)
		pdf.MoveTo(x0, y0)
		pdf.LineTo((x0+x1)/2+offset, (y0+y1)/2)
		pdf.LineTo(x1, y1)
		pdf.DrawPath("D")

		// Draw thin helping line
		pdf.SetDrawColor(0xFF, 0x33, 0x33)
		pdf.SetLineWidth(2.56)
		pdf.MoveTo(x0, y0)
		pdf.LineTo((x0+x1)/2+offset, (y0+y1)/2)
		pdf.LineTo(x1, y1)
		pdf.DrawPath("D")

	}
	x := 35.0
	caps := []string{"butt", "square", "round"}
	joins := []string{"bevel", "miter", "round"}
	for i := range caps {
		draw(caps[i], joins[i], x, 50, x, 160)
		x += offset
	}
	fileStr := example.Filename("Fpdf_SetLineJoinStyle_caps")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
}
