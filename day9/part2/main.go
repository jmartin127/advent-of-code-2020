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

	invalidNum := partA(list)
	startIndex, endIndex := findStartEnd(list, invalidNum)

	max := -1
	min := 10000000000000
	var sum int
	for i := startIndex; i < endIndex+1; i++ {
		if list[i] > max {
			max = list[i]
		}
		if list[i] < min {
			min = list[i]
		}
		sum += list[i]
	}

	fmt.Printf("%d\n", max+min)
}

func findStartEnd(list []int, invalidNum int) (int, int) {
	for i := 0; i < len(list); i++ {
		start := i

		sum := list[i]
		for j := i + 1; j < len(list); j++ {
			sum += list[j]
			if sum == invalidNum {
				end := j
				return start, end
			}
		}
	}

	return -1, -1
}

func partA(list []int) int {
	for i := preambleLen; i < len(list); i++ {
		if !isSumOfPrior(list, list[i], i-preambleLen) {
			return list[i]
		}
	}

	return -1
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
