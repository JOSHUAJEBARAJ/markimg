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
	filename := flag.String("i", "", "file location of the image")

	flag.Parse()

	if *filename == "" {

		flag.PrintDefaults()
		os.Exit(1)
	}

	// reading the file
	imgFile, err := os.Open(*filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer imgFile.Close()
	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	imgBase64Str := base64.StdEncoding.EncodeToString(buf)

	// converting the image
	first := "![Image](data:image/png;base64,"
	middle := imgBase64Str
	last := ")"
	result := first + middle + last
	clipboard.WriteAll(result)
}
