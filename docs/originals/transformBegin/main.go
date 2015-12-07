package main

import (
	"../"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	const (
		light = 200
		dark  = 0
	)
	var refX, refY float64
	var refStr string
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	color := func(val int) {
		pdf.SetDrawColor(val, val, val)
		pdf.SetTextColor(val, val, val)
	}
	reference := func(str string, x, y float64, val int) {
		color(val)
		pdf.Rect(x, y, 40, 10, "D")
		pdf.Text(x, y-1, str)
	}
	refDraw := func(str string, x, y float64) {
		refStr = str
		refX = x
		refY = y
		reference(str, x, y, light)
	}
	refDupe := func() {
		reference(refStr, refX, refY, dark)
	}

	titleStr := "Transformations"
	titlePt := 36.0
	titleHt := pdf.PointConvert(titlePt)
	pdf.SetFont("Helvetica", "", titlePt)
	titleWd := pdf.GetStringWidth(titleStr)
	titleX := (210 - titleWd) / 2
	pdf.Text(titleX, 10+titleHt, titleStr)
	pdf.TransformBegin()
	pdf.TransformMirrorVertical(10 + titleHt + 0.5)
	pdf.ClipText(titleX, 10+titleHt, titleStr, false)
	// Remember that the transform will mirror the gradient box too
	pdf.LinearGradient(titleX, 10, titleWd, titleHt+4, 120, 120, 120,
		255, 255, 255, 0, 0, 0, 0.6)
	pdf.ClipEnd()
	pdf.TransformEnd()

	pdf.SetFont("Helvetica", "", 12)

	// Scale by 150% centered by lower left corner of the rectangle
	refDraw("Scale", 50, 60)
	pdf.TransformBegin()
	pdf.TransformScaleXY(150, 50, 70)
	refDupe()
	pdf.TransformEnd()

	// Translate 7 to the right, 5 to the bottom
	refDraw("Translate", 125, 60)
	pdf.TransformBegin()
	pdf.TransformTranslate(7, 5)
	refDupe()
	pdf.TransformEnd()

	// Rotate 20 degrees counter-clockwise centered by the lower left corner of
	// the rectangle
	refDraw("Rotate", 50, 110)
	pdf.TransformBegin()
	pdf.TransformRotate(20, 50, 120)
	refDupe()
	pdf.TransformEnd()

	// Skew 30 degrees along the x-axis centered by the lower left corner of the
	// rectangle
	refDraw("Skew", 125, 110)
	pdf.TransformBegin()
	pdf.TransformSkewX(30, 125, 110)
	refDupe()
	pdf.TransformEnd()

	// Mirror horizontally with axis of reflection at left side of the rectangle
	refDraw("Mirror horizontal", 50, 160)
	pdf.TransformBegin()
	pdf.TransformMirrorHorizontal(50)
	refDupe()
	pdf.TransformEnd()

	// Mirror vertically with axis of reflection at bottom side of the rectangle
	refDraw("Mirror vertical", 125, 160)
	pdf.TransformBegin()
	pdf.TransformMirrorVertical(170)
	refDupe()
	pdf.TransformEnd()

	// Reflect against a point at the lower left point of rectangle
	refDraw("Mirror point", 50, 210)
	pdf.TransformBegin()
	pdf.TransformMirrorPoint(50, 220)
	refDupe()
	pdf.TransformEnd()

	// Mirror against a straight line described by a point and an angle
	angle := -20.0
	px := 120.0
	py := 220.0
	refDraw("Mirror line", 125, 210)
	pdf.TransformBegin()
	pdf.TransformRotate(angle, px, py)
	pdf.Line(px-1, py-1, px+1, py+1)
	pdf.Line(px-1, py+1, px+1, py-1)
	pdf.Line(px-5, py, px+60, py)
	pdf.TransformEnd()
	pdf.TransformBegin()
	pdf.TransformMirrorLine(angle, px, py)
	refDupe()
	pdf.TransformEnd()

	fileStr := example.Filename("Fpdf_TransformBegin")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
}
