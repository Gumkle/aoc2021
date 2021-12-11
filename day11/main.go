package main

import (
	"aoc/commons"
	"fmt"
)

const (
	steps = 100
)

type coords struct {
	x int
	y int
}

type coordQueue []*coords

func main() {
	fh := commons.NewHandler("input")
	octopusMap := make([][]uint8, 0)
	fh.ForEachLineWithoutBlanks(func(s string) {
		line := make([]uint8, 0)
		for _, char := range s {
			line = append(line, uint8(char-'0'))
		}
		octopusMap = append(octopusMap, line)
	})
	score, i, firstSyncStep := 0, 0, 0
	powerQueue := make(coordQueue, 0)
	for {
		exploded := make([]*coords, 0)
		for rowIndex := range octopusMap {
			for colIndex := range octopusMap[rowIndex] {
				powerQueue.push(&coords{x: rowIndex, y: colIndex})
			}
		}
		initialQueueSize := len(powerQueue)
		for {
			element, canContinue := powerQueue.pop()
			alreadyExploded := false
			if !canContinue {
				break
			}
			for _, exp := range exploded {
				if *exp == *element {
					alreadyExploded = true
					break
				}
			}
			if alreadyExploded {
				continue
			}
			octopusMap[element.x][element.y] += 1
			if octopusMap[element.x][element.y] == 10 {
				isLastInRow := element.y == len(octopusMap[element.x])-1
				isFirstInRow := element.y == 0
				isFirstInCol := element.x == 0
				isLastInCol := element.x == len(octopusMap)-1
				if i < steps {
					score += 1
				}
				octopusMap[element.x][element.y] = 0
				exploded = append(exploded, element)
				if !isFirstInCol && !isFirstInRow {
					powerQueue.push(&coords{x: element.x - 1, y: element.y - 1})
				}
				if !isFirstInCol {
					powerQueue.push(&coords{x: element.x - 1, y: element.y})
				}
				if !isFirstInCol && !isLastInRow {
					powerQueue.push(&coords{x: element.x - 1, y: element.y + 1})
				}
				if !isLastInRow {
					powerQueue.push(&coords{x: element.x, y: element.y + 1})
				}
				if !isLastInRow && !isLastInCol {
					powerQueue.push(&coords{x: element.x + 1, y: element.y + 1})
				}
				if !isLastInCol {
					powerQueue.push(&coords{x: element.x + 1, y: element.y})
				}
				if !isLastInCol && !isFirstInRow {
					powerQueue.push(&coords{x: element.x + 1, y: element.y - 1})
				}
				if !isFirstInRow {
					powerQueue.push(&coords{x: element.x, y: element.y - 1})
				}
			}
		}
		i++
		if initialQueueSize == len(exploded) {
			firstSyncStep = i
		}
		if firstSyncStep == i && i >= steps {
			break
		}
	}

	fmt.Println(score, firstSyncStep)
}

func (cq *coordQueue) pop() (*coords, bool) {
	if len(*cq) == 0 {
		return nil, false
	}
	pop := (*cq)[0]
	*cq = (*cq)[1:]
	return pop, true
}

func (cq *coordQueue) push(element *coords) {
	(*cq) = append((*cq), element)
}
