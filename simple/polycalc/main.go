package main

import (
	"../"
	"math"
)

func main() {
	const rowCount = 5
	const colCount = 4
	const ptSize = 36
	var x, y, radius, gap, advance float64
	var rgVal int
	var pts []simple.Coordinate
	vertices := func(count int) (res []simple.Coordinate) {
		var pt simple.Coordinate
		res = make([]simple.Coordinate, 0, count)
		mlt := 2.0 * math.Pi / float64(count)
		for j := 0; j < count; j++ {
			pt.Y, pt.X = math.Sincos(float64(j) * mlt)
			res = append(res, simple.Coordinate{
				X: x + radius*pt.X,
				Y: y + radius*pt.Y})
		}
		return
	}
	pdf := gofpdf.New("P", "mm", "A4", "") // A4 210.0 x 297.0
	pdf := simple.Document{
		Format: simple.Format{
			Orientation:   simple.Portrait,
			Unit:          simple.Millimeter,
			Size:          simple.A4,
			FontDirectory: "",
			OutputFile:    "polygon.pdf",
		},
		Items: ItemList{
			AddPage{},
		},
	}
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", ptSize)
	pdf.SetDrawColor(0, 80, 180)
	gap = 12.0
	pdf.SetY(gap)
	pdf.CellFormat(190.0, gap, "Equilateral polygons", "", 1, "C", false, 0, "")
	radius = (210.0 - float64(colCount+1)*gap) / (2.0 * float64(colCount))
	advance = gap + 2.0*radius
	y = 2*gap + pdf.PointConvert(ptSize) + radius
	rgVal = 230
	for row := 0; row < rowCount; row++ {
		pdf.SetFillColor(rgVal, rgVal, 0)
		rgVal -= 12
		x = gap + radius
		for col := 0; col < colCount; col++ {
			pts = vertices(row*colCount + col + 3)
			pdf.Polygon(pts, "FD")
			x += advance
		}
		y += advance
	}
	fileStr := example.Filename("Fpdf_Polygon")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
}
