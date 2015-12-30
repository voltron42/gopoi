package main

import (
	"../"
	"encoding/xml"
	"io/ioutil"
	"os"
)

func main() {
	o, u, s := simple.Portrait, simple.Millimeter, simple.A4
	format := simple.Format{
		Orientation:   &o,
		Unit:          &u,
		Size:          &s,
		FontDirectory: "",
		OutputFile:    "transform.pdf",
	}
	titleStr := "Transformations"
	titlePt := 36.0
	titleHt := simple.PointConvert(titlePt, u)
	font := simple.Font{
		Family: "Helvetica",
		Size:   titlePt,
	}
	titleWd := simple.GetStringWidth(titleStr, &format, font)
	titleX := (210 - titleWd) / 2
	transforms := []transRef{
		transRef{
			ref: ref{
				refStr: "Scale",
				refX:   50,
				refY:   60,
			},
			transform: simple.ScaleXY{
				ScalePer: 150,
				Point:    simple.Coordinate{X: 50, Y: 70},
			},
		},
		transRef{
			ref: ref{
				refStr: "Translate",
				refX:   125,
				refY:   60,
			},
			transform: simple.Translate{
				Point: simple.Coordinate{X: 7, Y: 5},
			},
		},
		transRef{
			ref: ref{
				refStr: "Rotate",
				refX:   50,
				refY:   110,
			},
			transform: simple.Rotate{
				Angle: 20,
				Point: simple.Coordinate{X: 50, Y: 120},
			},
		},
		transRef{
			ref: ref{
				refStr: "Skew",
				refX:   125,
				refY:   110,
			},
			transform: simple.SkewX{
				AngleX: 30,
				Point:  simple.Coordinate{X: 125, Y: 110},
			},
		},
		transRef{
			ref: ref{
				refStr: "Mirror horizontal",
				refX:   50,
				refY:   160,
			},
			transform: simple.HorizontalMirror{
				Xaxis: 50,
			},
		},
		transRef{
			ref: ref{
				refStr: "Mirror vertical",
				refX:   125,
				refY:   160,
			},
			transform: simple.VerticalMirror{
				Yaxis: 170,
			},
		},
	}
	pdf := simple.Document{
		Format: format,
		Items: simple.ItemList{
			simple.AddPage{},
			simple.SetFont{
				Font: font,
			},
			simple.Text{
				TextString: titleStr,
				Point:      simple.Coordinate{X: titleX, Y: 10 + titleHt},
			},
			simple.Transform{
				Transformations: simple.TransformationList{
					simple.VerticalMirror{
						Yaxis: 10 + titleHt + 0.5,
					},
				},
				Items: simple.ItemList{
					simple.Clip{
						Clipper: simple.ClipWrapper{
							ClipItem: simple.ClipText{
								Text:  titleStr,
								Point: simple.Coordinate{X: titleX, Y: 10 + titleHt},
							},
						},
						Items: simple.ItemList{
							simple.LinearGradient{
								Frame:  simple.Frame{X: titleX, Y: 10, W: titleWd, H: titleHt + 4},
								Color1: simple.RGB{R: 120, G: 120, B: 120},
								Color2: simple.RGB{R: 255, G: 255, B: 255},
								Point1: simple.Coordinate{X: 0, Y: 0},
								Point2: simple.Coordinate{X: 0, Y: 0.6},
							},
						},
					},
				},
			},
			simple.SetFont{
				Font: simple.Font{Family: "Helvetica", Size: 12},
			},
		},
	}
	for _, transform := range transforms {
		pdf.Items = append(pdf.Items, transform.draw()...)
	}
	bytes, err := xml.MarshalIndent(pdf, "", "  ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("transform.xml", bytes, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func color(val int) simple.ItemList {
	c := simple.RGB{R: val, G: val, B: val}
	return simple.ItemList{
		simple.SetDrawColor{
			Color: c,
		},
		simple.SetTextColor{
			Color: c,
		},
	}
}

func reference(str string, x, y float64, val int) simple.ItemList {
	return append(color(val), simple.Rect{
		DrawStyle: simple.DrawStyle{Draw: true},
		Frame:     simple.Frame{X: x, Y: y, W: 40, H: 10},
	}, simple.Text{
		TextString: str,
		Point:      simple.Coordinate{X: x, Y: y - 1},
	})
}

const (
	light = 200
	dark  = 0
)

type ref struct {
	refX, refY float64
	refStr     string
}

func (r *ref) draw(shade int) simple.ItemList {
	return reference(r.refStr, r.refX, r.refY, shade)
}

type transRef struct {
	ref       ref
	transform simple.Transformation
}

func (r *transRef) draw() simple.ItemList {
	return append(r.ref.draw(light), simple.Transform{
		Transformations: simple.TransformationList{
			r.transform,
		},
		Items: r.ref.draw(dark),
	})
}
