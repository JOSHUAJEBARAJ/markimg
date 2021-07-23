package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"os"

	"github.com/atotto/clipboard"
)

func main() {
	// flag
	var filename string
	flag.StringVar(&filename, "i", "", "file location of the image")

	flag.Parse()

	if filename == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// reading the file
	imageBase64Str, err := readFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// converting the image

	result := fmt.Sprintf("![Image](data:image/png;base64,%s)", imageBase64Str)

	err = clipboard.WriteAll(result)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func readFile(filename string) (string, error) {
	imgFile, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer imgFile.Close()

	fInfo, _ := imgFile.Stat()
	var size = fInfo.Size()
	buf := make([]byte, size)

	fReader := bufio.NewReader(imgFile)
	_, err = fReader.Read(buf)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buf), nil
}
