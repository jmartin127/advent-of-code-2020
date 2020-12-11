package main

import (
	"bufio"
	"fmt"
	"os"
)

type row struct {
	seats []string
}

func main() {
	file, err := os.Open("/Users/jeff.martin@getweave.com/go/src/github.com/jmartin127/advent-of-code-2020/day11/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	matrix := make([]*row, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, lineToRow(line))
	}

	for {
		var numChanges int
		//printMatrix(matrix)
		newMatrix := copyMatrix(matrix)
		for j, r := range matrix {
			for i, s := range r.seats {
				if s == "L" {
					if numAdjacentOccupied(matrix, i, j) == 0 {
						newMatrix[j].seats[i] = "#"
						numChanges++
					}
				} else if s == "#" {
					if numAdjacentOccupied(matrix, i, j) >= 5 {
						newMatrix[j].seats[i] = "L"
						numChanges++
					}
				}
			}
		}
		matrix = newMatrix

		if numChanges == 0 {
			break
		}
	}

	var numOccupied int
	for j, r := range matrix {
		for i := range r.seats {
			if matrix[j].seats[i] == "#" {
				numOccupied++
			}
		}
	}

	fmt.Printf("answer: %d\n", numOccupied)
}

func printMatrix(matrix []*row) {
	for i, r := range matrix {
		for j := range r.seats {
			fmt.Printf(matrix[i].seats[j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func copyMatrix(matrix []*row) []*row {
	copy := make([]*row, 0)
	for _, r := range matrix {
		s := make([]string, 0)
		for _, v := range r.seats {
			s = append(s, v)
		}
		r := &row{
			seats: s,
		}
		copy = append(copy, r)
	}
	return copy
}

func numAdjacentOccupied(matrix []*row, i, j int) int {
	var numOccupied int

	for a := j + 1; a < len(matrix); a++ {
		v := getVal(matrix, i, a)
		if v == "L" {
			break
		}
		if v == "#" {
			numOccupied++
			break
		}
	}
	for a := j - 1; a >= 0; a-- {
		v := getVal(matrix, i, a)
		if v == "L" {
			break
		}
		if v == "#" {
			numOccupied++
			break
		}
	}

	for a := i + 1; a < len(matrix[0].seats); a++ {
		v := getVal(matrix, a, j)
		if v == "L" {
			break
		}
		if v == "#" {
			numOccupied++
			break
		}
	}
	for a := i - 1; a >= 0; a-- {
		v := getVal(matrix, a, j)
		if v == "L" {
			break
		}
		if v == "#" {
			numOccupied++
			break
		}
	}

	for a, b := i+1, j+1; a < len(matrix[0].seats) && b < len(matrix); a, b = a+1, b+1 {
		v := getVal(matrix, a, b)
		if v == "L" {
			break
		}
		if v == "#" {
			numOccupied++
			break
		}
	}

	for a, b := i+1, j-1; a < len(matrix[0].seats) && b >= 0; a, b = a+1, b-1 {
		v := getVal(matrix, a, b)
		//fmt.Printf("v a b, %s %d %d\n", v, a, b)
		if v == "L" {
			break
		}
		if v == "#" {
			numOccupied++
			break
		}
	}

	for a, b := i-1, j-1; a >= 0 && b >= 0; a, b = a-1, b-1 {
		v := getVal(matrix, a, b)
		if v == "L" {
			break
		}
		if v == "#" {
			numOccupied++
			break
		}
	}

	for a, b := i-1, j+1; a >= 0 && b < len(matrix); a, b = a-1, b+1 {
		v := getVal(matrix, a, b)
		if v == "L" {
			break
		}
		if v == "#" {
			numOccupied++
			break
		}
	}

	return numOccupied
}

func getVal(matrix []*row, i, j int) string {
	if i < 0 || i >= len(matrix[0].seats) {
		return "invalid"
	}

	if j < 0 || j >= len(matrix) {
		return "invalid"
	}

	v := matrix[j].seats[i]
	return v
}

/*
The seat layout fits neatly on a grid. Each position is either floor (.), an empty seat (L), or an occupied seat (#). For example, the initial seat layout might look like this:
*/
// L.LL.LL.LL
func lineToRow(seatsStr string) *row {
	seats := make([]string, 0)
	for _, r := range []rune(seatsStr) {
		seats = append(seats, string(r))
	}

	return &row{
		seats: seats,
	}
}
