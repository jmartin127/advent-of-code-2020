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

		v, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		list = append(list, v)
	}

	list = append(list, 0)
	sort.Ints(list)

	possiblePathCount := make(map[int]int, len(list))
	possiblePathCount[0] = 1
	for i := 0; i < len(list); i++ {
		currentElement := list[i]

		if i+1 <= len(list)-1 {
			nextElement := list[i+1]
			diff := nextElement - currentElement
			if diff <= 3 {
				possiblePathCount[i+1] = possiblePathCount[i+1] + possiblePathCount[i]
			}
		}

		if i+2 <= len(list)-1 {
			nextElement := list[i+2]

			diff := nextElement - currentElement
			if diff <= 3 {
				possiblePathCount[i+2] = possiblePathCount[i+2] + possiblePathCount[i]
			}
		}

		if i+3 <= len(list)-1 {
			nextElement := list[i+3]

			diff := nextElement - currentElement
			if diff <= 3 {
				possiblePathCount[i+3] = possiblePathCount[i+3] + possiblePathCount[i]
			}
		}
	}

	fmt.Printf("num combos %d\n", possiblePathCount[len(list)-1])

}
