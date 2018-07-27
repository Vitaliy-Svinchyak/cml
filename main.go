package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"cml/BlockTypes"
	"strconv"
)

func main() {
	var cmlTree []BlockTypes.BlockProperties
	b, err := ioutil.ReadFile("ex1.cml") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	var rows = strings.Split(string(b), "\n")
	var lastNesting = 0
	var lastElement *BlockTypes.BlockProperties

	for _, row := range rows {
		var nesting, row = detectNesting(row)
		var rowParameters = strings.Split(row, " ")
		var block = parseProperties(rowParameters)

		if nesting > lastNesting {
			lastElement.Children = append(lastElement.Children, block)
			block.Parent = lastElement
		} else if nesting == lastNesting {
			lastNesting = nesting
			cmlTree = append(cmlTree, block)
			lastElement = &cmlTree[len(cmlTree)-1]
		} else {
			lastNesting = nesting
			cmlTree = append(cmlTree, block)
		}
	}

	fmt.Println(cmlTree)
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
