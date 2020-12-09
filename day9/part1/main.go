package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	preambleLen = 25
)

func main() {
	file, err := os.Open("/Users/jeff.martin@getweave.com/go/src/github.com/jmartin127/advent-of-code-2020/day9/input.txt")
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

	for i := preambleLen; i < len(list); i++ {
		if !isSumOfPrior(list, list[i], i-preambleLen) {
			fmt.Printf("%d\n", list[i])
			break
		}
	}
}

func isSumOfPrior(list []int, num int, index int) bool {
	for i := index; i < index+preambleLen; i++ {
		for j := index; j < i; j++ {
			if i != j {
				if list[i]+list[j] == num {
					return true
				}
			}
		}
	}
	return false
}
