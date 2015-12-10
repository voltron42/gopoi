package simple

import (
	"encoding/xml"
)

type Path struct {
	XMLName    xml.Name     `xml:"path" json:"-"`
	Close      bool         `xml:"close,attr"`
	Style      PathPaint    `xml:"style,attr"`
	Start      Coordinate   `xml:"start>point"`
	Items      PathItemList `xml:"path-items"`
	ShapeStyle *ShapeStyle  `xml:"style"`
}

func (p Path) Draw(ctx *context) error {
	if p.ShapeStyle != nil {
		p.ShapeStyle.set(ctx)
	}
	ctx.pdf.MoveTo(p.Start.X, p.Start.Y)
	err := ctx.pdf.Error()
	if err != nil {
		return err
	}
	for _, item := range p.Items {
		err = item.DrawPath(ctx)
		if err != nil {
			return err
		}
	}
	if p.Close {
		ctx.pdf.ClosePath()
		err = ctx.pdf.Error()
		if err != nil {
			return err
		}
	}
	ctx.pdf.DrawPath(p.Style.String())
	return ctx.pdf.Error()
}

type PathItemList []PathItem

type PathItem interface {
	DrawPath(ctx *context) error
}

type ArcTo struct {
	XMLName   xml.Name   `xml:"arc-to" json:"-"`
	Next      Coordinate `xml:"next>point"`
	Radius    Coordinate `xml:"radius>point"`
	DegRotate float64    `xml:"deg-rotate,attr"`
	DegStart  float64    `xml:"deg-start,attr"`
	DegEnd    float64    `xml:"deg-end,attr"`
}

func (a ArcTo) DrawPath(ctx *context) error {
	ctx.pdf.ArcTo(a.Next.X, a.Next.Y, a.Radius.X, a.Radius.Y, a.DegRotate, a.DegStart, a.DegEnd)
	return ctx.pdf.Error()
}

type LineTo struct {
	XMLName xml.Name   `xml:"line-to" json:"-"`
	Next    Coordinate `xml:"point"`
}

func (a LineTo) DrawPath(ctx *context) error {
	ctx.pdf.LineTo(a.Next.X, a.Next.Y)
	return ctx.pdf.Error()
}

type CurveTo struct {
	XMLName xml.Name   `xml:"curve-to" json:"-"`
	Center  Coordinate `xml:"center>point"`
	Next    Coordinate `xml:"next>point"`
}

func (a CurveTo) DrawPath(ctx *context) error {
	ctx.pdf.CurveTo(a.Center.X, a.Center.Y, a.Next.X, a.Next.Y)
	return ctx.pdf.Error()
}

type BezierCubicCurveTo struct {
	XMLName xml.Name   `xml:"bezier-cubic-curve-to" json:"-"`
	Center1 Coordinate `xml:"center1>point"`
	Center2 Coordinate `xml:"center2>point"`
	Next    Coordinate `xml:"next>point"`
}

func (a BezierCubicCurveTo) DrawPath(ctx *context) error {
	ctx.pdf.CurveBezierCubicTo(a.Center1.X, a.Center1.Y, a.Center2.X, a.Center2.Y, a.Next.X, a.Next.Y)
	return ctx.pdf.Error()
}
