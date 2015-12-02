package model

import (
	"encoding/xml"
	"errors"
	"strings"
)

type Orientation int

const (
	Portrait Orientation = iota
	Landscape
)

var orientations = strings.Split("portrait|landscape", "|")

func (o *Orientation) String() string {
	return orientations[*o]
}

func (o *Orientation) UnmarshalXMLAttr(attr xml.Attr) error {
	index, err := find(orientations, attr.Value, "orientation")
	if err != nil {
		return err
	}
	*o = Orientation(index)
	return nil
}

type Unit int

const (
	Millimeter Unit = iota
	Centimeter
	Point
	Inch
)

var units = strings.Split("millimeter|centimeter|point|inch", "|")

func (u *Unit) String() string {
	return units[*u]
}

func (u *Unit) UnmarshalXMLAttr(attr xml.Attr) error {
	index, err := find(units, attr.Value, "unit")
	if err != nil {
		return err
	}
	*u = Unit(index)
	return nil
}

type Size int

const (
	A4 Size = iota
	A3
	A5
	Letter
	Legal
)

var sizes = strings.Split("A4|A3|A5|letter|legal", "|")

func (s *Size) String() string {
	return sizes[*s]
}

func (s *Size) UnmarshalXMLAttr(attr xml.Attr) error {
	index, err := find(sizes, attr.Value, "size")
	if err != nil {
		return err
	}
	*s = Size(index)
	return nil
}

type VerticalOrientation int

type ImageFormat int

const (
	Jpg ImageFormat = iota
	Jpeg
	Png
	Gif
)

var imageFormats = strings.Split("jpg|jpeg|png|gif", "|")

func (i *ImageFormat) String() string {
	return imageFormats[*i]
}

func (i *ImageFormat) UnmarshalXMLAttr(attr xml.Attr) error {
	index, err := find(imageFormats, attr.Value, "image-format")
	if err != nil {
		return err
	}
	*i = ImageFormat(index)
	return nil
}

type CurveType int

const (
	Bezier CurveType = iota
	Cubic
)

var curveTypes = strings.Split("bezier|cubic", "|")

func (c *CurveType) String() string {
	return curveTypes[*c]
}

func (c *CurveType) UnmarshalXMLAttr(attr xml.Attr) error {
	index, err := find(curveTypes, attr.Value, "curveType")
	if err != nil {
		return err
	}
	*c = CurveType(index)
	return nil
}

func find(haystack []string, needle string, label string) (int, error) {
	for index, item := range haystack {
		if item == needle {
			return index, nil
		}
	}
	return -1, errors.New("Invalid " + label + ": " + needle)
}
