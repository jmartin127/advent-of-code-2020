package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("/Users/jeff.martin@getweave.com/go/src/github.com/jmartin127/advent-of-code-2020/day5/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var highest int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		seatID := seatIDFromPass(line)
		if seatID > highest {
			highest = seatID
		}
	}

	fmt.Printf("Highest %d\n", highest)
}

func seatIDFromPass(pass string) int {

	low := 0
	high := 127
	var row int
	var column int

	colLow := 0
	colHigh := 7

	chars := []rune(pass)
	for i, c := range chars {
		if c == 'F' {
			if i == 6 {
				row = low
			}
			diff := int((high-low)/2) + 1
			high = high - diff
		} else if c == 'B' {
			if i == 6 {
				row = high
			}
			diff := int((high-low)/2) + 1
			low = low + diff
		} else if c == 'L' {
			if i == 9 {
				column = colLow
			}
			diff := int((colHigh-colLow)/2) + 1
			colHigh = colHigh - diff
		} else if c == 'R' {
			if i == 9 {
				column = colHigh
			}
			diff := int((colHigh-colLow)/2) + 1
			colLow = colLow + diff
		}
	}

	fmt.Printf("Row: %d, Column %d\n", row, column)

	return (row * 8) + column
}
