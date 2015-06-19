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
	for index, orientation := range orientations {
		if orientation == attr.Value {
			*o = Orientation(index)
			return nil
		}
	}
	return errors.New("Invalid orientation: " + attr.Value)
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
	for index, unit := range units {
		if unit == attr.Value {
			*u = Unit(index)
			return nil
		}
	}
	return errors.New("Invalid unit: " + attr.Value)
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
	for index, size := range sizes {
		if size == attr.Value {
			*s = Size(index)
			return nil
		}
	}
	return errors.New("Invalid size: " + attr.Value)
}
