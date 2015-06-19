package model

import (
	"encoding/xml"
	"errors"
	"strconv"
)

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
