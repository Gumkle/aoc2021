package main

import (
	"aoc/commons"
	"day"
	"fmt"
	"os"
)

func main() {
	fh := commons.NewHandler("../day/input")
	givenShots := day.ReadGivenShots(fh)
	boards := day.ReadBoards(fh)

	var takenShots day.Shots
	for index := range givenShots {
		takenShots = givenShots[:index]
		for _, board := range boards {
			if board.WinsBy(takenShots) {
				sum := board.SumOfUnmarkedBy(takenShots)
				fmt.Println(sum * uint64(takenShots[len(takenShots)-1]))
				os.Exit(0)
			}
		}
	}
}
