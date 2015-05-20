package model

type Document struct {
	Orientation   Orientation `json:"orientation,attr"`
	Unit          Unit        `json:"unit,attr"`
	Size          Size        `json:"size,attr"`
	FontDirectory string      `json:"font-directory,attr"`

	Metadata *Metadata `json:"meta"`
	Header   Page      `json:"header>page"`
	Footer   Page      `json:"footer>page"`
	Pages    []Page    `json:"page"`
}

type Metadata struct {
}

type Orientation int

func (o *Orientation) String() string {
	return ""
}

type Unit int

func (u *Unit) String() string {
	return ""
}

type Size int

func (s *Size) String() string {
	return ""
}

type Page struct {
}
