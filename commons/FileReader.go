package commons

import (
	"bufio"
	"fmt"
	"os"
)

type fileHandler struct {
	ch  chan string
	scn *bufio.Scanner
}

func NewHandler(path string) *fileHandler {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)

	return &fileHandler{
		ch:  make(chan string),
		scn: scanner,
	}
}

func (fh *fileHandler) ReadLine() (string, bool) {
	if fh.scn.Scan() {
		return fh.scn.Text(), true
	} else {
		return "", false
	}
}

func (fh *fileHandler) ForEachLine(callback func(string)) {
	for text, ok := fh.ReadLine(); ok; text, ok = fh.ReadLine() {
		callback(text)
	}
}

func (fh *fileHandler) ForEachLineWithoutBlanks(callback func(string)) {
	fh.ForEachLine(func(s string) {
		if s == "" {
			return
		}
		callback(s)
	})
}
