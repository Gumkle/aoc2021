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
	winningBoards := make([]*day.Board, 0)
	for index := range givenShots {
		if len(boards) <= 0 {
			break
		}
		takenShots = givenShots[:index]
		boardsWinningThisTurn, boardLeftAfterThisTurn := resolve(takenShots, boards)
		winningBoards = append(winningBoards, boardsWinningThisTurn...)
		boards = boardLeftAfterThisTurn
	}
	board := winningBoards[len(winningBoards)-1]
	sum := board.SumOfUnmarkedBy(takenShots)
	fmt.Println(sum * uint64(takenShots[len(takenShots)-1]))
	os.Exit(0)
}

func resolve(takenShots day.Shots, boards []*day.Board) ([]*day.Board, []*day.Board) {
	winningBoards := make([]*day.Board, 0)
	for index, board := range boards {
		if board.WinsBy(takenShots) {
			winningBoards = append(winningBoards, board)
			boardsToConsider := append(boards[:index], boards[index+1:]...)
			boardsWinning, boardsLeft := resolve(takenShots, boardsToConsider)
			return append(winningBoards, boardsWinning...), boardsLeft
		}
	}
	return winningBoards, boards
}
