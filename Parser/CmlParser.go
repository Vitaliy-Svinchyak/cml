package Parser

import (
	"cml/BlockTypes"
	"strings"
	"strconv"
	"reflect"
	"github.com/jroimartin/gocui"
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
		var rowParameters = getRowParameters(row)
		var blockType = rowParameters[0]
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

func getRowParameters(row string) []string {
	var params = strings.Split(row, " ")
	var formattedParams []string
	var concattedText string

	for _, param := range params {
		var propertySplitted = strings.Split(param, ":")
		if propertySplitted[0] == "text" {
			var lastSymbol = propertySplitted[1][len(propertySplitted[1])-1]
			if lastSymbol != 34 {
				concattedText = removeEscaping(param) + " "
				continue
			}
		}

		if len(concattedText) == 0 {
			formattedParams = append(formattedParams, param)
		} else {
			var lastSymbol = param[len(param)-1]
			var penultimateSymbol = param[len(param)-2]

			if lastSymbol == 34 && penultimateSymbol != 92 {
				concattedText += removeEscaping(param)
				formattedParams = append(formattedParams, concattedText)
				concattedText = ""
			} else {
				concattedText += removeEscaping(param) + " "
			}
		}
	}

	return formattedParams
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
				block.SetText(string([]rune(propertyValue)[1:len(propertyValue)-1]))
				break
			case "border":
				border, _ := strconv.Atoi(propertyValue)
				block.SetBorder(border)
				break
			case "bg-color":
				block.SetBgColor(detectColor(propertyValue))
				break
			case "fg-color":
				block.SetFgColor(detectColor(propertyValue))
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

func detectColor(color string) gocui.Attribute {
	switch color {
	case "red":
		return gocui.ColorRed
	case "black":
		return gocui.ColorBlack
	case "green":
		return gocui.ColorGreen
	case "yellow":
		return gocui.ColorYellow
	case "blue":
		return gocui.ColorBlue
	case "magenta":
		return gocui.ColorMagenta
	case "cyan":
		return gocui.ColorCyan
	case "white":
		return gocui.ColorWhite
	}

	return 0
}

func removeEscaping(str string) string {
	return strings.Replace(str, `\"`, `"`, -1)
}
