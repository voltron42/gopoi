package simple

import (
	"io"
)

type AddFont struct {
	Family   string
	Style    FontStyle
	FileName string
}

func (a AddFont) Draw(ctx *context) error {
	ctx.pdf.AddFont(a.Family, a.Style.String(), a.FileName)
	return ctx.pdf.Error()
}

type AddFontFromReader struct {
	Family     string
	Style      FontStyle
	FileReader io.Reader
}

func (a AddFontFromReader) Draw(ctx *context) error {
	ctx.pdf.AddFontFromReader(a.Family, a.Style.String(), a.FileReader)
	return ctx.pdf.Error()
}

type AddPage struct{}

func (a AddPage) Draw(ctx *context) error {
	ctx.pdf.AddPage()
	return ctx.pdf.Error()
}

type AddPageFormat struct {
	Orientation *Orientation
	SizeType    *SizeType
}

func (a AddPageFormat) Draw(ctx *context) error {
	ctx.pdf.AddPageFormat(a.Orientation.String(), a.SizeType.pdfSizeType())
	return ctx.pdf.Error()
}

type Image struct {
	FileName    string
	X, Y, W, H  float64
	Flow        bool
	ImageFormat *ImageFormat
	Link        *LinkId
	LinkString  *LinkString
}

func (i Image) Draw(ctx *context) error {
	ctx.pdf.Image(i.FileName, i.X, i.Y, i.W, i.H, i.Flow, i.ImageFormat.String(), int(*i.Link), string(*i.LinkString))
	return ctx.pdf.Error()
}

/*
func (f *Fpdf) SVGBasicWrite(sb *SVGBasicType, scale float64)
func (f *Fpdf) SetFontUnitSize(size float64)
Afunc (f *Fpdf) SetX(x float64)
func (f *Fpdf) SetXY(x, y float64)
func (f *Fpdf) SetY(y float64)
*/
