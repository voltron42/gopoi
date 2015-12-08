package main

import (
	"../"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

func main() {
	const rowCount = 2
	const colCount = 2
	const ptSize = 36
	var x, y, radius, gap, advance float64
	var rgVal int
	var pts []simple.Coordinate
	vertices := func(count int) (res []simple.Coordinate) {
		var pt simple.Coordinate
		res = make([]simple.Coordinate, 0, count)
		mlt := 2.0 * math.Pi / float64(count)
		for j := 0; j < count; j++ {
			pt.Y, pt.X = math.Sincos(float64(j) * mlt)
			res = append(res, simple.Coordinate{
				X: x + radius*pt.X,
				Y: y + radius*pt.Y})
		}
		return
	}
	o, u, s := simple.Portrait, simple.Millimeter, simple.A4
	gap = 12.0
	radius = (210.0 - float64(colCount+1)*gap) / (2.0 * float64(colCount))
	advance = gap + 2.0*radius
	y = 2*gap + simple.PointConvert(ptSize, simple.Millimeter) + radius
	rgVal = 230
	pdf := simple.Document{
		Format: simple.Format{
			Orientation:   &o,
			Unit:          &u,
			Size:          &s,
			FontDirectory: "",
			OutputFile:    "polygon.pdf",
		},
		Items: simple.ItemList{
			simple.AddPage{},
			simple.SetFont{
				Font: simple.Font{
					Family: "Helvetica",
					Style:  simple.SetFontStyle{false, false, false},
					Size:   ptSize,
				},
			},
			simple.SetDrawColor{Color: simple.RGB{R: 0, G: 80, B: 180}},
			simple.SetY{Y: gap},
			simple.CellFormat{
				Point: simple.Coordinate{
					X: 190.0,
					Y: gap,
				},
				TextString:   "Equilateral polygons",
				Border:       simple.Border{},
				NextPosition: simple.NextLine,
				Align: simple.Alignment{
					Horiz: simple.Center,
					Vert:  simple.DefaultVertAlign,
				},
				Fill: false,
			},
		},
	}
	for row := 0; row < rowCount; row++ {
		pdf.Items = append(pdf.Items, simple.SetFillColor{Color: simple.RGB{R: rgVal, G: rgVal, B: 0}})
		rgVal -= 12
		x = gap + radius
		for col := 0; col < colCount; col++ {
			pts = vertices(row*colCount + col + 3)
			pdf.Items = append(pdf.Items, simple.Polygon{Points: pts, DrawStyle: simple.DrawStyle{Fill: true, Draw: true}})
			x += advance
		}
		y += advance
	}
	fmt.Println(pdf)
	bytes, err := json.MarshalIndent(pdf, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
	bytes, err = xml.Marshal(pdf)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("polygon.xml", bytes, os.ModePerm)
}
