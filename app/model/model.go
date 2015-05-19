package model

import "github.com/jung-kurt/gofpdf"

var pdf *gofpdf.Fpdf

type PdfElement interface {
}

type Document struct {
	Orientation   *Orientation `json:"orientation,attr"`
	Unit          *Unit        `json:"unit,attr"`
	Size          *Size        `json:"size,attr"`
	FontDirectory string       `json:"font-directory,attr"`
	Pages         []Page       `json:"page"`
}

type Orientation int

type Unit int

type Size int

type Page struct {
}
