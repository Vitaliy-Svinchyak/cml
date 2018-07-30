package Parser

import (
	"cml/BlockTypes"
	"strings"
	"strconv"
	"reflect"
	//"fmt"
)

func ParseString(cml string) []*BlockTypes.Block {
	var cmlTree []*BlockTypes.Block
	var knownTypes = []string{"block"}

	var rows = strings.Split(string(cml), "\n")
	var lastNesting = 0
	var lastElement *BlockTypes.Block

	for rowNumber, row := range rows {
		if row == "" {
			continue
		}
		var nesting, row = detectNesting(row)
		var rowParameters = strings.Split(row, ",")
		var typeAndFirstParam = strings.Split(rowParameters[0], " ")
		rowParameters[0] = typeAndFirstParam[1]
		var blockType = typeAndFirstParam[0]
		var knownType, _ = inArray(blockType, knownTypes)
		if !knownType {
			panic("Unknown block type: " + blockType + " on line " + strconv.Itoa(rowNumber))
		}
		var block = parseProperties(rowParameters, rowNumber)

		if nesting > lastNesting {
			block.Parent = lastElement
			lastElement.Children = append(lastElement.Children, block)
		} else if nesting == lastNesting {
			lastNesting = nesting
			cmlTree = append(cmlTree, block)
			lastElement = cmlTree[len(cmlTree)-1]
		} else {
			lastNesting = nesting
			cmlTree = append(cmlTree, block)
		}
	}

	return cmlTree
}

func parseProperties(properties []string, rowNumber int) *BlockTypes.Block {
	var block = &BlockTypes.Block{}

	for _, property := range properties {
		var propertySplitted = strings.Split(property, ":")
		if len(propertySplitted) == 2 {
			var propertyName = propertySplitted[0]
			var propertyValue = propertySplitted[1]

			switch propertyName {
			case "width":
				block.SetWidth(propertyValue)
				break
			case "height":
				block.SetHeight(propertyValue)
				break
			case "id":
				block.SetId(propertyValue)
				break
			case "row":
				row, _ := strconv.Atoi(propertyValue)
				block.SetRow(row)
				break
			case "col":
				col, _ := strconv.Atoi(propertyValue)
				block.SetCol(col)
				break
			case "text":
				block.SetText(string([]rune(propertyValue)[1:len(propertyValue) - 1]))
				break
			case "border":
				border, _ := strconv.Atoi(propertyValue)
				block.SetBorder(border)
				break
			default:
				panic("Unknown property " + propertyName + " on line " + strconv.Itoa(rowNumber))
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

func inArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}
