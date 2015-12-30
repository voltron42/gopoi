package simple

import (
	"encoding/xml"
)

type LinearGradient struct {
	XMLName xml.Name   `xml:"linear-gradient" json:"-"`
	Frame   Frame      `xml:"frame"`
	Color1  RGB        `xml:"color1"`
	Color2  RGB        `xml:"color2"`
	Point1  Coordinate `xml:"point1"`
	Point2  Coordinate `xml:"point2"`
}

func (l LinearGradient) Draw(ctx *context) error {
	ctx.pdf.LinearGradient(
		l.Frame.X,
		l.Frame.Y,
		l.Frame.W,
		l.Frame.H,
		l.Color1.R,
		l.Color1.G,
		l.Color1.B,
		l.Color2.R,
		l.Color2.G,
		l.Color2.B,
		l.Point1.X,
		l.Point1.Y,
		l.Point2.X,
		l.Point2.Y,
	)
	return ctx.pdf.Error()
}

type RadialGradient struct {
	XMLName xml.Name   `xml:"radial-gradient" json:"-"`
	Radius  float64    `xml:"radius,attr"`
	Frame   Frame      `xml:"frame"`
	Color1  RGB        `xml:"color1"`
	Color2  RGB        `xml:"color2"`
	Point1  Coordinate `xml:"point1"`
	Point2  Coordinate `xml:"point2"`
}

func (l RadialGradient) Draw(ctx *context) error {
	ctx.pdf.RadialGradient(
		l.Frame.X,
		l.Frame.Y,
		l.Frame.W,
		l.Frame.H,
		l.Color1.R,
		l.Color1.G,
		l.Color1.B,
		l.Color2.R,
		l.Color2.G,
		l.Color2.B,
		l.Point1.X,
		l.Point1.Y,
		l.Point2.X,
		l.Point2.Y,
		l.Radius,
	)
	return ctx.pdf.Error()
}

type SetDrawColor struct {
	XMLName xml.Name `xml:"set-draw-color" json:"-"`
	Color   RGB      `xml:"rgb"`
}

func (s SetDrawColor) Draw(ctx *context) error {
	ctx.pdf.SetDrawColor(
		s.Color.R,
		s.Color.G,
		s.Color.B,
	)
	return ctx.pdf.Error()
}

type SetFillColor struct {
	XMLName xml.Name `xml:"set-fill-color" json:"-"`
	Color   RGB      `xml:"rgb"`
}

func (s SetFillColor) Draw(ctx *context) error {
	ctx.pdf.SetFillColor(
		s.Color.R,
		s.Color.G,
		s.Color.B,
	)
	return ctx.pdf.Error()
}

type SetTextColor struct {
	XMLName xml.Name `xml:"set-text-color" json:"-"`
	Color   RGB      `xml:"rgb"`
}

func (s SetTextColor) Draw(ctx *context) error {
	ctx.pdf.SetTextColor(
		s.Color.R,
		s.Color.G,
		s.Color.B,
	)
	return ctx.pdf.Error()
}

type SetMargins struct {
	XMLName xml.Name `xml:"set-margins" json:"-"`
	Left    float64  `xml:"left,attr"`
	Top     float64  `xml:"top,attr"`
	Right   float64  `xml:"right,attr"`
}

func (s SetMargins) Draw(ctx *context) error {
	ctx.pdf.SetMargins(
		s.Left,
		s.Top,
		s.Right,
	)
	return ctx.pdf.Error()
}

type SetRightMargin struct {
	XMLName xml.Name `xml:"set-right-margin" json:"-"`
	Margin  float64  `xml:"margin,attr"`
}

func (s SetRightMargin) Draw(ctx *context) error {
	ctx.pdf.SetRightMargin(s.Margin)
	return ctx.pdf.Error()
}

type SetLeftMargin struct {
	XMLName xml.Name `xml:"set-left-margin" json:"-"`
	Margin  float64  `xml:"margin,attr"`
}

func (s SetLeftMargin) Draw(ctx *context) error {
	ctx.pdf.SetLeftMargin(s.Margin)
	return ctx.pdf.Error()
}

type SetTopMargin struct {
	XMLName xml.Name `xml:"set-top-margin" json:"-"`
	Margin  float64  `xml:"margin,attr"`
}

func (s SetTopMargin) Draw(ctx *context) error {
	ctx.pdf.SetTopMargin(s.Margin)
	return ctx.pdf.Error()
}

type SetFontSize struct {
	XMLName xml.Name `xml:"set-font-size" json:"-"`
	Size    float64  `xml:"size,attr"`
}

func (s SetFontSize) Draw(ctx *context) error {
	ctx.pdf.SetFontSize(s.Size)
	return ctx.pdf.Error()
}

type SetLineWidth struct {
	XMLName xml.Name `xml:"set-line-width" json:"-"`
	Width   float64  `xml:"width,attr"`
}

func (s SetLineWidth) Draw(ctx *context) error {
	ctx.pdf.SetLineWidth(s.Width)
	return ctx.pdf.Error()
}

type SetAlpha struct {
	XMLName   xml.Name  `xml:"set-alpha json:"-""`
	Alpha     float64   `xml:"alpha,attr"`
	BlendMode BlendMode `xml:"blend-mode,attr"`
}

func (s SetAlpha) Draw(ctx *context) error {
	ctx.pdf.SetAlpha(s.Alpha, s.BlendMode.String())
	return ctx.pdf.Error()
}

type SetFont struct {
	XMLName xml.Name `xml:"set-font" json:"-"`
	Font    Font     `xml:"font"`
}

func (s SetFont) Draw(ctx *context) error {
	ctx.pdf.SetFont(s.Font.Family, s.Font.Style.String(), s.Font.Size)
	return ctx.pdf.Error()
}

type SetLineCapStyle struct {
	XMLName  xml.Name `xml:"set-line-cap-style" json:"-"`
	CapStyle CapStyle `xml:"cap-style,attr"`
}

func (s SetLineCapStyle) Draw(ctx *context) error {
	ctx.pdf.SetLineCapStyle(s.CapStyle.String())
	return ctx.pdf.Error()
}

type SetLineJoinStyle struct {
	XMLName   xml.Name  `xml:"set-line-join-style" json:"-"`
	JoinStyle JoinStyle `xml:"join-style,attr"`
}

func (s SetLineJoinStyle) Draw(ctx *context) error {
	ctx.pdf.SetLineJoinStyle(s.JoinStyle.String())
	return ctx.pdf.Error()
}

type SetDashPattern struct {
	XMLName     xml.Name    `xml:"set-dash-pattern" json:"-"`
	DashPattern DashPattern `xml:"dash-pattern"`
}

func (s SetDashPattern) Draw(ctx *context) error {
	ctx.pdf.SetDashPattern(s.DashPattern.Array, s.DashPattern.Phase)
	return ctx.pdf.Error()
}
