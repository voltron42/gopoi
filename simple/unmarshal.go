package simple

import (
	"../../xtraml/choice"
	"encoding/xml"
	"errors"
	"fmt"
)

var pathItemChoices = choice.ChoiceParser{
	"arc-to": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := ArcTo{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"line-to": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		l := LineTo{}
		err := d.DecodeElement(&l, &start)
		return l, err
	},
	"curve-to": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		c := CurveTo{}
		err := d.DecodeElement(&c, &start)
		return c, err
	},
	"bezier-cubic-curve-to": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		c := BezierCubicCurveTo{}
		err := d.DecodeElement(&c, &start)
		return c, err
	},
}

func (p *PathItemList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return pathItemChoices.ParseList(d, start, func(item interface{}) error {
		pathItem, ok := item.(PathItem)
		if !ok {
			return errors.New("Item is not a valid PathItem.")
		}
		*p = append(*p, pathItem)
		return nil
	})
}

var itemChoices = choice.ChoiceParser{
	"add-font": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := AddFont{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"add-page": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := AddPage{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"image": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := Image{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"path": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := Path{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"line": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := Line{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"rect": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := Rect{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"circle": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := Circle{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"polygon": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := Polygon{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"beziergon": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := Beziergon{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"ellipse": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := Ellipse{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"arc": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := Arc{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"curve": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := Curve{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"bezier-curve": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := BezierCurve{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"cubic-curve": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := CubicCurve{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"linear-gradient": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := LinearGradient{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"radial-gradient": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := RadialGradient{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"set-draw-color": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := SetDrawColor{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"set-fill-color": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := SetFillColor{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"set-text-color": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := SetTextColor{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"set-margins": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := SetMargins{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"set-right-margin": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := SetRightMargin{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"set-left-margin": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := SetLeftMargin{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"set-top-margin": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := SetTopMargin{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"set-font-size": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := SetFontSize{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"set-line-width": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := SetLineWidth{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"set-alpha": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := SetAlpha{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"set-font": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := SetFont{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"set-line-cap-style": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := SetLineCapStyle{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"set-line-join-style": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := SetLineJoinStyle{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"set-dash-pattern": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := SetDashPattern{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"layer": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := Layer{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
	"text": func(d *xml.Decoder, start xml.StartElement) (interface{}, error) {
		a := Text{}
		err := d.DecodeElement(&a, &start)
		return a, err
	},
}

func (i *ItemList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return itemChoices.ParseList(d, start, func(item interface{}) error {
		fmt.Printf("item: %v\n", item)
		pathItem, ok := item.(Drawable)
		if !ok {
			return errors.New("Item is not a valid Drawable.")
		}
		*i = append(*i, pathItem)
		return nil
	})
}
