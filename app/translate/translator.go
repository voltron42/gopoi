package translate

import (
	"../model"

	"github.com/jung-kurt/gofpdf"
)

type Translator struct {
}

func New() *Translator {
	return &Translator{}
}

func (t *Translator) Translate(doc model.Document, outfile string) error {
	pdf := gofpdf.New(doc.Orientation.String(), doc.Unit.String(), doc.Size.String(), doc.FontDirectory)

	if pdf.Err() {
		return pdf.Error()
	}

	return pdf.OutputFileAndClose(outfile)
}
