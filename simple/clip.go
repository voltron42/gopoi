package simple

import (
	"encoding/xml"
	"github.com/jung-kurt/gofpdf"
)

type Clip struct {
	XMLName xml.Name    `xml:"clip"`
	Clipper ClipWrapper `xml:"clip-item"`
	Items   ItemList    `xml:"body"`
}

func (c Clip) Draw(ctx *context) error {
	err := c.Clipper.ApplyClipItem(ctx)
	if err != nil {
		return err
	}
	for _, item := range c.Items {
		err := item.Draw(ctx)
		if err != nil {
			return err
		}
	}
	ctx.pdf.ClipEnd()
	return ctx.pdf.Error()
}

type ClipWrapper struct {
	ClipItem
}

type ClipItem interface {
	ApplyClipItem(ctx *context) error
}

type ClipCircle struct {
	XMLName xml.Name   `xml:"clip-circle"`
	Radius  float64    `xml:"radius,attr"`
	Center  Coordinate `xml:"center"`
	Outline bool       `xml:"outline,attr"`
}

func (c ClipCircle) ApplyClipItem(ctx *context) error {
	ctx.pdf.ClipCircle(c.Center.X, c.Center.Y, c.Radius, c.Outline)
	return ctx.pdf.Error()
}

type ClipEllipse struct {
	XMLName xml.Name   `xml:"clip-ellipse"`
	Radius  Coordinate `xml:"radius"`
	Center  Coordinate `xml:"center"`
	Outline bool       `xml:"outline,attr"`
}

func (c ClipEllipse) ApplyClipItem(ctx *context) error {
	ctx.pdf.ClipEllipse(c.Center.X, c.Center.Y, c.Radius.X, c.Radius.Y, c.Outline)
	return ctx.pdf.Error()
}

type ClipPolygon struct {
	XMLName xml.Name     `xml:"clip-polygon"`
	Points  []Coordinate `xml:"points>point"`
	Outline bool         `xml:"outline,attr"`
}

func (c ClipPolygon) ApplyClipItem(ctx *context) error {
	points := []gofpdf.PointType{}
	for _, point := range c.Points {
		points = append(points, point.toPointType())
	}
	ctx.pdf.ClipPolygon(points, c.Outline)
	return ctx.pdf.Error()
}

type ClipRect struct {
	XMLName xml.Name `xml:"clip-rect"`
	Frame   Frame    `xml:"frame"`
	Outline bool     `xml:"outline,attr"`
}

func (c ClipRect) ApplyClipItem(ctx *context) error {
	ctx.pdf.ClipRect(c.Frame.X, c.Frame.Y, c.Frame.W, c.Frame.H, c.Outline)
	return ctx.pdf.Error()
}

type ClipRoundedRect struct {
	XMLName xml.Name `xml:"clip-rounded-rect"`
	Radius  float64  `xml:"radius,attr"`
	Frame   Frame    `xml:"frame"`
	Outline bool     `xml:"outline,attr"`
}

func (c ClipRoundedRect) ApplyClipItem(ctx *context) error {
	ctx.pdf.ClipRoundedRect(c.Frame.X, c.Frame.Y, c.Frame.W, c.Frame.H, c.Radius, c.Outline)
	return ctx.pdf.Error()
}

type ClipText struct {
	XMLName xml.Name   `xml:"clip-text"`
	Text    string     `xml:"text,attr"`
	Point   Coordinate `xml:"point"`
	Outline bool       `xml:"outline,attr"`
}

func (c ClipText) ApplyClipItem(ctx *context) error {
	ctx.pdf.ClipText(c.Point.X, c.Point.Y, c.Text, c.Outline)
	return ctx.pdf.Error()
}
