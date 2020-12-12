package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type instruction struct {
	direction string
	num       int
}

func main() {
	file, err := os.Open("/Users/jeff.martin@getweave.com/go/src/github.com/jmartin127/advent-of-code-2020/day12/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	instructions := make([]instruction, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		instructions = append(instructions, parseInstruction(line))
	}

	direction := 0
	y := 0
	x := 0
	for _, i := range instructions {
		//fmt.Printf("x %d, y %d\n", x, y)
		//fmt.Printf("instruction %+v\n", i)

		if i.direction == "N" {
			y += i.num
		} else if i.direction == "S" {
			y -= i.num
		} else if i.direction == "E" {
			x += i.num
		} else if i.direction == "W" {
			x -= i.num
		} else if i.direction == "L" {
			direction = direction + i.num
			if direction >= 360 {
				direction = direction - 360
			}
			fmt.Printf("Direction!!! %d\n", direction)

		} else if i.direction == "R" {
			direction = direction - i.num
			if direction < 0 {
				direction = direction + 360
			}
			fmt.Printf("Direction!!! %d\n", direction)
		} else if i.direction == "F" {
			if direction < 0 || direction > 360 {
				os.Exit(1)
			}

			if direction == 0 {
				x += i.num
			} else if direction == 180 {
				x -= i.num
			} else if direction == 90 {
				y += i.num
			} else if direction == 270 {
				y -= i.num
			} else {
				fmt.Printf("invalid %d\n", direction)
				os.Exit(1)

			}
		}
	}

	fmt.Printf("x %d, y %d\n", x, y)
	a := math.Abs(float64(x)) + math.Abs(float64(y))
	fmt.Printf("A: %f\n", a)

}

// F10
func parseInstruction(line string) instruction {
	direction := line[:1]
	numString := line[1:]

	num, err := strconv.Atoi(numString)
	if err != nil {
		panic(err)
	}
	return instruction{
		direction: direction,
		num:       num,
	}
}
