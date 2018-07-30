package main

import (
	"io/ioutil"
	"fmt"
	"cml/Parser"
	//"github.com/k0kubun/pp"
	"cml/Renderer"
)

func main() {
	cml, err := ioutil.ReadFile("ex1.cml") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	var cmlTree = Parser.ParseString(string(cml))
	Renderer.PaintGui(cmlTree)
	//pp.Println(cmlTree)
}
