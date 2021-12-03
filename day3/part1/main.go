package main

import (
	"aoc/commons"
	"fmt"
)

const (
	wordLen = 12
)

func main() {
	fh := commons.NewHandler("input")
	var bytes [wordLen]int
	listLen := 0
	fh.ForEachLineWithoutBlanks(func(s string) {
		listLen++
		for i := 0; i < wordLen; i++ {
			if s[i] == '1' {
				bytes[i]++
			}
		}
	})
	var gamma, epsilon uint64
	for i := 0; i < wordLen; i++ {
		if bytes[i] > listLen/2 {
			gamma |= uint64(1 << (wordLen - i - 1))
		}
	}
	epsilon = 0b0000111111111111 ^ gamma
	fmt.Printf("%d %012b\n", gamma, gamma)
	fmt.Printf("%d %012b\n", epsilon, epsilon)
	fmt.Println(epsilon * gamma)
}
