package Parser

import (
	"cml/BlockTypes"
	"strings"
	"strconv"
)

func ParseString(cml string) []BlockTypes.BlockProperties {
	var cmlTree []BlockTypes.BlockProperties

	var rows = strings.Split(string(cml), "\n")
	var lastNesting = 0
	var lastElement *BlockTypes.BlockProperties

	for _, row := range rows {
		var nesting, row = detectNesting(row)
		var rowParameters = strings.Split(row, " ")
		var block = parseProperties(rowParameters)

		if nesting > lastNesting {
			block.Parent = lastElement
			lastElement.Children = append(lastElement.Children, block)
		} else if nesting == lastNesting {
			lastNesting = nesting
			cmlTree = append(cmlTree, block)
			lastElement = &cmlTree[len(cmlTree)-1]
		} else {
			lastNesting = nesting
			cmlTree = append(cmlTree, block)
		}
	}

	return cmlTree
}

func parseProperties(properties []string) BlockTypes.BlockProperties {
	var block = BlockTypes.BlockProperties{Children: []BlockTypes.BlockProperties{}}

	for _, property := range properties {
		var propertySplitted = strings.Split(property, ":")
		if len(propertySplitted) == 2 {
			var propertyValue = propertySplitted[1]
			switch propertySplitted[0] {
			case "width":
				block.Width = propertyValue
				break
			case "height":
				block.Height = propertyValue
				break
			case "row":
				i, _ := strconv.Atoi(propertyValue)
				block.Row = i
				break
			case "col":
				i, _ := strconv.Atoi(propertyValue)
				block.Col = i
				break
			}
		}
	}

	return block
}

func detectNesting(row string) (int, string) {
	var length = 0
	var rowRunes = []rune(row)

	for i := 0; i < len(row); i++ {
		if rowRunes[i] == 32 {
			length++
		} else {
			break
		}
	}

	return length / 4, string(rowRunes[length:len(row)])
}
