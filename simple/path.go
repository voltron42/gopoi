package simple

type Path struct {
	Style          PathPaint
	StartX, StartY float64
	Items          []PathItem
}

func (p Path) Draw(ctx *context) error {
	ctx.pdf.MoveTo(p.StartX, p.StartY)
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

type PathItem interface {
	DrawPath(ctx *context) error
}

type ArcTo struct {
	X, Y, RX, RY, DegRotate, DegStart, DegEnd float64
}

func (a ArcTo) DrawPath(ctx *context) error {
	ctx.pdf.ArcTo(a.X, a.Y, a.RX, a.RY, a.DegRotate, a.DegStart, a.DegEnd)
	return ctx.pdf.Error()
}

/*
func (f *Fpdf) CurveBezierCubicTo(cx0, cy0, cx1, cy1, x, y float64)
func (f *Fpdf) CurveTo(cx, cy, x, y float64)
func (f *Fpdf) LineTo(x, y float64)
*/
