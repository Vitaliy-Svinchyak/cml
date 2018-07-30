package Renderer

import (
	"github.com/jroimartin/gocui"
	"log"
	"cml/BlockTypes"
)

func PaintGui(cmlTree []*BlockTypes.Block) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		panic(err)
	}
	defer g.Close()

	var managers []gocui.Manager
	for _, block := range cmlTree {
		managers = append(managers, block)
	}

	g.SetManager(managers...)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
