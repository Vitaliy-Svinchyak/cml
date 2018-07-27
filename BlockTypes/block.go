package BlockTypes

import (
	"regexp"
)

type BlockProperties struct {
	Id          string
	Text        string
	Width       string
	Height      string
	Row         int
	Col         int
	Border      int
	BorderColor string
	Children    []BaseBlockI
	Parent      BaseBlockI
}

type BaseBlockI interface {
	SetWidth(width string)
}

type Block struct {
	BlockProperties
}

func (b Block) SetWidth(width string) {
	matched, err := regexp.MatchString("[0-9]+[pt%]", width)

	if err != nil {
		panic(err)
	}
	if !matched {
		panic("Width " + width + " is not valid")
	}

	b.Width = width
	// not working
}
