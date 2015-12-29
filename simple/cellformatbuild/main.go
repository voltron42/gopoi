package main

import (
	"../"
	"bytes"
	"encoding/csv"
	"encoding/xml"
	"io/ioutil"
	"os"
	"strings"
)

var header = []string{"Country", "Capital", "Area (sq km)", "Pop. (thousands)"}

func main() {
	o, u, s := simple.Portrait, simple.Millimeter, simple.A4
	pdf := simple.Document{
		Format: simple.Format{
			Orientation:   &o,
			Unit:          &u,
			Size:          &s,
			FontDirectory: "",
			OutputFile:    "cellFormat.pdf",
		},
		Items: simple.ItemList{
			simple.SetFont{
				Font: simple.Font{
					Family: "Arial",
					Size:   14,
					Style:  simple.SetFontStyle{},
				},
			},
			simple.AddPage{},
		},
	}
	records, err := loadData("countries.csv")
	if err != nil {
		panic(err)
	}
	pdf.Items = append(pdf.Items, basicTable(records)...)
	pdf.Items = append(pdf.Items, simple.AddPage{})
	pdf.Items = append(pdf.Items, improvedTable(records)...)
	pdf.Items = append(pdf.Items, simple.AddPage{})
	pdf.Items = append(pdf.Items, fancyTable(records)...)
	pdf.Items = append(pdf.Items, simple.AddPage{})
	bytes, err := xml.MarshalIndent(pdf, "", "  ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("cellformat1.xml", bytes, os.ModePerm)
}

func basicTable(records []countryType) simple.ItemList {
	items := simple.ItemList{}
	for _, str := range header {
		items = append(items, basicHeader(str))
	}
	items = append(items, simple.LineBreak{Ordinate: -1})
	for _, c := range records {
		items = append(items,
			basicCell(c.nameStr),
			basicCell(c.capitalStr),
			basicCell(c.areaStr),
			basicCell(c.popStr),
			simple.LineBreak{Ordinate: -1},
		)
	}
	return items
}

func basicHeader(text string) simple.CellFormat {
	return coreCell(40, 7, text, false, simple.FullBorder(), simple.Alignment{})
}
func basicCell(text string) simple.CellFormat {
	return coreCell(40, 6, text, false, simple.FullBorder(), simple.Alignment{})
}

func coreCell(x, y float64, text string, fill bool, border simple.Border, align simple.Alignment) simple.CellFormat {
	return simple.CellFormat{
		Point: simple.Coordinate{
			X: x,
			Y: y,
		},
		TextString:   text,
		Border:       border,
		NextPosition: simple.ToTheRight,
		Align:        align,
		Fill:         fill,
	}
}

func cell(x, y float64, text string, top, left, bottom, right bool, vert simple.VertAlign, horiz simple.HorizAlign) simple.CellFormat {
	return coreCell(x, y, text, false, simple.Border{
		Top:    top,
		Left:   left,
		Bottom: bottom,
		Right:  right,
	}, simple.Alignment{
		Horiz: horiz,
		Vert:  vert,
	})
}

func improvedTable(records []countryType) simple.ItemList {
	items := simple.ItemList{}
	// Column widths
	w := []float64{40.0, 35.0, 40.0, 45.0}
	wSum := 0.0
	for _, v := range w {
		wSum += v
	}
	// 	Header
	for j, str := range header {
		items = append(items, coreCell(w[j], 7, str, false, simple.FullBorder(), simple.Alignment{Horiz: simple.Center}))
	}
	items = append(items, simple.LineBreak{Ordinate: -1})
	for _, c := range records {
		items = append(items,
			coreCell(w[0], 6, c.nameStr, false, simple.Border{Left: true, Right: true}, simple.Alignment{}),
			coreCell(w[1], 6, c.capitalStr, false, simple.Border{Left: true, Right: true}, simple.Alignment{}),
			coreCell(w[2], 6, c.areaStr, false, simple.Border{Left: true, Right: true}, simple.Alignment{Horiz: simple.Right}),
			coreCell(w[3], 6, c.popStr, false, simple.Border{Left: true, Right: true}, simple.Alignment{Horiz: simple.Right}),
			simple.LineBreak{Ordinate: -1})
	}
	items = append(items, coreCell(wSum, 0, "", false, simple.Border{Top: true}, simple.Alignment{}))
	return items
}

func fancyTable(records []countryType) simple.ItemList {
	items := simple.ItemList{
		simple.SetFillColor{
			Color: simple.RGB{
				R: 255,
			},
		},
		simple.SetTextColor{
			Color: simple.RGB{
				R: 255,
				G: 255,
				B: 255,
			},
		},
		simple.SetDrawColor{
			Color: simple.RGB{
				R: 128,
			},
		},
		simple.SetLineWidth{
			Width: 0.3,
		},
		simple.SetFont{
			Font: simple.Font{
				Style: simple.SetFontStyle{
					Bold: true,
				},
			},
		},
	}
	w := []float64{40, 35, 40, 45}
	wSum := 0.0
	for _, v := range w {
		wSum += v
	}
	for j, str := range header {
		items = append(items, coreCell(w[j], 7, str, true, simple.FullBorder(), simple.Alignment{Horiz: simple.Center}))
	}
	items = append(items,
		simple.LineBreak{Ordinate: -1},
		simple.SetFillColor{
			Color: simple.RGB{
				R: 224,
				G: 235,
				B: 255,
			},
		},
		simple.SetTextColor{},
		simple.SetFont{},
	)
	fill := false
	for _, c := range records {
		items = append(items,
			coreCell(w[0], 6, c.nameStr, fill, simple.Border{Left: true, Right: true}, simple.Alignment{}),
			coreCell(w[1], 6, c.capitalStr, fill, simple.Border{Left: true, Right: true}, simple.Alignment{}),
			coreCell(w[2], 6, c.areaStr, fill, simple.Border{Left: true, Right: true}, simple.Alignment{Horiz: simple.Right}),
			coreCell(w[3], 6, c.popStr, fill, simple.Border{Left: true, Right: true}, simple.Alignment{Horiz: simple.Right}),
			simple.LineBreak{Ordinate: -1})
		fill = !fill
	}
	items = append(items, coreCell(wSum, 0, "", false, simple.Border{Top: true}, simple.Alignment{}))
	return items
}

type countryType struct {
	nameStr, capitalStr, areaStr, popStr string
}

func loadData(fileStr string) ([]countryType, error) {
	data, err := ioutil.ReadFile(fileStr)
	if err != nil {
		return nil, err
	}
	rows, err := csv.NewReader(bytes.NewReader(data)).ReadAll()
	if err != nil {
		return nil, err
	}
	list := []countryType{}
	for _, row := range rows {
		c := countryType{
			nameStr:    row[0],
			capitalStr: row[1],
			areaStr:    row[2],
			popStr:     row[3],
		}
		list = append(list, c)
	}
	return list, nil
}

func strDelimit(record string, delim string, index int) string {
	list := strings.Split(record, delim)
	return list[index]
}
