package main

import (
	"io/ioutil"
	"fmt"
	"time"
	"math/rand"
	"cml/Parser"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	cml, err := ioutil.ReadFile("examples/ex2.cml")
	if err != nil {
		fmt.Print(err)
	}

	Parser.NewDocument(string(cml))
	//var cmlSlice, cmlTree = Parser.ParseString(string(cml))
	//	pp.Println(cmlTree)
}
