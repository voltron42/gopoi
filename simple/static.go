package simple

import (
	"github.com/jung-kurt/gofpdf"
)

func PointConvert(pt float64, unit Unit) float64 {
	return gofpdf.New("", unit.String(), "", "").PointConvert(pt)
}
