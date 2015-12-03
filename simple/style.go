package simple

type LinearGradient struct {
	Frame  Frame      `xml:"frame"`
	Color1 RGB        `xml:"color1>rgb"`
	Color2 RGB        `xml:"color2>rgb"`
	Point1 Coordinate `xml:"point1>point"`
	Point2 Coordinate `xml:"point2>point"`
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
	Radius float64    `xml:"radius,attr"`
	Frame  Frame      `xml:"frame"`
	Color1 RGB        `xml:"color1>rgb"`
	Color2 RGB        `xml:"color2>rgb"`
	Point1 Coordinate `xml:"point1>point"`
	Point2 Coordinate `xml:"point2>point"`
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
	Color RGB `xml:"rgb"`
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
	Color RGB `xml:"rgb"`
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
	Color RGB `xml:"rgb"`
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
	Left  float64 `xml:"left,attr"`
	Top   float64 `xml:"top,attr"`
	Right float64 `xml:"right,attr"`
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
	Margin float64 `xml:"margin,attr"`
}

func (s SetRightMargin) Draw(ctx *context) error {
	ctx.pdf.SetRightMargin(s.Margin)
	return ctx.pdf.Error()
}

type SetLeftMargin struct {
	Margin float64 `xml:"margin,attr"`
}

func (s SetLeftMargin) Draw(ctx *context) error {
	ctx.pdf.SetLeftMargin(s.Margin)
	return ctx.pdf.Error()
}

type SetTopMargin struct {
	Margin float64 `xml:"margin,attr"`
}

func (s SetTopMargin) Draw(ctx *context) error {
	ctx.pdf.SetTopMargin(s.Margin)
	return ctx.pdf.Error()
}

type SetFontSize struct {
	Size float64 `xml:"size,attr"`
}

func (s SetFontSize) Draw(ctx *context) error {
	ctx.pdf.SetFontSize(s.Size)
	return ctx.pdf.Error()
}

type SetLineWidth struct {
	Width float64 `xml:"width,attr"`
}

func (s SetLineWidth) Draw(ctx *context) error {
	ctx.pdf.SetLineWidth(s.Width)
	return ctx.pdf.Error()
}

type SetAlpha struct {
	Alpha     float64   `xml:"alpha,attr"`
	BlendMode BlendMode `xml:"blendmode,attr"`
}

func (s SetAlpha) Draw(ctx *context) error {
	ctx.pdf.SetAlpha(s.Alpha, s.BlendMode.String())
	return ctx.pdf.Error()
}

type SetFont struct {
	Family string       `xml:"family,attr"`
	Style  SetFontStyle `xml:"style,attr"`
	Size   float64      `xml:"size,attr"`
}

func (s SetFont) Draw(ctx *context) error {
	ctx.pdf.SetFont(s.Family, s.Style.String(), s.Size)
	return ctx.pdf.Error()
}

type SetLineCapStyle struct {
	CapStyle CapStyle `xml:"cap-style,attr"`
}

func (s SetLineCapStyle) Draw(ctx *context) error {
	ctx.pdf.SetLineCapStyle(s.CapStyle.String())
	return ctx.pdf.Error()
}

type SetLineJoinStyle struct {
	JoinStyle JoinStyle `xml:"join-style,attr"`
}

func (s SetLineJoinStyle) Draw(ctx *context) error {
	ctx.pdf.SetLineJoinStyle(s.JoinStyle.String())
	return ctx.pdf.Error()
}

type SetDashPattern struct {
	DashPattern DashPattern `xml:"dash-pattern"`
}

func (s SetDashPattern) Draw(ctx *context) error {
	ctx.pdf.SetDashPattern(s.DashPattern.Array, s.DashPattern.Phase)
	return ctx.pdf.Error()
}
