package main

import (
	"fmt"
	"log"
	"os"
)

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getFileContents(f *os.File) string {
	finfo, err := f.Stat()
	errCheck(err)

	byteContent := make([]byte, finfo.Size())
	_, err = f.Read(byteContent)

	return string(byteContent)
}

func main() {
	file, err := os.Open("test.fnot")
	errCheck(err)
	defer file.Close()

	code := getFileContents(file)
	fmt.Println(code)
}
