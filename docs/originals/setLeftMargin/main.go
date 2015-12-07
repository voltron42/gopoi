package main

import (
	"../"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
)

func main() {
	var y0 float64
	var crrntCol int
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetDisplayMode("fullpage", "TwoColumnLeft")
	titleStr := "20000 Leagues Under the Seas"
	pdf.SetTitle(titleStr, false)
	pdf.SetAuthor("Jules Verne", false)
	setCol := func(col int) {
		// Set position at a given column
		crrntCol = col
		x := 10.0 + float64(col)*65.0
		pdf.SetLeftMargin(x)
		pdf.SetX(x)
	}
	chapterTitle := func(chapNum int, titleStr string) {
		// Arial 12
		pdf.SetFont("Arial", "", 12)
		// Background color
		pdf.SetFillColor(200, 220, 255)
		// Title
		pdf.CellFormat(0, 6, fmt.Sprintf("Chapter %d : %s", chapNum, titleStr),
			"", 1, "L", true, 0, "")
		// Line break
		pdf.Ln(4)
		y0 = pdf.GetY()
	}
	chapterBody := func(fileStr string) {
		// Read text file
		txtStr, err := ioutil.ReadFile(fileStr)
		if err != nil {
			pdf.SetError(err)
		}
		// Font
		pdf.SetFont("Times", "", 12)
		// Output text in a 6 cm width column
		pdf.MultiCell(60, 5, string(txtStr), "", "", false)
		pdf.Ln(-1)
		// Mention
		pdf.SetFont("", "I", 0)
		pdf.Cell(0, 5, "(end of excerpt)")
		// Go back to first column
		setCol(0)
	}
	printChapter := func(num int, titleStr, fileStr string) {
		// Add chapter
		pdf.AddPage()
		chapterTitle(num, titleStr)
		chapterBody(fileStr)
	}
	pdf.SetAcceptPageBreakFunc(func() bool {
		// Method accepting or not automatic page break
		if crrntCol < 2 {
			// Go to next column
			setCol(crrntCol + 1)
			// Set ordinate to top
			pdf.SetY(y0)
			// Keep on page
			return false
		}
		// Go back to first column
		setCol(0)
		// Page break
		return true
	})
	pdf.SetHeaderFunc(func() {
		// Arial bold 15
		pdf.SetFont("Arial", "B", 15)
		// Calculate width of title and position
		wd := pdf.GetStringWidth(titleStr) + 6
		pdf.SetX((210 - wd) / 2)
		// Colors of frame, background and text
		pdf.SetDrawColor(0, 80, 180)
		pdf.SetFillColor(230, 230, 0)
		pdf.SetTextColor(220, 50, 50)
		// Thickness of frame (1 mm)
		pdf.SetLineWidth(1)
		// Title
		pdf.CellFormat(wd, 9, titleStr, "1", 1, "C", true, 0, "")
		// Line break
		pdf.Ln(10)
		// Save ordinate
		y0 = pdf.GetY()
	})
	pdf.SetFooterFunc(func() {
		// Position at 1.5 cm from bottom
		pdf.SetY(-15)
		// Arial italic 8
		pdf.SetFont("Arial", "I", 8)
		// Text color in gray
		pdf.SetTextColor(128, 128, 128)
		// Page number
		pdf.CellFormat(0, 10, fmt.Sprintf("Page %d", pdf.PageNo()),
			"", 0, "C", false, 0, "")
	})
	printChapter(1, "A RUNAWAY REEF", example.TextFile("20k_c1.txt"))
	printChapter(2, "THE PROS AND CONS", example.TextFile("20k_c2.txt"))
	fileStr := example.Filename("Fpdf_SetLeftMargin_multicolumn")
	err := pdf.OutputFileAndClose(fileStr)
	example.Summary(err, fileStr)
}
