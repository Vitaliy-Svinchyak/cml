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
	Row      int
	Col      int
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

func (b *Block) SetRow(row int) {
	b.Row = row
}

func (b *Block) SetCol(row int) {
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
	if b.Parent != nil {
		return b.Parent.CalcCol() + b.Col
	}

	return b.Col
}

func (b Block) CalcRow() int {
	if b.Parent != nil {
		return b.Parent.CalcRow() + b.Row
	}

	return b.Row
}
