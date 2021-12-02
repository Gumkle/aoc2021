package main

import (
	"aoc/commons"
	"fmt"
	"os"
	"strings"
)

const (
	down    = "down"
	forward = "forward"
	up      = "up"
)

func main() {
	fh, err := commons.NewHandler("input")
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}

	var dir string
	var val int
	var reader *strings.Reader
	moves := make(map[string]int)

	fh.ForEachLine(func(line string) {
		if line == "" {
			return
		}
		reader = strings.NewReader(line)
		fmt.Fscanf(reader, "%s %d", &dir, &val)
		moves[dir] += val
	})

	depth := moves[down] - moves[up]
	pos := depth * moves[forward]

	fmt.Println(moves, pos)
}
