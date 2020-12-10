package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("/Users/jeff.martin@getweave.com/go/src/github.com/jmartin127/advent-of-code-2020/day10/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	list := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("line %s\n", line)

		v, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		list = append(list, v)
	}

	fmt.Printf("List %+v", list)

	diffs := make(map[int]int, 0)
	sort.Ints(list)
	for i := 0; i < len(list)-1; i++ {
		first := list[i]
		second := list[i+1]

		diff := second - first
		fmt.Printf("Diff %d\n", diff)

		if _, ok := diffs[diff]; !ok {
			diffs[diff] = 1
		} else {
			diffs[diff] = diffs[diff] + 1
		}
	}

	diffs[1] = diffs[1] + 1
	diffs[3] = diffs[3] + 1

	fmt.Printf("diff of 1 %d\n", diffs[1])
	fmt.Printf("diff of 3 %d\n", diffs[3])

	answer := diffs[1] * diffs[3]
	fmt.Printf("Anser: %d\n", answer)

}
