package BlockTypes

import (
	"testing"
)

func TestCalcWidth(t *testing.T) {
	var block = Block{}
	block.SetWidth("20pt")
	block.SetCol("20")

	var w = block.CalcWidth()
	if w != 40 {
		t.Error("Failed: Expected 40, got ", w)
	}
}

func TestCalcWidthWithParent(t *testing.T) {
	var parentBlock = Block{}
	parentBlock.SetWidth("40pt")
	parentBlock.SetCol("20")

	var block = Block{}
	block.SetWidth("50%")
	block.SetCol("0")
	block.Parent = &parentBlock

	var w = block.CalcWidth()
	if w != 40 {
		t.Error("Failed: Expected 40, got ", w)
	}
}
