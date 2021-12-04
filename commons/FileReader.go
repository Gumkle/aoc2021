package commons

import (
	"bufio"
	"fmt"
	"os"
)

type FileHandler struct {
	scn  *bufio.Scanner
	file *os.File
}

func NewHandler(path string) *FileHandler {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)

	return &FileHandler{
		scn:  scanner,
		file: file,
	}
}

func (fh *FileHandler) ReadLine() (string, bool) {
	if fh.scn.Scan() {
		return fh.scn.Text(), true
	} else {
		return "", false
	}
}

func (fh *FileHandler) ForEachLine(callback func(string)) {
	for text, ok := fh.ReadLine(); ok; text, ok = fh.ReadLine() {
		callback(text)
	}
}

func (fh *FileHandler) ForEachLineWithoutBlanks(callback func(string)) {
	fh.ForEachLine(func(s string) {
		if s == "" {
			return
		}
		callback(s)
	})
}

func (fh *FileHandler) MovePointer(distance int64, relative int) {
	fh.file.Seek(distance, relative)
}
