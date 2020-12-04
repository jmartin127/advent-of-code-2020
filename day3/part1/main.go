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

	rise := 1
	run := 3

	var treeCount int
	var desiredJIndex int
	for i := rise; i < len(matrix); i += rise {
		row := matrix[i]

		desiredJIndex += run
		j := desiredJIndex
		if desiredJIndex >= len(row)-1 {
			j = desiredJIndex % len(row)
		}

		fmt.Printf("i %d, j %d\n", i, j)
		isTree := row[j]

		if isTree {
			treeCount++
		}
	}

	fmt.Printf("Tree Count %d\n", treeCount)

}
