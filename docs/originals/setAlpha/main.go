package main

import (
	"../"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	const (
		gapX  = 10.0
		gapY  = 9.0
		rectW = 40.0
		rectH = 58.0
		pageW = 210
		pageH = 297
	)
	modeList := []string{"Normal", "Multiply", "Screen", "Overlay",
		"Darken", "Lighten", "ColorDodge", "ColorBurn", "HardLight", "SoftLight",
		"Difference", "Exclusion", "Hue", "Saturation", "Color", "Luminosity"}
	pdf := gofpdf.New("", "", "", "")
	pdf.SetLineWidth(2)
	pdf.SetAutoPageBreak(false, 0)
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", 18)
	pdf.SetXY(0, gapY)
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(pageW, gapY, "Alpha Blending Modes", "", 0, "C", false, 0, "")
	j := 0
	y := 3 * gapY
	for col := 0; col < 4; col++ {
		x := gapX
		for row := 0; row < 4; row++ {
			pdf.Rect(x, y, rectW, rectH, "D")
			pdf.SetFont("Helvetica", "B", 12)
			pdf.SetFillColor(0, 0, 0)
			pdf.SetTextColor(250, 250, 230)
			pdf.SetXY(x, y+rectH-4)
			pdf.CellFormat(rectW, 5, modeList[j], "", 0, "C", true, 0, "")
			pdf.SetFont("Helvetica", "I", 150)
			pdf.SetTextColor(80, 80, 120)
			pdf.SetXY(x, y+2)
			pdf.CellFormat(rectW, rectH, "A", "", 0, "C", false, 0, "")
			pdf.SetAlpha(0.5, modeList[j])
			pdf.Image(example.ImageFile("golang-gopher.png"),
				x-gapX, y, rectW+2*gapX, 0, false, "", 0, "")
			pdf.SetAlpha(1.0, "Normal")
			x += rectW + gapX
			j++
		}
		y += rectH + gapY
	}
	fileStr := example.Filename("Fpdf_SetAlpha_transparency")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
}
