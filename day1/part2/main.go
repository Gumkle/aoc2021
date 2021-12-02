package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var ch chan int

func main() {
	ch = make(chan int)
	go readFile()
	groupValues := []int{}
	result := 0
	i := 0
	for {
		value, open := <-ch
		if !open {
			break
		}
		groupValues = append(groupValues, value)
		if i > 0 {
			groupValues[i-1] += value
		}
		if i > 1 {
			groupValues[i-2] += value
		}
		i++
	}
	for index := range groupValues {
		if index >= 1 && index <= len(groupValues)-2 && groupValues[index] > groupValues[index-1] {
			result++
		}
	}
	fmt.Println(result)
}

func readFile() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("File opening errror")
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		value, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Line conversion error")
			os.Exit(1)
		}
		ch <- value
	}
	close(ch)
}
