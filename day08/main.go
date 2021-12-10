package main

import (
	"aoc/commons"
	"day8/part1"
	"day8/part2"
	"fmt"
	"strings"
)

const (
	delimiter = "|"
)

func main() {
	fh := commons.NewHandler("input")
	signals := make([][]string, 0)
	outputs := make([][]string, 0)
	fh.ForEachLineWithoutBlanks(func(s string) {
		input := strings.Split(s, "|")
		signals = append(signals, strings.Split(input[0], " "))
		outputs = append(outputs, strings.Split(input[1], " "))
	})

	fmt.Printf("Answer to part1: %d\n", part1.CountDigits(outputs))
	fmt.Printf("Answer to part2: %d\n", part2.SumOutputValues(signals, outputs))
}
