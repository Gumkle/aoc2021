package main

import (
	"aoc/commons"
	"fmt"
	"strconv"
	"strings"
)

const (
	valueSeparator       = ","
	ages                 = 9
	reproductionInterval = 6
)

func main() {
	numericValues := make([]uint8, 0)
	fh := commons.NewHandler("input")
	line, ok := fh.ReadLine()
	if !ok {
		fmt.Println("Failed to read values from file!")
	}
	values := strings.Split(line, valueSeparator)
	for _, value := range values {
		numericValue, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Faced error while converting values to numbers!")
		}

		numericValues = append(numericValues, uint8(numericValue))
	}
	fmt.Printf("Answer for part1: %d\n", SolveFor(numericValues, 80))
	fmt.Printf("Answer for part2: %d\n", SolveFor(numericValues, 256))
}

func SolveFor(input []uint8, days uint) uint64 {
	ageGroups := [ages]uint64{}
	newAgeGroups := [ages]uint64{}

	for _, value := range input {
		ageGroups[value]++
	}

	for i := uint(0); i < days; i++ {
		for j := 0; j < ages; j++ {
			if j == ages-1 {
				newAgeGroups[j] = ageGroups[0]
				continue
			}
			newAgeGroups[j] = ageGroups[j+1]
			if j == reproductionInterval {
				newAgeGroups[j] += ageGroups[0]
			}
		}
		ageGroups = newAgeGroups
	}

	sum := uint64(0)
	for i := 0; i < ages; i++ {
		sum += ageGroups[i]
	}
	return sum
}
