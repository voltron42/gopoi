package simple

import (
	"encoding/xml"
	"strings"
)

type Orientation int

const (
	Landscape Orientation = iota
	Portrait
	defaultOrientation
)

var orientations = []string{"Landscape", "Portrait", ""}

func (o *Orientation) String() string {
	return orientations[int(*o)]
}

func (o *Orientation) UnmarshalXMLAttr(attr xml.Attr) error {
	switch strings.ToLower(attr.Value) {
	case "l":
		*o = Landscape
	case "landscape":
		*o = Landscape
	case "p":
		*o = Portrait
	case "portrait":
		*o = Portrait
	default:
		*o = defaultOrientation
	}
	return nil
}

type Unit int

const (
	Point Unit = iota
	Millimeter
	Centimeter
	Inch
	defaultUnit
)

var units = []string{"pt", "mm", "cm", "in", ""}

func (u *Unit) String() string {
	return units[int(*u)]
}

func (u *Unit) UnmarshalXMLAttr(attr xml.Attr) error {
	switch strings.ToLower(attr.Value) {
	case "pt":
		*u = Point
	case "mm":
		*u = Millimeter
	case "cm":
		*u = Centimeter
	case "in":
		*u = Inch
	default:
		*u = defaultUnit
	}
	return nil
}

type Size int

const (
	A3 Size = iota
	A4
	A5
	Letter
	Legal
	defaultSize
)

var sizes = []string{"A3", "A4", "A5", "Letter", "Legal", ""}

func (s *Size) String() string {
	return sizes[int(*s)]
}

func (s *Size) UnmarshalXMLAttr(attr xml.Attr) error {
	switch strings.ToLower(attr.Value) {
	case "a3":
		*s = A3
	case "a4":
		*s = A4
	case "a5":
		*s = A5
	case "letter":
		*s = Letter
	case "legal":
		*s = Legal
	default:
		*s = defaultSize
	}
	return nil
}

type FontStyle int

const (
	Bold FontStyle = iota
	Italic
	BoldItalic
	normal
)

var fontStyles = []string{"B", "I", "BI", ""}

func (f *FontStyle) String() string {
	return fontStyles[int(*f)]
}

func (f *FontStyle) UnmarshalXMLAttr(attr xml.Attr) error {
	switch strings.ToLower(attr.Value) {
	case "b":
		*f = Bold
	case "bold":
		*f = Bold
	case "i":
		*f = Italic
	case "italic":
		*f = Italic
	case "bi":
		*f = BoldItalic
	case "bolditalic":
		*f = BoldItalic
	default:
		*f = normal
	}
	return nil
}

type ImageFormat int

const (
	JPG ImageFormat = iota
	JPEG
	PNG
	GIF
	defaultFormat
)

var formats = []string{"JPG", "JPEG", "PNG", "GIF", ""}

func (i *ImageFormat) String() string {
	return formats[int(*i)]
}

func (i *ImageFormat) UnmarshalXMLAttr(attr xml.Attr) error {
	switch strings.ToLower(attr.Value) {
	case "jpg":
		*i = JPG
	case "jpeg":
		*i = JPEG
	case "png":
		*i = PNG
	case "gif":
		*i = GIF
	default:
		*i = defaultFormat
	}
	return nil
}

type PathPaint int

const (
	Fill PathPaint = iota
	Draw
	DrawFill
	Stroke
	StrokeClose
	FillNonZero
	FillEvenOdd
	FillStrokeNonZero
	FillStrokeEvenOdd
	FillStrokeCloseNonZero
	FillStrokeCloseEvenOdd
	defaultPaint
)

var paints = []string{"F", "D", "FD", "S", "s", "f", "f*", "B", "B*", "b", "b*", ""}

func (p *PathPaint) String() string {
	return paints[int(*p)]
}

func (p *PathPaint) UnmarshalXMLAttr(attr xml.Attr) error {
	switch strings.ToLower(attr.Value) {
	case "fill":
		*p = Fill
	case "draw":
		*p = Draw
	case "drawfill":
		*p = DrawFill
	case "stroke":
		*p = Stroke
	case "strokeclose":
		*p = StrokeClose
	case "fillnonzero":
		*p = FillNonZero
	case "fillevenodd":
		*p = FillEvenOdd
	case "fillstrokenonzero":
		*p = FillStrokeNonZero
	case "fillstrokeevenodd":
		*p = FillStrokeEvenOdd
	case "fillstrokeclosenonzero":
		*p = FillStrokeCloseNonZero
	case "fillstrokecloseevenodd":
		*p = FillStrokeCloseEvenOdd
	default:
		*p = defaultPaint
	}
	return nil
}

type BlendMode int

const (
	Normal BlendMode = iota
	Multiply
	Screen
	Overlay
	Darken
	Lighten
	ColorDodge
	ColorBurn
	HardLight
	SoftLight
	Difference
	Exclusion
	Hue
	Saturation
	Color
	Luminosity
	defaultBlendMode
)

var blendModes = []string{
	"Normal",
	"Multiply",
	"Screen",
	"Overlay",
	"Darken",
	"Lighten",
	"ColorDodge",
	"ColorBurn",
	"HardLight",
	"SoftLight",
	"Difference",
	"Exclusion",
	"Hue",
	"Saturation",
	"Color",
	"Luminosity",
}

func (b *BlendMode) String() string {
	return blendModes[int(*b)]
}

func (b *BlendMode) UnmarshalXMLAttr(attr xml.Attr) error {
	switch strings.ToLower(attr.Value) {
	case "normal":
		*b = Normal
	case "multiply":
		*b = Multiply
	case "screen":
		*b = Screen
	case "overlay":
		*b = Overlay
	case "darken":
		*b = Darken
	case "lighten":
		*b = Lighten
	case "colordodge":
		*b = ColorDodge
	case "Colorburn":
		*b = ColorBurn
	case "hardlight":
		*b = HardLight
	case "softlight":
		*b = SoftLight
	case "difference":
		*b = Difference
	case "exclusion":
		*b = Exclusion
	case "hue":
		*b = Hue
	case "saturation":
		*b = Saturation
	case "color":
		*b = Color
	case "luminosity":
		*b = Luminosity
	default:
		*b = defaultBlendMode
	}
	return nil
}

type CapStyle int

const (
	Butt CapStyle = iota
	RoundCap
	Square
	defaultCapStyle
)

var capstyles = []string{"butt", "round", "square", ""}

func (c *CapStyle) String() string {
	return capstyles[int(*c)]
}

func (c *CapStyle) UnmarshalXMLAttr(attr xml.Attr) error {
	switch strings.ToLower(attr.Value) {
	case "butt":
		*c = Butt
	case "round":
		*c = RoundCap
	case "square":
		*c = Square
	default:
		*c = defaultCapStyle
	}
	return nil
}

type JoinStyle int

const (
	Miter JoinStyle = iota
	RoundJoin
	Bevel
	defaultJoinStyle
)

var joinstyles = []string{"miter", "round", "bevel", ""}

func (c *JoinStyle) String() string {
	return joinstyles[int(*c)]
}

func (c *JoinStyle) UnmarshalXMLAttr(attr xml.Attr) error {
	switch strings.ToLower(attr.Value) {
	case "miter":
		*c = Miter
	case "round":
		*c = RoundJoin
	case "bevel":
		*c = Bevel
	default:
		*c = defaultJoinStyle
	}
	return nil
}
