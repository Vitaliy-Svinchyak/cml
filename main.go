package main

import (
	"io/ioutil"
	"fmt"
	"cml/Parser"
	"github.com/k0kubun/pp"
	"os"
	"cml/Renderer"
)

func main() {
	argsWithProg := os.Args
	var debug = false
	if len(argsWithProg) > 1 && argsWithProg[1] == "--debug" {
		debug = true
	}

	cml, err := ioutil.ReadFile("ex1.cml")
	if err != nil {
		fmt.Print(err)
	}

	var cmlTree = Parser.ParseString(string(cml))
	if debug {
		pp.Println(cmlTree)
	} else {
		Renderer.PaintGui(cmlTree)
	}
}
