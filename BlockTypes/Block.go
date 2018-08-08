package BlockTypes

import (
	"github.com/jroimartin/gocui"
	"fmt"
	"strconv"
	"cml/Config"
)

type Styled interface {
	SetBorder(border int)
	SetBgColor(color gocui.Attribute)
	SetFgColor(color gocui.Attribute)
}

type Located interface {
	SetWidth(width string)
	SetRow(row int)
	SetCol(col int)
}

type Identifying interface {
	SetId(id string)
}

type Block struct {
	Id       string
	Text     string
	Width    string
	Height   string
	Row      string
	Col      string
	Border   int
	BgColor  gocui.Attribute
	FgColor  gocui.Attribute
	Children []*Block
	Parent   *Block
}

func (b *Block) SetWidth(width string) {
	var matched = numberOrPercent(width)

	if !matched {
		panic("Width " + width + " is not valid")
	}

	b.Width = width
}

func (b *Block) SetHeight(height string) {
	var matched = numberOrPercent(height)

	if !matched {
		panic("Height " + height + " is not valid")
	}

	b.Height = height
}

func (b *Block) SetRow(row string) {
	b.Row = row
}

func (b *Block) SetCol(row string) {
	b.Col = row
}

func (b *Block) SetId(id string) {
	b.Id = id
}

func (b *Block) SetText(text string) {
	b.Text = text
}

func (b *Block) SetBorder(border int) {
	b.Border = border
}

func (b *Block) SetBgColor(color gocui.Attribute) {
	b.BgColor = color
}
func (b *Block) SetFgColor(color gocui.Attribute) {
	b.FgColor = color
}

func (b Block) Layout(g *gocui.Gui) error {
	v, err := g.SetView(b.Id, b.CalcCol(), b.CalcRow(), b.CalcWidth(), b.CalcHeight())
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		if b.Border == 1 {
			v.Frame = true
		} else {
			v.Frame = false
		}

		if b.BgColor != 0 {
			v.BgColor = b.BgColor
		}
		if b.FgColor != 0 {
			v.FgColor = b.FgColor
		}

		fmt.Fprint(v, b.Text)
	}
	return nil
}

func (b Block) CalcWidth() int {
	if b.Width[len(b.Width)-1] == 37 {
		var widthPercent, _ = strconv.Atoi(b.Width[0:len(b.Width)-1])
		var availableWidth int

		if b.Parent != nil {
			var parentWidth = b.Parent.CalcWidth()
			availableWidth = parentWidth - b.Parent.CalcCol()
		} else {
			availableWidth, _ = Config.GetTerminalSize()
		}
		var newWidth = availableWidth * widthPercent / 100

		return b.CalcCol() + newWidth
	}

	var width, _ = strconv.Atoi(b.Width[0:len(b.Width)-2])

	return b.CalcCol() + width
}

func (b Block) CalcHeight() int {
	if b.Height[len(b.Height)-1] == 37 {
		var heightPercent, _ = strconv.Atoi(b.Height[0:len(b.Height)-1])
		var availableHeight int

		if b.Parent != nil {
			var parentHeight = b.Parent.CalcHeight()
			availableHeight = parentHeight - b.Parent.CalcRow()
		} else {
			availableHeight, _ = Config.GetTerminalSize()
		}
		var newHeight = availableHeight * heightPercent / 100

		return b.CalcRow() + newHeight
	}

	var height, _ = strconv.Atoi(b.Height[0:len(b.Height)-2])
	return b.CalcRow() + height
}

func (b Block) CalcCol() int {
	var calculatedCol int
	var currentCol = b.Col

	if currentCol[len(currentCol)-1] == 37 {
		var colPercent, _ = strconv.Atoi(b.Col[0:len(b.Col)-1])
		var parentWidth = b.Parent.CalcWidth() - b.Parent.CalcCol()
		calculatedCol = parentWidth * colPercent / 100
	} else {
		calculatedCol, _ = strconv.Atoi(currentCol)
	}

	if b.Parent != nil {
		return b.Parent.CalcCol() + calculatedCol
	}

	return calculatedCol
}

func (b Block) CalcRow() int {
	var calculatedRow int
	var currentRow = b.Row

	if currentRow[len(currentRow)-1] == 37 {
		var rowPercent, _ = strconv.Atoi(b.Col[0:len(b.Col)-1])
		var parentWidth = b.Parent.CalcHeight() - b.Parent.CalcRow()
		calculatedRow = parentWidth * rowPercent / 100
	} else {
		calculatedRow, _ = strconv.Atoi(currentRow)
	}

	if b.Parent != nil {
		return b.Parent.CalcRow() + calculatedRow
	}

	return calculatedRow
}
