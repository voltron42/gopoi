package simple

import (
	"encoding/xml"
)

type AddFont struct {
	XMLName  xml.Name  `xml:"add-font"`
	Family   string    `xml:"family,attr"`
	Style    FontStyle `xml:"style,attr"`
	FileName string    `xml:"file-name,attr"`
}

func (a AddFont) Draw(ctx *context) error {
	ctx.pdf.AddFont(a.Family, a.Style.String(), a.FileName)
	return ctx.pdf.Error()
}

type AddPage struct {
	XMLName xml.Name `xml:"add-page"`
}

func (a AddPage) Draw(ctx *context) error {
	ctx.pdf.AddPage()
	return ctx.pdf.Error()
}

type AddPageFormat struct {
	XMLName     xml.Name `xml:"add-page-format"`
	Orientation *Orientation
	SizeType    *SizeType
}

func (a AddPageFormat) Draw(ctx *context) error {
	ctx.pdf.AddPageFormat(a.Orientation.String(), a.SizeType.pdfSizeType())
	return ctx.pdf.Error()
}

type Image struct {
	XMLName     xml.Name     `xml:"image"`
	FileName    string       `xml:"filename,attr"`
	Frame       Frame        `xml:"frame"`
	Flow        bool         `xml:"flow,attr"`
	ImageFormat *ImageFormat `xml:"format,attr"`
	Link        *LinkId      `xml:"link,attr"`
	LinkString  *LinkString  `xml:"link-string,attr"`
}

func (i Image) Draw(ctx *context) error {
	ctx.pdf.Image(i.FileName, i.Frame.X, i.Frame.Y, i.Frame.W, i.Frame.H, i.Flow, i.ImageFormat.String(), int(*i.Link), string(*i.LinkString))
	return ctx.pdf.Error()
}

type SetX struct {
	XMLName xml.Name `xml:"set-x"`
	X       float64  `xml:"x,attr"`
}

func (s SetX) Draw(ctx *context) error {
	ctx.pdf.SetX(s.X)
	return ctx.pdf.Error()
}

type SetXY struct {
	XMLName xml.Name `xml:"set-xy"`
	X       float64  `xml:"x,attr"`
	Y       float64  `xml:"y,attr"`
}

func (s SetXY) Draw(ctx *context) error {
	ctx.pdf.SetXY(s.X, s.Y)
	return ctx.pdf.Error()
}

type SetY struct {
	XMLName xml.Name `xml:"set-x"`
	Y       float64  `xml:"Y,attr"`
}

func (s SetY) Draw(ctx *context) error {
	ctx.pdf.SetY(s.Y)
	return ctx.pdf.Error()
}

/*
func (f *Fpdf) SVGBasicWrite(sb *SVGBasicType, scale float64)
func (f *Fpdf) SetFontUnitSize(size float64)
*/
