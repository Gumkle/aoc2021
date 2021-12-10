package main

import (
	"aoc/commons"
	"fmt"
	"strings"
)

const (
	up      = "up"
	down    = "down"
	forward = "forward"
)

func main() {
	fh := commons.NewHandler("input")

	var aim, val, depth, distance int
	var dir string
	var reader *strings.Reader

	fh.ForEachLineWithoutBlanks(func(line string) {
		reader = strings.NewReader(line)
		fmt.Fscanf(reader, "%s %d", &dir, &val)
		if dir == up {
			aim -= val
		}
		if dir == down {
			aim += val
		}
		if dir == forward {
			depth += val * aim
			distance += val
		}
	})

	fmt.Println(depth * distance)
}
