package simple

import (
	"encoding/xml"
)

type Layer struct {
	XMLName      xml.Name `xml:"layer"`
	Name         string   `xml:"name,attr"`
	Visible      bool     `xml:"visible,attr"`
	LayeredItems ItemList `xml:"items"`
}

func (l Layer) Draw(ctx *context) error {
	layerId := ctx.pdf.AddLayer(l.Name, l.Visible)
	ctx.pdf.BeginLayer(layerId)
	for _, item := range l.LayeredItems {
		err := item.Draw(ctx)
		if err != nil {
			return err
		}
	}
	ctx.pdf.EndLayer()
	return ctx.pdf.Error()
}
