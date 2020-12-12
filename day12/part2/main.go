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

	y := 0
	x := 0
	wx := 10
	wy := 1
	for _, i := range instructions {
		fmt.Printf("instruction %+v\n", i)

		if i.direction == "N" {
			wy += i.num
		} else if i.direction == "S" {
			wy -= i.num
		} else if i.direction == "E" {
			wx += i.num
		} else if i.direction == "W" {
			wx -= i.num
		} else if i.direction == "L" {
			wx, wy = rotateLeft(wx, wy, i.num)
		} else if i.direction == "R" {
			wx, wy = rotateRight(wx, wy, i.num)
		} else if i.direction == "F" {
			x += wx * i.num
			y += wy * i.num
		}
		fmt.Printf("x %d, y %d\n", x, y)
		fmt.Printf("wx %d, wy %d\n", wx, wy)
	}

	a := math.Abs(float64(x)) + math.Abs(float64(y))
	fmt.Printf("A: %f\n", a)

}

func rotateLeft(x, y, degrees int) (int, int) {
	r, a := toPolar(x, y)

	times := float64(degrees / 90) // number of 90 degree intervals
	a = a + (times * math.Pi / 2)
	return toCartesian(r, a)
}

func rotateRight(x, y, degrees int) (int, int) {
	r, a := toPolar(x, y)

	times := float64(degrees / 90) // number of 90 degree intervals
	a = a - (times * math.Pi / 2)

	return toCartesian(r, a)
}

func toPolar(x, y int) (float64, float64) {
	r := math.Sqrt(float64(x*x + y*y))
	if x < 0 && y < 0 {
		r = r * -1
	}
	a := math.Atan(float64(y) / float64(x))
	return r, a
}

func toCartesian(r, a float64) (int, int) {
	x := r * math.Cos(a)
	y := r * math.Sin(a)

	intX := int(math.Round(x))
	intY := int(math.Round(y))
	return intX, intY
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
