package main

import (
	"../"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"path/filepath"
)

func main() {
	const (
		margin = 10
		wd     = 210
		ht     = 297
	)
	fileList := []string{
		"logo-gray.png",
		"logo.jpg",
		"logo.png",
		"logo-rgb.png",
		"logo-progressive.jpg",
	}
	var infoPtr *gofpdf.ImageInfoType
	var imageFileStr string
	var imgWd, imgHt, lf, tp float64
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetMargins(10, 10, 10)
	pdf.SetFont("Helvetica", "", 15)
	for j, str := range fileList {
		imageFileStr = example.ImageFile(str)
		infoPtr = pdf.RegisterImage(imageFileStr, "")
		imgWd, imgHt = infoPtr.Extent()
		switch j {
		case 0:
			lf = margin
			tp = margin
		case 1:
			lf = wd - margin - imgWd
			tp = margin
		case 2:
			lf = (wd - imgWd) / 2.0
			tp = (ht - imgHt) / 2.0
		case 3:
			lf = margin
			tp = ht - imgHt - margin
		case 4:
			lf = wd - imgWd - margin
			tp = ht - imgHt - margin
		}
		pdf.Image(imageFileStr, lf, tp, imgWd, imgHt, false, "", 0, "")
	}
	fileStr := example.Filename("Fpdf_RegisterImage")
	// Test the image information retrieval method
	infoShow := func(imageStr string) {
		imageStr = example.ImageFile(imageStr)
		info := pdf.GetImageInfo(imageStr)
		if info != nil {
			if info.Width() > 0.0 {
				fmt.Printf("Image %s is registered\n", filepath.ToSlash(imageStr))
			} else {
				fmt.Printf("Incorrect information for image %s\n", filepath.ToSlash(imageStr))
			}
		} else {
			fmt.Printf("Image %s is not registered\n", filepath.ToSlash(imageStr))
		}
	}
	infoShow(fileList[0])
	infoShow("foo.png")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
}
