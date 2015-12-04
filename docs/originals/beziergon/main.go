package main

import (
	"../"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	const (
		margin      = 10
		wd          = 210
		unit        = (wd - 2*margin) / 6
		ht          = 297
		fontSize    = 15
		msgStr      = `Demonstration of Beziergon function`
		coefficient = 0.6
		delta       = coefficient * unit
		ln          = fontSize * 25.4 / 72
		offsetX     = (wd - 4*unit) / 2.0
		offsetY     = offsetX + 2*ln
	)

	srcList := []gofpdf.PointType{
		{X: 0, Y: 0},
		{X: 1, Y: 0},
		{X: 1, Y: 1},
		{X: 2, Y: 1},
		{X: 2, Y: 2},
		{X: 3, Y: 2},
		{X: 3, Y: 3},
		{X: 4, Y: 3},
		{X: 4, Y: 4},
		{X: 1, Y: 4},
		{X: 1, Y: 3},
		{X: 0, Y: 3},
	}

	ctrlList := []gofpdf.PointType{
		{X: 1, Y: -1},
		{X: 1, Y: 1},
		{X: 1, Y: 1},
		{X: 1, Y: 1},
		{X: 1, Y: 1},
		{X: 1, Y: 1},
		{X: 1, Y: 1},
		{X: 1, Y: 1},
		{X: -1, Y: 1},
		{X: -1, Y: -1},
		{X: -1, Y: -1},
		{X: -1, Y: -1},
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", fontSize)
	for j, src := range srcList {
		srcList[j].X = offsetX + src.X*unit
		srcList[j].Y = offsetY + src.Y*unit
	}
	for j, ctrl := range ctrlList {
		ctrlList[j].X = ctrl.X * delta
		ctrlList[j].Y = ctrl.Y * delta
	}
	jPrev := len(srcList) - 1
	srcPrev := srcList[jPrev]
	curveList := []gofpdf.PointType{srcPrev} // point [, control 0, control 1, point]*
	control := func(x, y float64) {
		curveList = append(curveList, gofpdf.PointType{X: x, Y: y})
	}
	for j, src := range srcList {
		ctrl := ctrlList[jPrev]
		control(srcPrev.X+ctrl.X, srcPrev.Y+ctrl.Y) // Control 0
		ctrl = ctrlList[j]
		control(src.X-ctrl.X, src.Y-ctrl.Y) // Control 1
		curveList = append(curveList, src)  // Destination
		jPrev = j
		srcPrev = src
	}
	pdf.MultiCell(wd-margin-margin, ln, msgStr, "", "C", false)
	pdf.SetDashPattern([]float64{0.8, 0.8}, 0)
	pdf.SetDrawColor(160, 160, 160)
	pdf.Polygon(srcList, "D")
	pdf.SetDashPattern([]float64{}, 0)
	pdf.SetDrawColor(64, 64, 128)
	pdf.SetLineWidth(pdf.GetLineWidth() * 3)
	pdf.Beziergon(curveList, "D")
	fileStr := example.Filename("Fpdf_Beziergon")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
}
