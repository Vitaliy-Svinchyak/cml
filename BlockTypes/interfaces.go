package BlockTypes

import "github.com/jroimartin/gocui"

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

type Sheet interface {
	SetId(id string)
	GetId() string

	SetParent(Sheet)
	GetParent() Sheet
	AddChild(Sheet)
	GetChildren() []Sheet

	CalcWidth() int
	CalcCol() int
	CalcHeight() int
	CalcRow() int

	InitializeDefaultParams()
	GetMetaInfo() map[string]int

	gocui.Manager
}

