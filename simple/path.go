package simple

type Path struct {
	Style PathPaint    `xml:"style,attr"`
	Start Coordinate   `xml:"start>point"`
	Items PathItemList `xml:"items"`
}

func (p Path) Draw(ctx *context) error {
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
	ctx.pdf.ClosePath()
	err = ctx.pdf.Error()
	if err != nil {
		return err
	}
	ctx.pdf.DrawPath(p.Style.String())
	return ctx.pdf.Error()
}

type PathItemList []PathItem

type PathItem interface {
	DrawPath(ctx *context) error
}

type ArcTo struct {
	End       Coordinate `xml:"end>point"`
	Radius    Coordinate `xml:"radius>point"`
	DegRotate float64    `xml:"deg-rotate,attr"`
	DegStart  float64    `xml:"deg-start,attr"`
	DegEnd    float64    `xml:"deg-end,attr"`
}

func (a ArcTo) DrawPath(ctx *context) error {
	ctx.pdf.ArcTo(a.End.X, a.End.Y, a.Radius.X, a.Radius.Y, a.DegRotate, a.DegStart, a.DegEnd)
	return ctx.pdf.Error()
}

/*
func (f *Fpdf) CurveBezierCubicTo(cx0, cy0, cx1, cy1, x, y float64)
func (f *Fpdf) CurveTo(cx, cy, x, y float64)
func (f *Fpdf) LineTo(x, y float64)
*/
