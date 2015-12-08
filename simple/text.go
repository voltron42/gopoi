package simple

import (
	"encoding/xml"
)

type Text struct {
	XMLName    xml.Name   `xml:"text"`
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

type Cell struct {
	XMLName    xml.Name   `xml:"cell"`
	TextString string     `xml:"string,attr"`
	CellMargin *float64   `xml:"margin,attr"`
	Point      Coordinate `xml:"point"`
	TextStyle  *TextStyle `xml:"text-style"`
}

func (t Cell) Draw(ctx *context) error {
	if t.TextStyle != nil {
		t.TextStyle.set(ctx)
	}
	if t.CellMargin != nil {
		ctx.pdf.SetCellMargin(*t.CellMargin)
	}
	ctx.pdf.Cell(t.Point.X, t.Point.Y, t.TextString)
	return ctx.pdf.Error()
}

type CellFormat struct {
	XMLName      xml.Name     `xml:"cell-format"`
	TextString   string       `xml:"string,attr"`
	CellMargin   *float64     `xml:"margin,attr"`
	NextPosition NextPosition `xml:"next-pos,attr"`
	Fill         bool         `xml:"fill,attr"`
	TextStyle    *TextStyle   `xml:"text-style"`
	Point        Coordinate   `xml:"point"`
	Border       Border       `xml:"border"`
	Align        Alignment    `xml:"alignment"`
}

func (t CellFormat) Draw(ctx *context) error {
	if t.TextStyle != nil {
		t.TextStyle.set(ctx)
	}
	if t.CellMargin != nil {
		ctx.pdf.SetCellMargin(*t.CellMargin)
	}
	ctx.pdf.CellFormat(t.Point.X, t.Point.Y, t.TextString, t.Border.String(), int(t.NextPosition), t.Align.String(), t.Fill, 0, "")
	return ctx.pdf.Error()
}

/*
func (f *Fpdf) Ln(h float64)
func (f *Fpdf) Bookmark(txtStr string, level int, y float64)
func (f *Fpdf) CellFormat(w, h float64, txtStr string, borderStr string, ln int, alignStr string, fill bool, link int, linkStr string)
func (f *Fpdf) Cellf(w, h float64, fmtStr string, args ...interface{})
func (f *Fpdf) MultiCell(w, h float64, txtStr, borderStr, alignStr string, fill bool)
func (f *Fpdf) RawWriteBuf(buf *bytes.Buffer)
func (f *Fpdf) RawWriteStr(str string)
func (f *Fpdf) SetCellMargin(margin float64)
func (f *Fpdf) Write(h float64, txtStr string)
func (f *Fpdf) WriteAligned(width, lineHeight float64, textStr, alignStr string)
func (f *Fpdf) WriteLinkID(h float64, displayStr string, linkID int)
func (f *Fpdf) WriteLinkString(h float64, displayStr, targetStr string)
func (f *Fpdf) Writef(h float64, fmtStr string, args ...interface{})
*/
