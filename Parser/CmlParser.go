package Parser

import (
	"cml/BlockTypes"
	"strings"
	"strconv"
	"reflect"
	"github.com/jroimartin/gocui"
)

func ParseString(cml string) ([]*BlockTypes.Block, []*BlockTypes.Block) {
	var blocksById = make(map[string]*BlockTypes.Block)
	var lastBlockOnNesting []*BlockTypes.Block
	var cmlTree []*BlockTypes.Block
	var cmlSlice []*BlockTypes.Block
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
			lastElement = block
		} else if nesting == 0 {
			cmlTree = append(cmlTree, block)
			lastElement = cmlTree[len(cmlTree)-1]
		} else if nesting < lastNesting {
			var parent = lastBlockOnNesting[nesting-1]
			block.Parent = parent
			parent.Children = append(parent.Children, block)
			lastElement = block
		} else if nesting == lastNesting {
			var parent = lastBlockOnNesting[nesting-1]
			block.Parent = parent
			parent.Children = append(parent.Children, block)
			lastElement = block
		} else {
			cmlTree = append(cmlTree, block)
		}

		lastNesting = nesting
		if len(lastBlockOnNesting) < nesting+1 {
			lastBlockOnNesting = append(lastBlockOnNesting, block)
		} else {
			lastBlockOnNesting[nesting] = block
		}
		cmlSlice = append(cmlSlice, block)

		if block.Id != "" {
			if blocksById[block.Id] != nil {
				panic("Block with Id:" + block.Id + " already exists." + "Duplicated id found on line " + strconv.Itoa(rowNumber))
			}
			blocksById[block.Id] = block
		}
	}

	return cmlSlice, cmlTree
}

func getRowParameters(row string) []string {
	var params = strings.Split(row, " ")
	var formattedParams []string
	var concattedText string

	for _, param := range params {
		var propertySplitted = strings.Split(param, ":")
		if propertySplitted[0] == "text" {
			var lastSymbol = propertySplitted[1][len(propertySplitted[1])-1]
			if (lastSymbol != 34 && len(propertySplitted[1]) > 1) || (lastSymbol == 34 && len(propertySplitted[1]) == 1) {
				concattedText = removeEscaping(param) + " "
				continue
			}
		}

		if len(concattedText) == 0 {
			formattedParams = append(formattedParams, param)
		} else {
			var penultimateSymbol uint8
			if len(param) == 0 || len(param) == 1 {
				param = " "
				penultimateSymbol = 32
			} else {
				penultimateSymbol = param[len(param)-2]
			}
			var lastSymbol = param[len(param)-1]

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
				// allow % col
				row, _ := strconv.Atoi(propertyValue)
				block.SetRow(row)
				break
			case "col":
				// allow % col
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
	default:
		panic("Unknown color " + color)
	}

	return 0
}

func removeEscaping(str string) string {
	return strings.Replace(str, `\"`, `"`, -1)
}
