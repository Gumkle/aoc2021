package main

import (
	"day"
	"fmt"
)

func main() {
	input := day.GetInputHandler("../input")
	worldMap := day.NewMap()
	input.ForEachPerpendicularLine(func(x1, x2, y1, y2 uint) {
		worldMap.MarkPerpendicularLine(x1, x2, y1, y2)
	})
	overlaps := 0
	for i := 0; i < day.MapSize; i++ {
		for j := 0; j < day.MapSize; j++ {
			if worldMap[i][j] > 1 {
				overlaps++
			}
		}
	}
	fmt.Println(overlaps)
}
