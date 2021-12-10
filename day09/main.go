package main

import (
	"aoc/commons"
	"fmt"
	"sort"
)

const (
	markedBasin    = 10
	maxBasinHeight = 8
)

type cavernsMap [][]uint8

func main() {
	fh := commons.NewHandler("input")
	cMap := make(cavernsMap, 0, 100)
	line := 0
	fh.ForEachLineWithoutBlanks(func(s string) {
		mapLine := make([]uint8, 0, 100)
		for _, char := range s {
			mapLine = append(mapLine, uint8(char-'0'))
		}
		cMap = append(cMap, mapLine)
		line++
	})

	fmt.Println("Answer to part1:", sumRiskLevels(cMap))
	fmt.Println("Answer to part2:", multiplyThreeLargestBasins(cMap))
}

func multiplyThreeLargestBasins(cMap cavernsMap) uint {
	basinSizes := make([]int, 0)
	for i := 0; i < len(cMap); i++ {
		for j := 0; j < len(cMap[i]); j++ {
			if cMap[i][j] > maxBasinHeight {
				continue
			}
			basinSizes = append(basinSizes, cMap.getBasinSizeAndMarkPlace(i, j))
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	return uint(basinSizes[0] * basinSizes[1] * basinSizes[2])
}

func (cm cavernsMap) getBasinSizeAndMarkPlace(initialX, initialY int) int {
	size := 0
	checkQueue := make([][2]int, 0)
	checkQueue = append(checkQueue, [2]int{initialX, initialY})

	for len(checkQueue) > 0 {
		coords := checkQueue[0]
		x, y := coords[0], coords[1]
		checkQueue = checkQueue[1:]
		if cm[x][y] > maxBasinHeight {
			continue
		}
		size++
		cm[x][y] = markedBasin
		if x != 0 && cm[x-1][y] <= maxBasinHeight {
			checkQueue = append(checkQueue, [2]int{x - 1, y})
		}
		if x != len(cm)-1 && cm[x+1][y] <= maxBasinHeight {
			checkQueue = append(checkQueue, [2]int{x + 1, y})
		}
		if y != 0 && cm[x][y-1] <= maxBasinHeight {
			checkQueue = append(checkQueue, [2]int{x, y - 1})
		}
		if y != len(cm[x])-1 && cm[x][y+1] <= maxBasinHeight {
			checkQueue = append(checkQueue, [2]int{x, y + 1})
		}
	}

	return size
}

func sumRiskLevels(cMap cavernsMap) uint64 {
	risk := uint64(0)
	for i := 0; i < len(cMap); i++ {
		for j := 0; j < len(cMap[i]); j++ {
			if cMap.hasLowPointIn(i, j) {
				risk += uint64(cMap[i][j]) + 1
			}
		}
	}
	return risk
}

func (cm cavernsMap) hasLowPointIn(i, j int) bool {
	pointValue := cm[i][j]
	if i != 0 {
		if cm[i-1][j] <= pointValue {
			return false
		}
	}
	if i != len(cm)-1 {
		if cm[i+1][j] <= pointValue {
			return false
		}
	}
	if j != 0 {
		if cm[i][j-1] <= pointValue {
			return false
		}
	}
	if j != len(cm[i])-1 {
		if cm[i][j+1] <= pointValue {
			return false
		}
	}
	return true
}
