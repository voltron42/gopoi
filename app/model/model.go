package model

import (
	"../../../prism"
)

type Document struct {
	Orientation   Orientation `json:"orientation,attr"`
	Unit          Unit        `json:"unit,attr"`
	Size          Size        `json:"size,attr"`
	FontDirectory string      `json:"font-directory,attr"`

	Metadata *Metadata     `json:"meta"`
	Header   []PageElement `json:"header>elements"`
	Footer   []PageElement `json:"footer>elements"`
	Pages    []Page        `json:"page"`
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
	PageFormat   *PageFormat   `xml:"format"`
	PageElements []PageElement `xml:"elements"`
}

type PageFormat struct {
	Orientation Orientation `json:"orientation,attr"`
	Width       floatstr    `xml:"width,attr"`
	Height      floatstr    `xml:"height,attr"`
}

type PageElements struct {
	Elements []PageElement
}

type PageElement struct {
	Bookmark  *Bookmark  `xml:"bookmark"`
	LineBreak *LineBreak `xml:"line-break"`
	Block     *Block     `xml:"block"`
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
	RGB   *prism.RGB `xml:"rgb"`
	Named *Named     `xml:"named"`
	Hex   *Hex       `xml:"hex"`
}

type Named struct {
	Name string `xml:"color,attr"`
}

type Hex struct {
	Code string `xml:",chardata"`
}

type Transform struct {
	Skew      *Skew       `xml:"skew"`
	Rotate    *Rotate     `xml:"rotate"`
	Translate *Coordinate `xml:"translate>point"`
	Mirror    *Line       `xml:"mirror>line"`
	Scale     *Scale      `xml:"scale"`
}

type Skew struct {
	Angle   Coordinate `xml:"angle>point"`
	Station Coordinate `xml:"station>point"`
}

type Rotate struct {
	Angle  floatstr   `xml:"angle,attr"`
	Center Coordinate `xml:"center>point"`
}

type Scale struct {
	Factor Coordinate `xml:"factor>point"`
	Center Coordinate `xml:"center>point"`
}

type Clip struct {
	// todo --
}

type Object struct {
}

type Image struct {
}

type Polygon struct {
	Fill    boolstr `xml:"fill,attr"`
	Outline boolstr `xml:"outline,attr"`

	Points []Coordinate `xml:"point"`
}

type Beziergon struct {
	Fill    boolstr `xml:"fill,attr"`
	Outline boolstr `xml:"outline,attr"`

	Points []Coordinate `xml:"point"`
}

type Curve struct {
	Type    string  `xml:"type,attr"`
	Fill    boolstr `xml:"fill,attr"`
	Outline boolstr `xml:"outline,attr"`

	Points []Coordinate `xml:"point"`
}

type Line struct {
	Bounds Bounds `xml:"bounds"`
}

type Bounds struct {
	Coordinate Coordinate `xml:"point"`
	Opposite   Coordinate `xml:"opposite>point"`
}

type Coordinate struct {
	X floatstr `xml:"x,attr"`
	Y floatstr `xml:"y,attr"`
}
