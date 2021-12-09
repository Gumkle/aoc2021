package part2

import (
	"day8/part1"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	letterCount = 7
)

type sequence struct {
	string
}

type translator struct {
	toDigit map[string]int
	toCode  map[int]*sequence
}

func SumOutputValues(signals, outputs [][]string) uint64 {
	sum := uint64(0)
	for index := range signals {
		sum += getOutputValue(signals[index], outputs[index])
	}
	return sum
}

func getOutputValue(signals, outputs []string) uint64 {
	merged := mergeWithoutBlanks(signals, outputs)
	translator := createTranslatorBasedOn(merged)
	value := ""
	for _, code := range outputs {
		seq := sequence{code}
		seq.sort()
		if code == "" {
			continue
		}
		digit, ok := translator.toDigit[seq.string]
		if !ok {
			fmt.Println("Error retrieving digit from translator")
			os.Exit(1)
		}
		value += strconv.Itoa(digit)
	}
	result, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("Error occured while parsing output value")
		os.Exit(1)
	}
	return uint64(result)
}

func createTranslatorBasedOn(merged []*sequence) *translator {
	t := translator{
		toDigit: make(map[string]int),
		toCode:  make(map[int]*sequence),
	}
	toIterateOver := merged
	for len(toIterateOver) > 0 {
		iteratingOver := toIterateOver
		toIterateOver = make([]*sequence, 0)
		for _, value := range iteratingOver {
			oneCode, isOneSet := t.toCode[1]
			threeCode, isThreeSet := t.toCode[3]
			fourCode, isFourSet := t.toCode[4]
			fiveCode, isFiveSet := t.toCode[5]
			_, isNineSet := t.toCode[9]
			_, isSixSet := t.toCode[6]

			if digit, ok := part1.SimpleDigitMap[len(value.string)]; ok {
				t.toCode[digit] = value
				t.toDigit[value.string] = digit
				continue
			}

			if len(value.string) == 5 && isOneSet && value.contains(oneCode.string) {
				t.toCode[3] = value
				t.toDigit[value.string] = 3
				continue
			}

			if len(value.string) == 5 && isFourSet && isThreeSet && value.contains(fourCode.sub(threeCode.string)) {
				t.toCode[5] = value
				t.toDigit[value.string] = 5
				continue
			}

			if len(value.string) == 5 && isThreeSet && isFiveSet {
				t.toCode[2] = value
				t.toDigit[value.string] = 2
				continue
			}

			if len(value.string) == 6 && isFourSet && isThreeSet && value.contains(fourCode.string) && value.contains(threeCode.string) {
				t.toCode[9] = value
				t.toDigit[value.string] = 9
				continue
			}

			if len(value.string) == 6 && isFiveSet && value.contains(fiveCode.string) {
				t.toCode[6] = value
				t.toDigit[value.string] = 6
				continue
			}

			if len(value.string) == 6 && isNineSet && isSixSet {
				t.toCode[0] = value
				t.toDigit[value.string] = 0
				continue
			}

			toIterateOver = append(toIterateOver, value)
		}
	}
	return &t
}

func mergeWithoutBlanks(signals, outputs []string) []*sequence {
	mergedRaw := append(signals, outputs...)
	merged := make([]*sequence, 0)
	i := 0
	for _, value := range mergedRaw {
		if len(value) > 1 {
			merged = append(merged, &sequence{value})
			merged[i].sort()
			i++
		}
	}
	return merged[:i]
}

func (s *sequence) sort() {
	values := [letterCount]uint8{}
	for _, char := range s.string {
		values[char-'a'] = 1
	}
	s.string = ""
	for index, value := range values {
		if value != 0 {
			s.string += string(index + 'a')
		}
	}
}

func (s *sequence) contains(s2 string) bool {
	for _, char := range s2 {
		if !strings.ContainsRune(s.string, char) {
			return false
		}
	}
	return true
}

func (s *sequence) sub(s2 string) string {
	result := ""
	for _, char := range s.string {
		if !strings.ContainsRune(s2, char) {
			result += string(char)
		}
	}
	return result
}
