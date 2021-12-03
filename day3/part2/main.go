package main

import (
	"aoc/commons"
	"fmt"
	"strings"
)

const (
	wordLen      = 12
	byMostCommon = 0
	byLessCommon = 1
)

func main() {
	fh := commons.NewHandler("input")
	words := make([]uint16, 0)
	var word uint16
	var scanner *strings.Reader
	fh.ForEachLineWithoutBlanks(func(s string) {
		scanner = strings.NewReader(s)
		fmt.Fscanf(scanner, "%12b", &word)
		words = append(words, word)
	})

	oxygen := uint64(findNumberIn(words, wordLen-1, byMostCommon))
	co2 := uint64(findNumberIn(words, wordLen-1, byLessCommon))

	fmt.Println(oxygen * co2)
}

func findNumberIn(words []uint16, shift uint8, order int) uint16 {
	ones := make([]uint16, 0)
	zeros := make([]uint16, 0)
	for i := range words {
		if words[i]&(1<<shift) > 0 {
			ones = append(ones, words[i])
		} else {
			zeros = append(zeros, words[i])
		}
	}
	fmt.Printf("Order: %d, shift: %d, ones: %d, zeros: %d\n", order, shift, len(ones), len(zeros))
	if len(zeros) == 1 && len(ones) == 0 || (shift == 0 && order == byLessCommon) {
		return zeros[0]
	}
	if len(ones) == 1 && len(zeros) == 0 || (shift == 0 && order == byMostCommon) {
		return ones[0]
	}

	if order == byMostCommon {
		if len(zeros) > len(ones) {
			return findNumberIn(zeros, shift-1, order)
		}
		return findNumberIn(ones, shift-1, order)
	} else {
		if len(zeros) <= len(ones) {
			return findNumberIn(zeros, shift-1, order)
		}
		return findNumberIn(ones, shift-1, order)
	}
}
