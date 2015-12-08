package simple

import (
	"encoding/xml"
	"github.com/jung-kurt/gofpdf"
	"strconv"
	"strings"
)

type SizeType struct {
	Width, Height float64
}

func (s *SizeType) pdfSizeType() gofpdf.SizeType {
	return gofpdf.SizeType{s.Width, s.Height}
}

type RGB struct {
	R int `xml:"r,attr"`
	G int `xml:"g,attr"`
	B int `xml:"b,attr"`
}

type Coordinate struct {
	X float64 `xml:"x,attr"`
	Y float64 `xml:"y,attr"`
}

func (c Coordinate) toPointType() gofpdf.PointType {
	return gofpdf.PointType{c.X, c.Y}
}

type Frame struct {
	X float64 `xml:"x,attr"`
	Y float64 `xml:"y,attr"`
	W float64 `xml:"w,attr"`
	H float64 `xml:"h,attr"`
}

type SetFontStyle struct {
	Bold      bool
	Italic    bool
	Underline bool
}

func (s SetFontStyle) String() string {
	out := ""
	if s.Bold {
		out += "B"
	}
	if s.Italic {
		out += "I"
	}
	if s.Underline {
		out += "U"
	}
	return out
}

func (s *SetFontStyle) UnmarshalXMLAttr(attr xml.Attr) error {
	value := strings.ToLower(attr.Value)
	s.Bold = strings.Contains(value, "b")
	s.Italic = strings.Contains(value, "i")
	s.Underline = strings.Contains(value, "u")
	return nil
}

type DrawStyle struct {
	Fill bool
	Draw bool
}

func (s DrawStyle) String() string {
	out := ""
	if s.Fill {
		out += "F"
	}
	if s.Draw {
		out += "D"
	}
	return out
}

func (s *DrawStyle) UnmarshalXMLAttr(attr xml.Attr) error {
	value := strings.ToLower(attr.Value)
	s.Fill = strings.Contains(value, "f")
	s.Draw = strings.Contains(value, "d")
	return nil
}

type DashPattern struct {
	Array FloatList `xml:"dash-array,attr"`
	Phase float64   `xml:"dash-phase,attr"`
}

type FloatList []float64

func (f *FloatList) UnmarshalXmlAttr(attr xml.Attr) error {
	list := strings.Split(attr.Value, " ")
	for _, item := range list {
		num, err := strconv.ParseFloat(item, 64)
		if err != nil {
			return err
		}
		*f = append(*f, num)
	}
	return nil
}

type LineStyle struct {
	Width    *float64  `xml:"width,attr"`
	CapStyle *CapStyle `xml:"cap-style,attr"`
}

func (r *LineStyle) set(ctx *context) {
	if r.Width != nil {
		ctx.pdf.SetLineWidth(*r.Width)
	}
	if r.CapStyle != nil {
		ctx.pdf.SetLineCapStyle((*r.CapStyle).String())
	}
}

type ShapeStyle struct {
	JoinStyle *JoinStyle   `xml:"join-style,attr"`
	FillColor *RGB         `xml:"fill>rgb"`
	DrawColor *RGB         `xml:"draw>rgb"`
	Pattern   *DashPattern `xml:"dash-pattern"`
	LineStyle *LineStyle   `xml:"line-style"`
}

func (r *ShapeStyle) set(ctx *context) {
	if r.FillColor != nil {
		ctx.pdf.SetFillColor(r.FillColor.R, r.FillColor.G, r.FillColor.B)
	}
	if r.DrawColor != nil {
		ctx.pdf.SetDrawColor(r.DrawColor.R, r.DrawColor.G, r.DrawColor.B)
	}
	if r.Pattern != nil {
		ctx.pdf.SetDashPattern(r.Pattern.Array, r.Pattern.Phase)
	}
	if r.JoinStyle != nil {
		ctx.pdf.SetLineJoinStyle((*r.JoinStyle).String())
	}
	if r.LineStyle != nil {
		r.LineStyle.set(ctx)
	}
}

type Font struct {
	Family string       `xml:"family,attr"`
	Style  SetFontStyle `xml:"style,attr"`
	Size   float64      `xml:"size,attr"`
}

type TextStyle struct {
	Font  *Font `xml:"font"`
	Color *RGB  `xml:"rgb"`
}

func (t TextStyle) set(ctx *context) {
	if t.Font != nil {
		ctx.pdf.SetFont(t.Font.Family, t.Font.Style.String(), t.Font.Size)
	}
	if t.Color != nil {
		ctx.pdf.SetTextColor(t.Color.R, t.Color.G, t.Color.B)
	}
}

type Border struct {
	Top    bool `xml:"top,attr"`
	Left   bool `xml:"left,attr"`
	Right  bool `xml:"right,attr"`
	Bottom bool `xml:"bottom,attr"`
}

func (b Border) String() string {
	out := ""
	if b.Top {
		out += "T"
	}
	if b.Left {
		out += "L"
	}
	if b.Right {
		out += "R"
	}
	if b.Bottom {
		out += "B"
	}
	return out
}

type Alignment struct {
	Horiz HorizAlign `xml:"horiz,attr"`
	Vert  VertAlign  `xml:"vert,attr"`
}

func (a Alignment) String() string {
	return "LCR"[a.Horiz:1] + "TMB"[a.Vert:1]
}
