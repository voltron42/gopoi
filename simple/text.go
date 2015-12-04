package simple

type Text struct {
	TextString string     `xml:"string,attr"`
	Point      Coordinate `xml:"point"`
	TextStyle  *TextStyle `xml:"text-style"`
}

func (t Text) Draw(ctx *context) error {
	if t.TextStyle != nil {
		t.TextStyle.set(ctx)
	}
	ctx.pdf.Text(t.Point.X, t.Point.Y, t.TextString)
	return ctx.pdf.Error()
}

/*
func (f *Fpdf) Ln(h float64)
func (f *Fpdf) Bookmark(txtStr string, level int, y float64)
func (f *Fpdf) Cell(w, h float64, txtStr string)
func (f *Fpdf) CellFormat(w, h float64, txtStr string, borderStr string, ln int, alignStr string, fill bool, link int, linkStr string)
func (f *Fpdf) Cellf(w, h float64, fmtStr string, args ...interface{})
func (f *Fpdf) MultiCell(w, h float64, txtStr, borderStr, alignStr string, fill bool)
func (f *Fpdf) RawWriteBuf(buf *bytes.Buffer)
func (f *Fpdf) RawWriteStr(str string)
func (f *Fpdf) SetCellMargin(margin float64)
func (f *Fpdf) Text(x, y float64, txtStr string)
func (f *Fpdf) Write(h float64, txtStr string)
func (f *Fpdf) WriteAligned(width, lineHeight float64, textStr, alignStr string)
func (f *Fpdf) WriteLinkID(h float64, displayStr string, linkID int)
func (f *Fpdf) WriteLinkString(h float64, displayStr, targetStr string)
func (f *Fpdf) Writef(h float64, fmtStr string, args ...interface{})
*/
