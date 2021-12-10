package day

import (
	"aoc/commons"
	"fmt"
	"strings"
)

type LineInput struct {
	*commons.FileHandler
}

func GetInputHandler(path string) *LineInput {
	return &LineInput{
		commons.NewHandler(path),
	}
}

func (lp *LineInput) ForEachCoordsLine(callback func(x1, x2, y1, y2 uint)) {
	var scanner *strings.Reader
	var x1, x2, y1, y2 uint
	lp.ForEachLineWithoutBlanks(func(s string) {
		scanner = strings.NewReader(s)
		fmt.Fscanf(scanner, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		callback(x1, x2, y1, y2)
	})
}

func (lp *LineInput) ForEachPerpendicularLine(callback func(x1, x2, y1, y2 uint)) {
	lp.ForEachCoordsLine(func(x1, x2, y1, y2 uint) {
		if AreCoordsPerpendicularLine(x1, x2, y1, y2) {
			callback(x1, x2, y1, y2)
		}
	})
}
