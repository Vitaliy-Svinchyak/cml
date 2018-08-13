package Parser

import (
	"cml/BlockTypes"
	"cml/Renderer"
)

type Document struct {
	blocksById map[string]BlockTypes.Sheet
}

func NewDocument(cml string) Document {
	var document = Document{make(map[string]BlockTypes.Sheet)}

	var cmlSlice, _ = ParseString(cml, &document)
	Renderer.PaintGui(cmlSlice)

	return document
}

func (b *Document) saveId(block BlockTypes.Sheet) bool {
	if b.blocksById[block.GetId()] != nil {
		return false
	}

	b.blocksById[block.GetId()] = block

	return true
}

func (b *Document) GetElementById(id string) BlockTypes.Sheet {
	if b.blocksById[id] != nil {
		return nil
	}

	return b.blocksById[id]
}
