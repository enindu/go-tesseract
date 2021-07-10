package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func runTesseract(i string, o string) {
	command := exec.Command("tesseract", "-l", *languageFlag, i, o)

	exception := command.Run()
	handle(exception)
}

func convertImageToText(i string, o *os.File, t string) {
	data, exception := os.ReadFile(t + ".txt")
	handle(exception)

	dataString := string(data)

	regex, exception := regexp.Compile(`\s+|\n+`)
	handle(exception)

	fileName := strings.TrimPrefix(strings.TrimSuffix(filepath.Base(t), ".txt"), "go-tesseract-")

	o.WriteString("=== Start " + fileName + " file ===\n")
	o.WriteString(strings.ToLower(regex.ReplaceAllString(dataString, " ")))
	o.WriteString("\n=== End " + fileName + " file ===\n\n")

	fmt.Println("File \"" + fileName + "\" converted")
	os.Remove(t + ".txt")
}

func convertFileToText(i string, o *os.File) {
	temporaryFile := "/tmp/go-tesseract-" + filepath.Base(i)

	runTesseract(i, temporaryFile)
	convertImageToText(i, o, temporaryFile)
}

func convertDirectoryFileToText(d string, i string, o *os.File) {
	temporaryFile := "/tmp/go-tesseract-" + i

	runTesseract(d+"/"+i, temporaryFile)
	convertImageToText(i, o, temporaryFile)
}
