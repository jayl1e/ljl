package main

import (
	"flag"
	"fmt"
	"jayjaylee.com/ljl/translate"
	"os"
)

var inputFileName string
var outputFileName string

func init() {
	flag.StringVar(&inputFileName, "i", "", "input file")
	flag.StringVar(&outputFileName, "o", "", "output file")
}

func main() {
	flag.Parse()
	if inputFileName == "" || outputFileName == "" {
		fmt.Println("no input or output")
		os.Exit(1)
	}
	file, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}
	out := translate.Translate(string(file))
	err = os.WriteFile(outputFileName, []byte(out), 0755)
	if err != nil {
		panic(err)
	}
}
