package simple

type AddFont struct {
	Family   string    `xml:"family,attr"`
	Style    FontStyle `xml:"style,attr"`
	FileName string    `xml:"file-name,attr"`
}

func (a AddFont) Draw(ctx *context) error {
	ctx.pdf.AddFont(a.Family, a.Style.String(), a.FileName)
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

/*
func (f *Fpdf) SVGBasicWrite(sb *SVGBasicType, scale float64)
func (f *Fpdf) SetFontUnitSize(size float64)
Afunc (f *Fpdf) SetX(x float64)
func (f *Fpdf) SetXY(x, y float64)
func (f *Fpdf) SetY(y float64)
*/
