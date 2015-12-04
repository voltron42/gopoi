package main

import (
	"../"
	"github.com/jung-kurt/gofpdf"
	"math"
	"strconv"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetDrawColor(0xff, 0x00, 0x00)
	pdf.SetFillColor(0x99, 0x99, 0x99)
	pdf.SetFont("Helvetica", "", 15)
	pdf.AddPage()
	pdf.SetAlpha(1, "Multiply")
	var (
		polygon = func(cx, cy, r, n, dir float64) {
			da := 2 * math.Pi / n
			pdf.MoveTo(cx+r, cy)
			pdf.Text(cx+r, cy, "0")
			i := 1
			for a := da; a < 2*math.Pi; a += da {
				x, y := cx+r*math.Cos(dir*a), cy+r*math.Sin(dir*a)
				pdf.LineTo(x, y)
				pdf.Text(x, y, strconv.Itoa(i))
				i++
			}
			pdf.ClosePath()
		}
		polygons = func(cx, cy, r, n, dir float64) {
			d := 1.0
			for rf := r; rf > 0; rf -= 10 {
				polygon(cx, cy, rf, n, d)
				d *= dir
			}
		}
		star = func(cx, cy, r, n float64) {
			da := 4 * math.Pi / n
			pdf.MoveTo(cx+r, cy)
			for a := da; a < 4*math.Pi+da; a += da {
				x, y := cx+r*math.Cos(a), cy+r*math.Sin(a)
				pdf.LineTo(x, y)
			}
			pdf.ClosePath()
		}
	)
	// triangle
	polygons(55, 45, 40, 3, 1)
	pdf.DrawPath("B")
	pdf.Text(15, 95, "B (same direction, non zero winding)")

	// square
	polygons(155, 45, 40, 4, 1)
	pdf.DrawPath("B*")
	pdf.Text(115, 95, "B* (same direction, even odd)")

	// pentagon
	polygons(55, 145, 40, 5, -1)
	pdf.DrawPath("B")
	pdf.Text(15, 195, "B (different direction, non zero winding)")

	// hexagon
	polygons(155, 145, 40, 6, -1)
	pdf.DrawPath("B*")
	pdf.Text(115, 195, "B* (different direction, even odd)")

	// star
	star(55, 245, 40, 5)
	pdf.DrawPath("B")
	pdf.Text(15, 290, "B (non zero winding)")

	// star
	star(155, 245, 40, 5)
	pdf.DrawPath("B*")
	pdf.Text(115, 290, "B* (even odd)")

	fileStr := example.Filename("Fpdf_DrawPath_fill")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
}
