package simple

import (
	"github.com/jung-kurt/gofpdf"
)

func PointConvert(pt float64, unit Unit) float64 {
	return gofpdf.New("", unit.String(), "", "").PointConvert(pt)
}

func GetStringWidth(str string, f *Format, font Font) float64 {
	ctx := f.newContext()
	SetFont{Font: font}.Draw(ctx)
	return ctx.pdf.GetStringWidth(str)
}
