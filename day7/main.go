package main

import (
	"aoc/commons"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fh := commons.NewHandler("input")
	line, ok := fh.ReadLine()
	if !ok {
		fmt.Println("Failed to read line from  file")
		os.Exit(1)
	}
	rawPositions := strings.Split(line, ",")
	positions := make([]uint, 0)
	maxPosition := 0
	for _, rawPos := range rawPositions {
		pos, err := strconv.Atoi(rawPos)
		if err != nil {
			fmt.Println("Error parsing positions")
			os.Exit(1)
		}
		positions = append(positions, uint(pos))
		if pos > maxPosition {
			maxPosition = pos
		}
	}

	groupedByPosition := make([]uint, maxPosition+1)
	for _, value := range positions {
		groupedByPosition[value]++
	}

	fmt.Printf("Determined minimal fuel consumption: %d\n", determineMinimalFuelConsumption(groupedByPosition))
	fmt.Printf("Determined minimal increasing fuel consumption: %d\n", determineMinimalIncreasingFuelConsumption(groupedByPosition))
}

func determineMinimalFuelConsumption(groupedByPosition []uint) uint {
	minFuel := ^uint(0)
	for i := 0; i < len(groupedByPosition); i++ {
		fuel := uint(0)
		for index, count := range groupedByPosition {
			if i > index {
				fuel += uint(i-index) * count
			} else {
				fuel += uint(index-i) * count
			}
		}
		if fuel < minFuel {
			minFuel = fuel
		}
	}
	return minFuel
}

func determineMinimalIncreasingFuelConsumption(groupedByPosition []uint) uint {
	minFuel := ^uint(0)
	for i := 0; i < len(groupedByPosition); i++ {
		fuel := uint(0)
		for index, count := range groupedByPosition {
			var n uint
			if i > index {
				n = uint(i - index)
			} else {
				n = uint(index - i)
			}
			fuel += (n + (n-1)*n/2) * count
		}
		if fuel < minFuel {
			minFuel = fuel
		}
	}
	return minFuel
}
