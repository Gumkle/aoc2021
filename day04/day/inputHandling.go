package day

import (
	"aoc/commons"
	"fmt"
	"strconv"
	"strings"
)

const (
	ShotsSeparator = ","
	Side           = 5
)

func ReadBoards(fh *commons.FileHandler) []*Board {
	boards := make([]*Board, 0)
	newBoard := new(Board)
	i, j := 0, 0
	var scanner *strings.Reader
	fh.ForEachLineWithoutBlanks(func(s string) {
		scanner = strings.NewReader(s)
		for j = 0; j < Side; j++ {
			fmt.Fscan(scanner, &newBoard[i][j])
		}
		i++
		if i >= Side {
			boards = append(boards, newBoard)
			i = 0
			newBoard = new(Board)
		}
	})
	return boards
}

func ReadGivenShots(fh *commons.FileHandler) Shots {
	line, _ := fh.ReadLine()
	words := strings.Split(line, ",")
	givenShots := make(Shots, 0)
	for _, word := range words {
		value, _ := strconv.Atoi(word)
		givenShots = append(givenShots, uint8(value))
	}
	return givenShots
}
