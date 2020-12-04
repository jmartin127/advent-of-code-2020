package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("/Users/jeff.martin@getweave.com/go/src/github.com/jmartin127/advent-of-code-2020/day3/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	matrix := make([][]bool, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := []rune(line)

		row := make([]bool, 0)
		for _, v := range values {
			var tree bool
			if string(v) == "#" {
				tree = true
			}
			row = append(row, tree)
		}

		matrix = append(matrix, row)
	}

	a := determineCount(1, 1, matrix)
	b := determineCount(1, 3, matrix)
	c := determineCount(1, 5, matrix)
	d := determineCount(1, 7, matrix)
	e := determineCount(2, 1, matrix)

	fmt.Printf("Tree Count %d\n", a*b*c*d*e)

}

func determineCount(rise, run int, matrix [][]bool) int {
	var treeCount int
	var desiredJIndex int
	for i := rise; i < len(matrix); i += rise {
		row := matrix[i]

		desiredJIndex += run
		j := desiredJIndex
		if desiredJIndex >= len(row)-1 {
			j = desiredJIndex % len(row)
		}

		isTree := row[j]

		if isTree {
			treeCount++
		}
	}

	return treeCount
}
