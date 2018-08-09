package BlockTypes

import (
	"strconv"
	"cml/Config"
	"strings"
)

type AbstractSheet struct {
	Id       string
	Children []Sheet
	Parent   Sheet
	Width    string
	Height   string
	Row      string
	Col      string
}

func (b *AbstractSheet) SetId(id string) {
	b.Id = id
}

func (b AbstractSheet) GetId() string {
	return b.Id
}

func (b *AbstractSheet) SetParent(parent Sheet) {
	b.Parent = parent
}

func (b *AbstractSheet) GetParent() Sheet {
	return b.Parent
}

func (b *AbstractSheet) AddChild(child Sheet) {
	b.Children = append(b.Children, child)
}

func (b *AbstractSheet) GetChildren() []Sheet {
	return b.Children
}

func (b *Block) SetWidth(width string) {
	var matched = ptOrPercent(width)

	if !matched {
		panic("Width " + width + " is not valid")
	}

	b.Width = width
}

func (b *Block) SetHeight(height string) {
	var matched = ptOrPercent(height)

	if !matched {
		panic("Height " + height + " is not valid")
	}

	b.Height = height
}

func (b *Block) SetRow(row string) {
	var matched = numberOrPercent(row)

	if !matched {
		panic("Row " + row + " is not valid")
	}

	b.Row = row
}

func (b *Block) SetCol(col string) {
	var matched = numberOrPercent(col)

	if !matched {
		panic("Col " + col + " is not valid")
	}

	b.Col = col
}

func (b AbstractSheet) CalcWidth() int {
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
	} else if strings.Contains(b.Width, "kt") {
		var metaInfo = b.Parent.GetMetaInfo()
		var ktSize = metaInfo["ktSizeWidth"]
		var widthNumber, _ = strconv.Atoi(b.Width[0:len(b.Width)-2])

		return b.CalcCol() + widthNumber*ktSize
	}

	var width, _ = strconv.Atoi(b.Width[0:len(b.Width)-2])

	return b.CalcCol() + width
}

func (b AbstractSheet) CalcHeight() int {
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
	} else if strings.Contains(b.Height, "kt") {
		var metaInfo = b.Parent.GetMetaInfo()
		var ktSize = metaInfo["ktSizeHeight"]
		var heightNumber, _ = strconv.Atoi(b.Height[0:len(b.Height)-2])

		return b.CalcRow() + heightNumber*ktSize
	}

	var height, _ = strconv.Atoi(b.Height[0:len(b.Height)-2])
	return b.CalcRow() + height
}

func (b AbstractSheet) CalcCol() int {
	var calculatedCol int
	var currentCol = b.Col

	if currentCol[len(currentCol)-1] == 37 {
		var colPercent, _ = strconv.Atoi(b.Col[0:len(b.Col)-1])
		var parentWidth = b.Parent.CalcWidth() - b.Parent.CalcCol()
		calculatedCol = parentWidth * colPercent / 100
	} else if strings.Contains(currentCol, "kt") {
		var metaInfo = b.Parent.GetMetaInfo()
		var ktSize = metaInfo["ktSizeWidth"]
		var colNumber, _ = strconv.Atoi(currentCol[0:len(currentCol)-2])

		calculatedCol = colNumber * ktSize
	} else {
		calculatedCol, _ = strconv.Atoi(currentCol)
	}

	if b.Parent != nil {
		return b.Parent.CalcCol() + calculatedCol
	}

	return calculatedCol
}

func (b AbstractSheet) CalcRow() int {
	var calculatedRow int
	var currentRow = b.Row

	if currentRow[len(currentRow)-1] == 37 {
		var rowPercent, _ = strconv.Atoi(b.Col[0:len(b.Col)-1])
		var parentWidth = b.Parent.CalcHeight() - b.Parent.CalcRow()
		calculatedRow = parentWidth * rowPercent / 100
	} else if strings.Contains(currentRow, "kt") {
		var metaInfo = b.Parent.GetMetaInfo()
		var ktSize = metaInfo["ktSizeHeight"]
		var rowNumber, _ = strconv.Atoi(currentRow[0:len(currentRow)-2])

		calculatedRow = rowNumber * ktSize
	} else {
		calculatedRow, _ = strconv.Atoi(currentRow)
	}

	if b.Parent != nil {
		return b.Parent.CalcRow() + calculatedRow
	}

	return calculatedRow
}
