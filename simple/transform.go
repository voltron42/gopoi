package simple

import "encoding/xml"

type Transform struct {
	XMLName         xml.Name           `xml:"transform"`
	Transformations TransformationList `xml:"transformation"`
	Items           ItemList           `xml:"body"`
}

func (t Transform) Draw(ctx *context) error {
	ctx.pdf.TransformBegin()
	for _, transformation := range t.Transformations {
		err := transformation.ApplyTransformation(ctx)
		if err != nil {
			return err
		}
	}
	for _, item := range t.Items {
		err := item.Draw(ctx)
		if err != nil {
			return err
		}
	}
	ctx.pdf.TransformEnd()
	return ctx.pdf.Error()
}

type Transformation interface {
	ApplyTransformation(ctx *context) error
}

type TransformationList []Transformation

type HorizontalMirror struct {
	XMLName xml.Name `xml:"horizontal-mirror"`
	Xaxis   float64  `xml:"x-axis,attr"`
}

func (t HorizontalMirror) ApplyTransformation(ctx *context) error {
	ctx.pdf.TransformMirrorHorizontal(t.Xaxis)
	return ctx.pdf.Error()
}

type PointMirror struct {
	XMLName xml.Name   `xml:"point-mirror"`
	Point   Coordinate `xml:"point"`
}

func (t PointMirror) ApplyTransformation(ctx *context) error {
	ctx.pdf.TransformMirrorPoint(t.Point.X, t.Point.Y)
	return ctx.pdf.Error()
}

type LineMirror struct {
	XMLName xml.Name   `xml:"line-mirror"`
	Angle   float64    `xml:"angle,attr"`
	Point   Coordinate `xml:"point"`
}

func (t LineMirror) ApplyTransformation(ctx *context) error {
	ctx.pdf.TransformMirrorLine(t.Angle, t.Point.X, t.Point.Y)
	return ctx.pdf.Error()
}

type VerticalMirror struct {
	XMLName xml.Name `xml:"vertical-mirror"`
	Yaxis   float64  `xml:"y-axis,attr"`
}

func (t VerticalMirror) ApplyTransformation(ctx *context) error {
	ctx.pdf.TransformMirrorVertical(t.Yaxis)
	return ctx.pdf.Error()
}

type Rotate struct {
	XMLName xml.Name   `xml:"rotate"`
	Angle   float64    `xml:"angle,attr"`
	Point   Coordinate `xml:"point"`
}

func (t Rotate) ApplyTransformation(ctx *context) error {
	ctx.pdf.TransformRotate(t.Angle, t.Point.X, t.Point.Y)
	return ctx.pdf.Error()
}

type Scale struct {
	XMLName  xml.Name   `xml:"scale"`
	ScalePer Coordinate `xml:"scale-percentage"`
	Point    Coordinate `xml:"point"`
}

func (t Scale) ApplyTransformation(ctx *context) error {
	ctx.pdf.TransformScale(t.ScalePer.X, t.ScalePer.Y, t.Point.X, t.Point.Y)
	return ctx.pdf.Error()
}

type ScaleX struct {
	XMLName    xml.Name   `xml:"scale-x"`
	ScaleWidth float64    `xml:"scale-width,attr"`
	Point      Coordinate `xml:"point"`
}

func (t ScaleX) ApplyTransformation(ctx *context) error {
	ctx.pdf.TransformScaleX(t.ScaleWidth, t.Point.X, t.Point.Y)
	return ctx.pdf.Error()
}

type ScaleY struct {
	XMLName     xml.Name   `xml:"scale-y"`
	ScaleHeight float64    `xml:"scale-height,attr"`
	Point       Coordinate `xml:"point"`
}

func (t ScaleY) ApplyTransformation(ctx *context) error {
	ctx.pdf.TransformScaleY(t.ScaleHeight, t.Point.X, t.Point.Y)
	return ctx.pdf.Error()
}

type ScaleXY struct {
	XMLName  xml.Name   `xml:"scale-xy"`
	ScalePer float64    `xml:"scale-percentage,attr"`
	Point    Coordinate `xml:"point"`
}

func (t ScaleXY) ApplyTransformation(ctx *context) error {
	ctx.pdf.TransformScaleXY(t.ScalePer, t.Point.X, t.Point.Y)
	return ctx.pdf.Error()
}

type Skew struct {
	XMLName xml.Name   `xml:"skew"`
	Angle   Coordinate `xml:"angle"`
	Point   Coordinate `xml:"point"`
}

func (t Skew) ApplyTransformation(ctx *context) error {
	ctx.pdf.TransformSkew(t.Angle.X, t.Angle.Y, t.Point.X, t.Point.Y)
	return ctx.pdf.Error()
}

type SkewX struct {
	XMLName xml.Name   `xml:"skew-x"`
	AngleX  float64    `xml:"angle-x,attr"`
	Point   Coordinate `xml:"point"`
}

func (t SkewX) ApplyTransformation(ctx *context) error {
	ctx.pdf.TransformSkewX(t.AngleX, t.Point.X, t.Point.Y)
	return ctx.pdf.Error()
}

type SkewY struct {
	XMLName xml.Name   `xml:"skew-y"`
	AngleY  float64    `xml:"angle-y,attr"`
	Point   Coordinate `xml:"point"`
}

func (t SkewY) ApplyTransformation(ctx *context) error {
	ctx.pdf.TransformSkewY(t.AngleY, t.Point.X, t.Point.Y)
	return ctx.pdf.Error()
}

type Translate struct {
	XMLName xml.Name   `xml:"translate"`
	Point   Coordinate `xml:"point"`
}

func (t Translate) ApplyTransformation(ctx *context) error {
	ctx.pdf.TransformTranslate(t.Point.X, t.Point.Y)
	return ctx.pdf.Error()
}

type TranslateX struct {
	XMLName xml.Name `xml:"translate-x"`
	X       float64  `xml:"x"`
}

func (t TranslateX) ApplyTransformation(ctx *context) error {
	ctx.pdf.TransformTranslateX(t.X)
	return ctx.pdf.Error()
}

type TranslateY struct {
	XMLName xml.Name `xml:"translate-y"`
	Y       float64  `xml:"y"`
}

func (t TranslateY) ApplyTransformation(ctx *context) error {
	ctx.pdf.TransformTranslateY(t.Y)
	return ctx.pdf.Error()
}
