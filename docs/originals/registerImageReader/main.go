package main

import (
	"../"
	"github.com/jung-kurt/gofpdf"
	"net/http"
)

func main() {
	const (
		margin   = 10
		wd       = 210
		ht       = 297
		fontSize = 15
		urlStr   = "https://github.com/jung-kurt/gofpdf/blob/master/image/gofpdf.png?raw=true"
		msgStr   = `Images from the web can be easily embedded when a PDF document is generated.`
	)

	var (
		rsp *http.Response
		err error
		tp  string
	)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", fontSize)
	ln := pdf.PointConvert(fontSize)
	pdf.MultiCell(wd-margin-margin, ln, msgStr, "", "L", false)
	rsp, err = http.Get(urlStr)
	if err == nil {
		tp = pdf.ImageTypeFromMime(rsp.Header["Content-Type"][0])
		infoPtr := pdf.RegisterImageReader(urlStr, tp, rsp.Body)
		if pdf.Ok() {
			imgWd, imgHt := infoPtr.Extent()
			pdf.Image(urlStr, (wd-imgWd)/2.0, pdf.GetY()+ln,
				imgWd, imgHt, false, tp, 0, "")
		}
	} else {
		pdf.SetError(err)
	}
	fileStr := example.Filename("Fpdf_RegisterImageReader_url")
	err = pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
}
