package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	inputFlag := flag.String("i", "", "Input file or directory")
	outputFlag := flag.String("o", "", "Output file")
	languageFlag := flag.String("l", "eng", "Tesseract OCR language")
	flag.Parse()

	if *inputFlag == "" || *outputFlag == "" {
		flag.PrintDefaults()
		return
	}

	output, exception := os.OpenFile(*outputFlag, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	handle(exception)

	defer output.Close()

	input, exception := os.Stat(*inputFlag)
	handle(exception)

	inputMode := input.Mode()
	if inputMode.IsRegular() {
		temporaryFile := "/tmp/go-tesseract-" + filepath.Base(*inputFlag)
		tesseract(*inputFlag, temporaryFile, *languageFlag)
		write(filepath.Base(*inputFlag), output, temporaryFile)

		fmt.Println("Done!")
		return
	}

	directory := strings.TrimPrefix(*inputFlag, "/")

	inputs, exception := os.ReadDir(*inputFlag)
	handle(exception)

	for _, v := range inputs {
		temporaryFile := "/tmp/go-tesseract-" + v.Name()
		tesseract(directory+"/"+v.Name(), temporaryFile, *languageFlag)
		write(v.Name(), output, temporaryFile)
	}

	fmt.Println("Done!")
}

func tesseract(i string, o string, l string) {
	command := exec.Command("tesseract", "-l", l, i, o)
	exception := command.Start()
	handle(exception)
}

func write(i string, o *os.File, t string) {
	data, exception := os.ReadFile(t + ".txt")
	handle(exception)

	regex, exception := regexp.Compile(`\s+|\n+`)
	handle(exception)

	o.WriteString("=== Start " + i + " file ===\n")
	o.WriteString(strings.ToLower(regex.ReplaceAllString(string(data), " ")))
	o.WriteString("\n=== End " + i + " file ===\n\n")

	os.Remove(t + ".txt")
	fmt.Println(i, "converted")
}
