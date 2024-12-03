package util

import (
	"bufio"
	"os"
)

func OpenFileOrPanicPlz(path string) *os.File {
	file, err := os.Open("./1/input.txt")
	if err != nil {
		panic(err)
	}
	return file
}

func CloseFileOrPanicPlz(file *os.File) {
	err := file.Close()
	if err != nil {
		panic(err)
	}
}

func GiveMeAScannerPlz(file *os.File) *bufio.Scanner {
	r := bufio.NewReader(file)
	scanner := bufio.NewScanner(r)
	return scanner
}
