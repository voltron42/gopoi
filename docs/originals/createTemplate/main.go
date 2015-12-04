package main

import (
	"../"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetCompression(false)
	// pdf.SetFont("Times", "", 12)
	template := pdf.CreateTemplate(func(tpl *gofpdf.Tpl) {
		tpl.Image(example.ImageFile("logo.png"), 6, 6, 30, 0, false, "", 0, "")
		tpl.SetFont("Arial", "B", 16)
		tpl.Text(40, 20, "Template says hello")
		tpl.SetDrawColor(0, 100, 200)
		tpl.SetLineWidth(2.5)
		tpl.Line(95, 12, 105, 22)
	})
	_, tplSize := template.Size()
	// fmt.Println("Size:", tplSize)
	// fmt.Println("Scaled:", tplSize.ScaleBy(1.5))

	template2 := pdf.CreateTemplate(func(tpl *gofpdf.Tpl) {
		tpl.UseTemplate(template)
		subtemplate := tpl.CreateTemplate(func(tpl2 *gofpdf.Tpl) {
			tpl2.Image(example.ImageFile("logo.png"), 6, 86, 30, 0, false, "", 0, "")
			tpl2.SetFont("Arial", "B", 16)
			tpl2.Text(40, 100, "Subtemplate says hello")
			tpl2.SetDrawColor(0, 200, 100)
			tpl2.SetLineWidth(2.5)
			tpl2.Line(102, 92, 112, 102)
		})
		tpl.UseTemplate(subtemplate)
	})

	pdf.SetDrawColor(200, 100, 0)
	pdf.SetLineWidth(2.5)
	pdf.SetFont("Arial", "B", 16)

	pdf.AddPage()
	pdf.UseTemplate(template)
	pdf.UseTemplateScaled(template, gofpdf.PointType{X: 0, Y: 30}, tplSize)
	pdf.UseTemplateScaled(template, gofpdf.PointType{X: 0, Y: 60}, tplSize.ScaleBy(1.4))
	pdf.Line(40, 210, 60, 210)
	pdf.Text(40, 200, "Template example page 1")

	pdf.AddPage()
	pdf.UseTemplate(template2)
	pdf.Line(60, 210, 80, 210)
	pdf.Text(40, 200, "Template example page 2")

	fileStr := example.Filename("Fpdf_CreateTemplate")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
}
