package BlockTypes

import (
	"github.com/jroimartin/gocui"
)

type Grid struct {
	AbstractSheet

	Rows   int
	Cols   int
}

func (b *Grid) SetRows(rows int) {
	b.Rows = rows
}

func (b *Grid) SetCols(cols int) {
	b.Cols = cols
}

func (b *Grid) SetRow(row string) {
	var matched = nmPrKt(row)

	if !matched {
		panic("Row " + row + " is not valid")
	}

	b.Row = row
}

func (b *Grid) SetCol(col string) {
	var matched = nmPrKt(col)

	if !matched {
		panic("Col " + col + " is not valid")
	}

	b.Col = col
}

func (b *Grid) SetWidth(width string) {
	var matched = ptPrKt(width)

	if !matched {
		panic("Width " + width + " is not valid")
	}

	b.Width = width
}

func (b *Grid) SetHeight(height string) {
	var matched = ptPrKt(height)

	if !matched {
		panic("Height " + height + " is not valid")
	}

	b.Height = height
}

func (b Grid) Layout(g *gocui.Gui) error {
	return nil
}

func (b *Grid) InitializeDefaultParams() {
	if b.Col == "" {
		b.Col = "0"
	}

	if b.Row == "" {
		b.Row = "0"
	}

	if b.Id == "" {
		b.SetId(RandStringRunes(32))
	}
	if b.Width == "" {
		b.SetWidth("1pt")
	}

	if b.Height == "" {
		b.SetHeight("1pt")
	}

}

func (b Grid) GetMetaInfo() map[string]int {
	return map[string]int{
		"ktSizeHeight": (b.CalcHeight() - b.CalcRow()) / b.Rows,
		"ktSizeWidth":  (b.CalcWidth() - b.CalcCol()) / b.Cols,
	}
}
