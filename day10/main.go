package main

import (
	"aoc/commons"
	"fmt"
	"sort"
)

var openings = map[rune]rune{
	']': '[',
	')': '(',
	'}': '{',
	'>': '<',
}

var closings = map[rune]rune{
	'[': ']',
	'(': ')',
	'{': '}',
	'<': '>',
}

var corruptionScores = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var completionScores = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func main() {
	fh := commons.NewHandler("input")
	lines := make([]string, 0, 100)
	fh.ForEachLineWithoutBlanks(func(s string) {
		lines = append(lines, s)
	})

	score, cleanLines := countScoreAndExcludeCorruptedLines(lines)
	fmt.Println("Answer to part1:", score)
	fmt.Println("Answer to part2:", countScoreOnIncompleteLines(cleanLines))
}

func countScoreAndExcludeCorruptedLines(lines []string) (int, []string) {
	score := 0
	corruptedLinesMarks := make([]int, len(lines))
	for index, line := range lines {
		corruptedLinesMarks[index] = 0
		symbolStack := make([]rune, 0)
		for _, char := range line {
			if counterpart, ok := openings[char]; !ok {
				symbolStack = append(symbolStack, char)
			} else {
				if symbolStack[len(symbolStack)-1] != counterpart {
					score += corruptionScores[char]
					corruptedLinesMarks[index] = 1
					break
				} else {
					symbolStack = symbolStack[:len(symbolStack)-1]
				}
			}
		}
	}
	cleanLines := make([]string, 0)
	for index, value := range corruptedLinesMarks {
		if value == 0 {
			cleanLines = append(cleanLines, lines[index])
		}
	}
	return score, cleanLines
}

func countScoreOnIncompleteLines(lines []string) int {
	scores := make([]int, len(lines))
	for index, line := range lines {
		score := 0
		stack := make([]rune, 0)
		for _, char := range line {
			if counterpart, ok := openings[char]; !ok {
				stack = append(stack, char)
			} else {
				if stack[len(stack)-1] == counterpart {
					stack = stack[:len(stack)-1]
				} else {
					break
				}
			}
		}
		for len(stack) > 0 {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			score = score*5 + completionScores[closings[pop]]
		}
		scores[index] = score
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}
