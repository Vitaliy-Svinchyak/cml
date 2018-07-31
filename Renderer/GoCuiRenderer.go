package Renderer

import (
	"github.com/jroimartin/gocui"
	"cml/BlockTypes"
	"cml/Config"
)

func PaintGui(cmlTree []*BlockTypes.Block) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		panic(err)
	}
	defer g.Close()

	Config.SetMaxXAndMaxY(g.Size())

	var managers []gocui.Manager
	//todo fix for any deep
	for _, block := range cmlTree {
		managers = append(managers, block)
		if len(block.Children) != 0 {
			for _, child := range block.Children {
				managers = append(managers, child)
			}
		}
	}

	g.SetManager(managers...)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		panic(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		panic(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
