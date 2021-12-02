package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var inp chan int

func main() {
	inp = make(chan int)
	go readFile()

	var currentValue, previousValue, largerCounter int
	var stillOpen bool
	currentValue = <-inp
	for {
		previousValue = currentValue
		currentValue, stillOpen = <-inp
		if stillOpen {
			if currentValue > previousValue {
				largerCounter += 1
			}
		} else {
			break
		}
	}
	fmt.Println(largerCounter)
}

func readFile() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		result, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Error while reading file line")
			os.Exit(1)
		}
		inp <- result
	}
	close(inp)
}
