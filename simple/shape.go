package simple

import (
	"github.com/jung-kurt/gofpdf"
)

type Line struct {
	Start     Coordinate `xml:"start>point"`
	End       Coordinate `xml:"end>point"`
	DrawColor *RGB       `xml:"draw>color"`
	LineStyle *LineStyle `xml:"line-style"`
}

func (r Line) Draw(ctx *context) error {
	if r.DrawColor != nil {
		ctx.pdf.SetDrawColor(r.DrawColor.R, r.DrawColor.G, r.DrawColor.B)
	}
	if r.LineStyle != nil {
		r.LineStyle.set(ctx)
	}
	ctx.pdf.Line(r.Start.X, r.Start.Y, r.End.X, r.End.Y)
	return ctx.pdf.Error()
}

type Rect struct {
	DrawStyle  DrawStyle   `xml:"style,attr"`
	Frame      Frame       `xml:"frame"`
	ShapeStyle *ShapeStyle `xml:"style"`
}

func (r Rect) Draw(ctx *context) error {
	if r.ShapeStyle != nil {
		r.ShapeStyle.set(ctx)
	}
	ctx.pdf.Rect(r.Frame.X, r.Frame.Y, r.Frame.W, r.Frame.H, r.DrawStyle.String())
	return ctx.pdf.Error()
}

type Circle struct {
	Radius     float64     `xml:"radius,attr"`
	DrawStyle  DrawStyle   `xml:"style,attr"`
	Center     Coordinate  `xml:"center>point"`
	ShapeStyle *ShapeStyle `xml:"style"`
}

func (c Circle) Draw(ctx *context) error {
	if c.ShapeStyle != nil {
		c.ShapeStyle.set(ctx)
	}
	ctx.pdf.Circle(c.Center.X, c.Center.Y, c.Radius, c.DrawStyle.String())
	return ctx.pdf.Error()
}

type Polygon struct {
	DrawStyle  DrawStyle    `xml:"style,attr"`
	ShapeStyle *ShapeStyle  `xml:"style"`
	Points     []Coordinate `xml:"points>point"`
}

func (p Polygon) Draw(ctx *context) error {
	if p.ShapeStyle != nil {
		p.ShapeStyle.set(ctx)
	}
	points := []gofpdf.PointType{}
	for _, point := range p.Points {
		points = append(points, point.toPointType())
	}
	ctx.pdf.Polygon(points, p.DrawStyle.String())
	return ctx.pdf.Error()
}

type Beziergon struct {
	DrawStyle  DrawStyle    `xml:"style,attr"`
	ShapeStyle *ShapeStyle  `xml:"style"`
	Points     []Coordinate `xml:"points>point"`
}

func (p Beziergon) Draw(ctx *context) error {
	if p.ShapeStyle != nil {
		p.ShapeStyle.set(ctx)
	}
	points := []gofpdf.PointType{}
	for _, point := range p.Points {
		points = append(points, point.toPointType())
	}
	ctx.pdf.Beziergon(points, p.DrawStyle.String())
	return ctx.pdf.Error()
}

type Ellipse struct {
	DegRotate  float64     `xml:"deg-rotate,attr"`
	DrawStyle  DrawStyle   `xml:"style,attr"`
	Radius     Coordinate  `xml:"radius>point"`
	Center     Coordinate  `xml:"center>point"`
	ShapeStyle *ShapeStyle `xml:"style"`
}

func (c Ellipse) Draw(ctx *context) error {
	if c.ShapeStyle != nil {
		c.ShapeStyle.set(ctx)
	}
	ctx.pdf.Ellipse(c.Center.X, c.Center.Y, c.Radius.X, c.Radius.Y, c.DegRotate, c.DrawStyle.String())
	return ctx.pdf.Error()
}

type Arc struct {
	DegRotate  float64     `xml:"deg-rotate,attr"`
	DegStart   float64     `xml:"deg-start,attr"`
	DegEnd     float64     `xml:"deg-end,attr"`
	DrawStyle  DrawStyle   `xml:"style,attr"`
	Radius     Coordinate  `xml:"radius>point"`
	Center     Coordinate  `xml:"center>point"`
	ShapeStyle *ShapeStyle `xml:"style"`
}

func (c Arc) Draw(ctx *context) error {
	if c.ShapeStyle != nil {
		c.ShapeStyle.set(ctx)
	}
	ctx.pdf.Arc(c.Center.X, c.Center.Y, c.Radius.X, c.Radius.Y, c.DegRotate, c.DegStart, c.DegEnd, c.DrawStyle.String())
	return ctx.pdf.Error()
}

type Curve struct {
	DrawStyle  DrawStyle   `xml:"style,attr"`
	Start      Coordinate  `xml:"start>point"`
	Center     Coordinate  `xml:"center>point"`
	End        Coordinate  `xml:"end>point"`
	ShapeStyle *ShapeStyle `xml:"style"`
}

func (c Curve) Draw(ctx *context) error {
	if c.ShapeStyle != nil {
		c.ShapeStyle.set(ctx)
	}
	ctx.pdf.Curve(c.Start.X, c.Start.Y, c.Center.X, c.Center.Y, c.End.X, c.End.Y, c.DrawStyle.String())
	return ctx.pdf.Error()
}

type BezierCurve struct {
	DrawStyle  DrawStyle   `xml:"style,attr"`
	Start      Coordinate  `xml:"start>point"`
	Center1    Coordinate  `xml:"center1>point"`
	Center2    Coordinate  `xml:"center2>point"`
	End        Coordinate  `xml:"end>point"`
	ShapeStyle *ShapeStyle `xml:"style"`
}

func (c BezierCurve) Draw(ctx *context) error {
	if c.ShapeStyle != nil {
		c.ShapeStyle.set(ctx)
	}
	ctx.pdf.CurveBezierCubic(c.Start.X, c.Start.Y, c.Center1.X, c.Center1.Y, c.Center2.X, c.Center2.Y, c.End.X, c.End.Y, c.DrawStyle.String())
	return ctx.pdf.Error()
}

type CubicCurve struct {
	DrawStyle  DrawStyle   `xml:"style,attr"`
	Start      Coordinate  `xml:"start>point"`
	Center1    Coordinate  `xml:"center1>point"`
	Center2    Coordinate  `xml:"center2>point"`
	End        Coordinate  `xml:"end>point"`
	ShapeStyle *ShapeStyle `xml:"style"`
}

func (c CubicCurve) Draw(ctx *context) error {
	if c.ShapeStyle != nil {
		c.ShapeStyle.set(ctx)
	}
	ctx.pdf.CurveCubic(c.Start.X, c.Start.Y, c.Center1.X, c.Center1.Y, c.End.X, c.End.Y, c.Center2.X, c.Center2.Y, c.DrawStyle.String())
	return ctx.pdf.Error()
}
