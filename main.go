package main

import (
	"flag"
	"os"
)

var (
	inputFlag    *string
	outputFlag   *string
	languageFlag *string
)

func main() {
	inputFlag = flag.String("I", "", "Input file or directory")
	outputFlag = flag.String("O", "", "Output file")
	languageFlag = flag.String("L", "eng", "Tesseract language")
	flag.Parse()

	if *inputFlag == "" || *outputFlag == "" {
		flag.PrintDefaults()
		return
	}

	output, exception := os.OpenFile(*outputFlag, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	handle(exception)

	defer output.Close()

	input, exception := os.Stat(*inputFlag)
	handle(exception)

	inputMode := input.Mode()
	if inputMode.IsRegular() {
		convertFileToText(*inputFlag, output)
		return
	}

	inputs, exception := os.ReadDir(*inputFlag)
	handle(exception)

	for _, v := range inputs {
		convertDirectoryFileToText(*inputFlag, v.Name(), output)
	}
}
