package model

import (
	"../../../prism"
	"../../../prism/hex"
	"../../../prism/named"
	"../../../sandbox/xmltogo"
	"encoding/xml"
	"errors"
	"strconv"
)

type Document struct {
	Orientation   Orientation `json:"orientation,attr"`
	Unit          Unit        `json:"unit,attr"`
	Size          Size        `json:"size,attr"`
	FontDirectory string      `json:"font-directory,attr"`

	Metadata *Metadata    `json:"meta"`
	Header   PageElements `json:"header>elements"`
	Footer   PageElements `json:"footer>elements"`
	Pages    []Page       `json:"page"`
}

type Orientation int

func (o *Orientation) String() string {
	return ""
}

type Unit int

func (u *Unit) String() string {
	return ""
}

type Size int

func (s *Size) String() string {
	return ""
}

type Metadata struct {
	Keywords    *Keywords    `xml:"keywords"`
	CustomFonts []CustomFont `xml:"custom-fonts>custom-font"`
	Layers      []Layer      `xml:"layers>layer"`
	Links       []string     `xml:"links>link-ref"`
	Margins     Margins      `xml:"margins"`
}

type Keywords struct {
	UTF8  boolstr  `xml:"utf8,attr"`
	Words []string `xml:"keyword"`
}

type CustomFont struct {
	Family   string  `xml:"family,attr"`
	Filepath string  `xml:"filepath,attr"`
	Bold     boolstr `xml:"bold,attr"`
	Italic   boolstr `xml:"italic,attr"`
}

type Layer struct {
	Name    string  `xml:"name,attr"`
	Visible boolstr `xml:"visible,attr"`
}

type Margins struct {
	Top   floatstr `xml:"top,attr"`
	Left  floatstr `xml:"left,attr"`
	Right floatstr `xml:"right,attr"`
}

type Page struct {
	PageFormat   *PageFormat  `xml:"format"`
	PageElements PageElements `xml:"elements"`
}

type PageFormat struct {
	Orientation Orientation `json:"orientation,attr"`
	Width       floatstr    `xml:"width,attr"`
	Height      floatstr    `xml:"height,attr"`
}

type PageElements struct {
	Elements []PageElement
}

func (p *PageElements) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return pageElementMarshaller.MarshalChildren(d, start, func(item interface{}) error {
		elem, ok := item.(PageElement)
		if ok {
			p.Elements = append(p.Elements, elem)
			return nil
		} else {
			return errors.New("Cannot cast item to page element")
		}
	})
}

type PageElement interface {
}

var pageElementMarshaller = xmltogo.InterfaceMarshaller{
	ChildMap: map[string]func() interface{}{
		"bookmark":   func() interface{} { return Bookmark{} },
		"line-break": func() interface{} { return LineBreak{} },
		"block":      func() interface{} { return Block{} },
	},
}

type Bookmark struct {
	Name     string   `xml:"name,attr"`
	Level    intstr   `xml:"level,attr"`
	Position floatstr `xml:"position,attr"`
}

type LineBreak struct {
	Position floatstr `xml:"position,attr"`
}

type Block struct {
	Layer string `xml:"layer,attr"`

	Setting    *Setting    `xml:"setting"`
	Transforms []Transform `xml:"transformations>transform"`
	Clip       *Clip       `xml:"clip"`
	Objects    []Object    `xml:"object"`
}

type Setting struct {
	X         floatstr `xml:"x,attr"`
	Y         floatstr `xml:"y,attr"`
	Margin    floatstr `xml:"margin,attr"`
	LineWidth floatstr `xml:"line-width,attr"`

	Font      *Font  `xml:"font"`
	TextColor *Color `xml:"text-color"`
	DrawColor *Color `xml:"draw-color"`
	FillColor *Color `xml:"fill-color"`
}

type Font struct {
	Family string  `xml:"family,attr"`
	Size   intstr  `xml:"size,attr"`
	Bold   boolstr `xml:"bold,attr"`
	Italic boolstr `xml:"italic,attr"`
}

type Color struct {
	RGB prism.RGB
}

func (c *Color) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	colors := []prism.RGB{}
	colorMarshaller.MarshalChildren(d, start, func(item interface{}) error {
		color := prism.RGB{}
		elem, ok := item.(prism.RGB)
		if ok {
			color = elem
		} else {
			var r, g, b uint32
			elem, ok := item.(Named)
			if ok {
				namedColor := namedFactory.GetColor(elem.Name)
				r, g, b, _ = namedColor.RGBA()
			} else {
				elem, ok := item.(Hex)
				if ok {
					hexColor := hexFactory.GetColor(elem.Code)
					r, g, b, _ = hexColor.RGBA()
				} else {
					return errors.New("cannot decode")
				}
			}
			color.R = r
			color.G = g
			color.B = b
		}
		colors = append(colors, color)
		return nil
	})
	return nil
}

var hexFactory = hex.NewFactory()
var namedFactory = named.NewFactory()

var colorMarshaller = xmltogo.InterfaceMarshaller{
	ChildMap: map[string]func() interface{}{
		"rgb":   func() interface{} { return prism.RGB{} },
		"named": func() interface{} { return Named{} },
		"hex":   func() interface{} { return Hex{} },
	},
}

type Named struct {
	Name string `xml:"color,attr"`
}

type Hex struct {
	Code string `xml:",chardata"`
}

type Transform struct {
	Skew      *Skew   `xml:"skew"`
	Rotate    *Rotate `xml:"rotate"`
	Translate *Point  `xml:"translate>point"`
	Mirror    *Line   `xml:"mirror>line"`
	Scale     *Scale  `xml:"scale"`
}

type Skew struct {
	Angle   Point `xml:"angle>point"`
	Station Point `xml:"station>point"`
}

type Rotate struct {
	Angle  floatstr `xml:"angle,attr"`
	Center Point    `xml:"center>point"`
}

type Scale struct {
	Factor Point `xml:"factor>point"`
	Center Point `xml:"center>point"`
}

type Clip struct {
	// todo --
}

type Object struct {
}

type Line struct {
	Bounds Bounds `xml:"bounds"`
}

type Bounds struct {
	Point    Point `xml:"point"`
	Opposite Point `xml:"opposite>point"`
}

type Point struct {
	X floatstr `xml:"x,attr"`
	Y floatstr `xml:"y,attr"`
}

type intstr int

func (i *intstr) UnmarshalXMLAttr(attr xml.Attr) error {
	num, err := strconv.Atoi(attr.Value)
	if err != nil {
		return err
	}
	*i = intstr(num)
	return nil
}

type floatstr float64

func (f *floatstr) UnmarshalXMLAttr(attr xml.Attr) error {
	num, err := strconv.ParseFloat(attr.Value, 64)
	if err != nil {
		return err
	}
	*f = floatstr(num)
	return nil
}

type boolstr bool

func (b *boolstr) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "yes":
		*b = boolstr(true)
	case "no":
		*b = boolstr(false)
	default:
		return errors.New("invalid value")
	}
	return nil
}
