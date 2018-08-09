package BlockTypes

import (
	"github.com/jroimartin/gocui"
	"fmt"
)

type Block struct {
	AbstractSheet

	Text    string
	Border  int
	BgColor gocui.Attribute
	FgColor gocui.Attribute
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

func (b *Block) InitializeDefaultParams() {
	if b.Col == "" {
		b.SetCol("0")
	}

	if b.Row == "" {
		b.SetRow("0")
	}

	if b.Id == "" {
		b.SetId(RandStringRunes(32))
	}
}

func (b Block) GetMetaInfo() map[string]int {
	return map[string]int{}
}
