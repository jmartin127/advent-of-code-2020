package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("/Users/jeff.martin@getweave.com/go/src/github.com/jmartin127/advent-of-code-2020/day1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	entries := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		entry, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		entries = append(entries, entry)
	}

	for _, e1 := range entries {
		for _, e2 := range entries {
			if e1+e2 == 2020 {
				a := e1 * e2
				fmt.Printf("Answer: %d\n", a)
				break
			}
		}
	}
}
