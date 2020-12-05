package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("/Users/jeff.martin@getweave.com/go/src/github.com/jmartin127/advent-of-code-2020/day5/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var highest int
	scanner := bufio.NewScanner(file)
	seats := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		seatID := seatIDFromPass(line)
		fmt.Printf("Seat %d\n", seatID)
		if seatID > highest {
			highest = seatID
		}
		seats = append(seats, seatID)
	}

	fmt.Printf("Missing seat %d\n", findMissingSeat(seats))
}

func findMissingSeat(seats []int) int {
	sort.Ints(seats)
	for index, i := range seats {
		a := seats[index]
		b := seats[index+1]
		fmt.Printf("i %d\n", i)
		if b-a > 1 {
			return i + 1
		}
	}

	return -1
}

func seatIDFromPass(pass string) int {
	var row int
	var column int

	low := 0
	high := 127

	colLow := 0
	colHigh := 7

	chars := []rune(pass)
	for i, c := range chars {
		if c == 'F' {
			if i == 6 {
				row = low
			}
			high = high - (((high - low) / 2) + 1)
		} else if c == 'B' {
			if i == 6 {
				row = high
			}
			low = low + (((high - low) / 2) + 1)
		} else if c == 'L' {
			if i == 9 {
				column = colLow
			}
			colHigh = colHigh - (((colHigh - colLow) / 2) + 1)
		} else if c == 'R' {
			if i == 9 {
				column = colHigh
			}
			colLow = colLow + (((colHigh - colLow) / 2) + 1)
		}
	}

	fmt.Printf("Row: %d, Column %d\n", row, column)

	return (row * 8) + column
}
