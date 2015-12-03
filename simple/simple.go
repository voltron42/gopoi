package simple

import (
	"encoding/xml"
	"github.com/jung-kurt/gofpdf"
)

type Document struct {
	Format Format   `xml:"format"`
	Items  ItemList `xml:"body"`
}

type ItemList []interface{}

func (i *ItemList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// TODO --
	return nil
}

type Format struct {
	Orientation   *Orientation
	Unit          *Unit
	Size          *Size
	FontDirectory string
	OutputFile    string
}

type context struct {
	pdf     *gofpdf.Fpdf
	outfile string
}

func NewContext(format Format) *context {
	return &context{pdf: gofpdf.New(format.Orientation.String(), format.Unit.String(), format.Size.String(), format.FontDirectory), outfile: format.OutputFile}
}

type Drawable interface {
	Draw(ctx *context) error
}

func (c *context) Publish(items ...Drawable) error {
	for _, item := range items {
		err := item.Draw(c)
		if err != nil {
			return err
		}
	}
	return c.pdf.OutputFileAndClose(c.outfile)
}

func main() {
	pdf := gofpdf.New("", "", "", "")
	pdf.ClosePath()
}

/*
func (f *Fpdf) SetDisplayMode(zoomStr, layoutStr string)
func (f *Fpdf) AliasNbPages(aliasStr string)
func (f *Fpdf) ClearError()
func (f *Fpdf) Err() bool
func (f *Fpdf) Error() error
func (f *Fpdf) GetFontDesc(familyStr, styleStr string) FontDescType
func (f *Fpdf) GetImageInfo(imageStr string) (info *ImageInfoType)
func (f *Fpdf) GetStringWidth(s string) float64
func (f *Fpdf) ImageTypeFromMime(mimeStr string) (tp string)
func (f *Fpdf) OpenLayerPane()
func (f *Fpdf) Output(w io.Writer) error
func (f *Fpdf) OutputAndClose(w io.WriteCloser) error
func (f *Fpdf) OutputFileAndClose(fileStr string) error
func (f *Fpdf) PageNo() int
func (f *Fpdf) PageSize(pageNum int) (wd, ht float64, unitStr string)
func (f *Fpdf) PointConvert(pt float64) (u float64)
func (f *Fpdf) PointToUnitConvert(pt float64) (u float64)
func (f *Fpdf) RegisterImage(fileStr, tp string) (info *ImageInfoType)
func (f *Fpdf) RegisterImageReader(imgName, tp string, r io.Reader) (info *ImageInfoType)
func (f *Fpdf) SetAcceptPageBreakFunc(fnc func() bool)
func (f *Fpdf) SetAuthor(authorStr string, isUTF8 bool)
func (f *Fpdf) SetAutoPageBreak(auto bool, margin float64)
func (f *Fpdf) SetCatalogSort(flag bool)
func (f *Fpdf) SetCompression(compress bool)
func (f *Fpdf) SetCreationDate(tm time.Time)
func (f *Fpdf) SetCreator(creatorStr string, isUTF8 bool)
func (f *Fpdf) SetFontLoader(loader FontLoader)
func (f *Fpdf) SetFontLocation(fontDirStr string)
func (f *Fpdf) SetKeywords(keywordsStr string, isUTF8 bool)
func (f *Fpdf) SetProtection(actionFlag byte, userPassStr, ownerPassStr string)
func (f *Fpdf) SetSubject(subjectStr string, isUTF8 bool)
func (f *Fpdf) SetFooterFunc(fnc func())
func (f *Fpdf) SetHeaderFunc(fnc func())
func (f *Fpdf) SetTitle(titleStr string, isUTF8 bool)
func (f *Fpdf) String() string
func (f *Fpdf) SplitLines(txt []byte, w float64) [][]byte
func (f *Fpdf) UnicodeTranslatorFromDescriptor(cpStr string) (rep func(string) string)
func (f *Fpdf) UnitToPointConvert(u float64) (pt float64)
*/