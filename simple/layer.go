package simple

type Layer struct {
	Name         string
	Visible      bool
	LayeredItems []Drawable
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
