# Go Tesseract
Image to text converter with [Tesseract OCR](https://github.com/tesseract-ocr/tesseract).

**Note:** This program will run on only Linux distributions. First of all, you have to install Tesseract OCR on your device.

## Arch Linux
```
# pacman -S tesseract tesseract-data-eng
```

## Build
Use Go compiler to build executable.

```
$ go build
```

## Usage
There are 3 flags you can use.

```
-i Input file or directory
-o Output file
-l Tesseract OCR language (Default value is "eng")
```

## Examples
Convert single image to text.

```
$ go-tesseract -i path/to/input.jpg -o output.txt
```

Convert multiple images in directory to text.

```
$ go-tesseract -i /path/to/directory/ -o output.txt
```

Convert single image to text using multiple languages.

```
$ go-tesseract -i path/to/input.jpg -o output.txt -l eng sin
```
