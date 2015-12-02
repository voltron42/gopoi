package simple

import (
	"github.com/jung-kurt/gofpdf"
)

type SizeType struct {
	Width, Height float64
}

func (s *SizeType) pdfSizeType() gofpdf.SizeType {
	return gofpdf.SizeType{s.Width, s.Height}
}
