package main

import (
	"../"
	"encoding/xml"
	"io/ioutil"
	"os"
)

func main() {
	const offset = 75.0
	o, u, s := simple.Landscape, simple.Millimeter, simple.A4
	pdf := simple.Document{
		Format: simple.Format{
			Orientation:   &o,
			Unit:          &u,
			Size:          &s,
			FontDirectory: "",
			OutputFile:    "setLineJoinStyle.pdf",
		},
		Items: simple.ItemList{
			simple.AddPage{},
		},
	}
	x := 35.0
	caps := []simple.CapStyle{simple.Butt, simple.Square, simple.RoundCap}
	joins := []simple.JoinStyle{simple.Bevel, simple.Miter, simple.RoundJoin}
	var draw = func(cap *simple.CapStyle, join *simple.JoinStyle, x0, y0, x1, y1 float64) []simple.Drawable {
		wide, narrow := 30.0, 2.56
		items := []simple.Drawable{
			simple.Path{
				Style: simple.Draw,
				ShapeStyle: &simple.ShapeStyle{
					JoinStyle: join,
					DrawColor: &simple.RGB{
						R: 0x33,
						G: 0x33,
						B: 0x33,
					},
					LineStyle: &simple.LineStyle{
						Width:    &wide,
						CapStyle: cap,
					},
				},
				Start: simple.Coordinate{
					X: x0,
					Y: y0,
				},
				Items: simple.PathItemList{
					simple.LineTo{
						Next: simple.Coordinate{
							X: (x0+x1)/2 + offset,
							Y: (y0 + y1) / 2,
						},
					},
					simple.LineTo{
						Next: simple.Coordinate{
							X: x1,
							Y: y1,
						},
					},
				},
			},
			simple.Path{
				Style: simple.Draw,
				ShapeStyle: &simple.ShapeStyle{
					JoinStyle: join,
					DrawColor: &simple.RGB{
						R: 0xFF,
						G: 0x33,
						B: 0x33,
					},
					LineStyle: &simple.LineStyle{
						Width:    &narrow,
						CapStyle: cap,
					},
				},
				Start: simple.Coordinate{
					X: x0,
					Y: y0,
				},
				Items: simple.PathItemList{
					simple.LineTo{
						Next: simple.Coordinate{
							X: (x0+x1)/2 + offset,
							Y: (y0 + y1) / 2,
						},
					},
					simple.LineTo{
						Next: simple.Coordinate{
							X: x1,
							Y: y1,
						},
					},
				},
			},
		}
		return items
	}
	for i := range caps {
		pdf.Items = append(pdf.Items, draw(&caps[i], &joins[i], x, 50, x, 160)...)
		x += offset
	}
	bytes, err := xml.MarshalIndent(pdf, "", "  ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("joinstyle.xml", bytes, os.ModePerm)
}
