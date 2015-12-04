package main

import (
	"../"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("", "", "", "")
	y := 10.0
	pdf.AddPage()

	pdf.SetFont("Helvetica", "", 24)
	pdf.SetXY(0, y)
	pdf.ClipText(10, y+12, "Clipping examples", false)
	pdf.RadialGradient(10, y, 100, 20, 128, 128, 160, 32, 32, 48,
		0.25, 0.5, 0.25, 0.5, 0.2)
	pdf.ClipEnd()

	y += 12
	pdf.SetFont("Helvetica", "B", 120)
	pdf.SetDrawColor(64, 80, 80)
	pdf.SetLineWidth(.5)
	pdf.ClipText(10, y+40, pdf.String(), true)
	pdf.RadialGradient(10, y, 200, 50, 220, 220, 250, 80, 80, 220,
		0.25, 0.5, 0.25, 0.5, 1)
	pdf.ClipEnd()

	y += 55
	pdf.ClipRect(10, y, 105, 20, true)
	pdf.SetFillColor(255, 255, 255)
	pdf.Rect(10, y, 105, 20, "F")
	pdf.ClipCircle(40, y+10, 15, false)
	pdf.RadialGradient(25, y, 30, 30, 220, 250, 220, 40, 60, 40, 0.3,
		0.85, 0.3, 0.85, 0.5)
	pdf.ClipEnd()
	pdf.ClipEllipse(80, y+10, 20, 15, false)
	pdf.RadialGradient(60, y, 40, 30, 250, 220, 220, 60, 40, 40, 0.3,
		0.85, 0.3, 0.85, 0.5)
	pdf.ClipEnd()
	pdf.ClipEnd()

	y += 28
	pdf.ClipEllipse(26, y+10, 16, 10, true)
	pdf.Image(example.ImageFile("logo.jpg"), 10, y, 32, 0, false, "JPG", 0, "")
	pdf.ClipEnd()

	pdf.ClipCircle(60, y+10, 10, true)
	pdf.RadialGradient(50, y, 20, 20, 220, 220, 250, 40, 40, 60, 0.3,
		0.7, 0.3, 0.7, 0.5)
	pdf.ClipEnd()

	pdf.ClipPolygon([]gofpdf.PointType{{80, y + 20}, {90, y},
		{100, y + 20}}, true)
	pdf.LinearGradient(80, y, 20, 20, 250, 220, 250, 60, 40, 60, 0.5,
		1, 0.5, 0.5)
	pdf.ClipEnd()

	y += 30
	pdf.SetLineWidth(.1)
	pdf.SetDrawColor(180, 180, 180)
	pdf.ClipRoundedRect(10, y, 120, 20, 5, true)
	pdf.RadialGradient(10, y, 120, 20, 255, 255, 255, 240, 240, 220,
		0.25, 0.75, 0.25, 0.75, 0.5)
	pdf.SetXY(5, y-5)
	pdf.SetFont("Times", "", 12)
	pdf.MultiCell(130, 5, lorem(), "", "", false)
	pdf.ClipEnd()

	fileStr := example.Filename("Fpdf_ClipText")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
}

func lorem() string {
	// TODO --
	return ""
}
