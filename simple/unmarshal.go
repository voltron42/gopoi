package simple

import (
	"../../xtraml/choice"
	"encoding/xml"
)

var pathItemChoices = choice.ChoiceParser{
	"arc-to": func() interface{} {
		return &ArcTo{}
	},
	"line-to": func() interface{} {
		return &LineTo{}
	},
	"curve-to": func() interface{} {
		return &CurveTo{}
	},
	"bezier-cubic-curve-to": func() interface{} {
		return &BezierCubicCurveTo{}
	},
}

func (p *PathItemList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var pathItem *PathItem
	return pathItemChoices.ParseList(d, start, p, pathItem, choice.Append)
}

func (p PathItemList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return choice.WrapList(e, xml.Name{Local: "path-items"}, p)
}

var itemChoices = choice.ChoiceParser{
	"add-font": func() interface{} {
		return &AddFont{}
	},
	"add-page": func() interface{} {
		return &AddPage{}
	},
	"image": func() interface{} {
		return &Image{}
	},
	"path": func() interface{} {
		return &Path{}
	},
	"line": func() interface{} {
		return &Line{}
	},
	"rect": func() interface{} {
		return &Rect{}
	},
	"circle": func() interface{} {
		return &Circle{}
	},
	"polygon": func() interface{} {
		return &Polygon{}
	},
	"beziergon": func() interface{} {
		return &Beziergon{}
	},
	"ellipse": func() interface{} {
		return &Ellipse{}
	},
	"arc": func() interface{} {
		return &Arc{}
	},
	"curve": func() interface{} {
		return &Curve{}
	},
	"bezier-curve": func() interface{} {
		return &BezierCurve{}
	},
	"cubic-curve": func() interface{} {
		return &CubicCurve{}
	},
	"linear-gradient": func() interface{} {
		return &LinearGradient{}
	},
	"radial-gradient": func() interface{} {
		return &RadialGradient{}
	},
	"set-draw-color": func() interface{} {
		return &SetDrawColor{}
	},
	"set-fill-color": func() interface{} {
		return &SetFillColor{}
	},
	"set-text-color": func() interface{} {
		return &SetTextColor{}
	},
	"set-margins": func() interface{} {
		return &SetMargins{}
	},
	"set-right-margin": func() interface{} {
		return &SetRightMargin{}
	},
	"set-left-margin": func() interface{} {
		return &SetLeftMargin{}
	},
	"set-top-margin": func() interface{} {
		return &SetTopMargin{}
	},
	"set-font-size": func() interface{} {
		return &SetFontSize{}
	},
	"set-line-width": func() interface{} {
		return &SetLineWidth{}
	},
	"set-alpha": func() interface{} {
		return &SetAlpha{}
	},
	"set-font": func() interface{} {
		return &SetFont{}
	},
	"set-line-cap-style": func() interface{} {
		return &SetLineCapStyle{}
	},
	"set-line-join-style": func() interface{} {
		return &SetLineJoinStyle{}
	},
	"set-dash-pattern": func() interface{} {
		return &SetDashPattern{}
	},
	"layer": func() interface{} {
		return &Layer{}
	},
	"text": func() interface{} {
		return &Text{}
	},
	"set-x": func() interface{} {
		return &SetX{}
	},
	"set-xy": func() interface{} {
		return &SetXY{}
	},
	"set-y": func() interface{} {
		return &SetY{}
	},
	"cell-format": func() interface{} {
		return &CellFormat{}
	},
}

func (i *ItemList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var item *Drawable
	return itemChoices.ParseList(d, start, i, item, choice.Append)
}

func (p ItemList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return choice.WrapList(e, xml.Name{Local: "body"}, p)
}
