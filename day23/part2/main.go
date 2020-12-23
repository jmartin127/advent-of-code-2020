package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	//input := []int{1, 2, 3, 4, 8, 7, 5, 9, 6}

	// add 10 million values?
	for i := 10; i <= 10000000; i++ {
		input = append(input, i)
	}
	fmt.Println("Done creating array")

	var currentIndex int
	for i := 0; i < 10; i++ {
		fmt.Printf("\n-- move %d\n", i+1)
		// fmt.Printf("-- cups %+v\n", input)
		input, currentIndex = executeMove(input, currentIndex)
	}

	fmt.Println("\n-- final --")
	fmt.Printf("cups: %+v\n", input)
	fmt.Printf("Current cup %d\n", input[currentIndex])
}

func executeMove(input []int, currentCupIndex int) ([]int, int) {
	// fmt.Printf("Current cup %d\n", input[currentCupIndex])
	currentCupLabel := input[currentCupIndex]
	cupsRemoved, newInput := pickUpCups(input, currentCupIndex+1, 3)
	// fmt.Printf("pick up: %+v\n", cupsRemoved)

	destIndex := determineDestination(currentCupLabel, newInput)
	// fmt.Printf("destination: %d\n", newInput[destIndex])

	result := insertCups(newInput, cupsRemoved, destIndex)

	// determine a new currentIndex
	// if the insert happened before the currentIndex, need to add 3
	var numToAdd int
	if destIndex < currentCupIndex {
		numToAdd = 3
	}
	newCurrentIndex := currentCupIndex + 1 + numToAdd
	if newCurrentIndex >= len(result) {
		newCurrentIndex = 0
	}

	return result, newCurrentIndex
}

/*
The crab places the cups it just picked up so that they are immediately clockwise of the destination cup. They keep the same order as when they were picked up.
*/
func insertCups(input []int, cupsRemoved []int, destIndex int) []int {
	insertLocation := destIndex + 1

	result := insertAtIndex(input, insertLocation, cupsRemoved[0])
	result = insertAtIndex(result, insertLocation+1, cupsRemoved[1])
	result = insertAtIndex(result, insertLocation+2, cupsRemoved[2])

	return result
}

func insertAtIndex(input []int, index int, value int) []int {
	input = append(input, 0)
	copy(input[index+1:], input[index:])
	input[index] = value
	return input
}

/*
The crab selects a destination cup: the cup with a label equal to the current cup's label minus one.
If this would select one of the cups that was just picked up, the crab will keep subtracting one until
it finds a cup that wasn't just picked up. If at any point in this process the value goes below the
lowest value on any cup's label, it wraps around to the highest value on any cup's label instead.

Returns the destination index
*/
func determineDestination(currentCupLabel int, cups []int) int {
	desired := currentCupLabel - 1

	// load them into a map
	max := -1
	indexByLabel := make(map[int]int, 0)
	for i, v := range cups {
		indexByLabel[v] = i
		if v > max {
			max = v
		}
	}

	// find the destination
	for i := desired; i > 0; i-- {
		if v, ok := indexByLabel[i]; ok {
			return v
		}
	}

	// didn't find, return max
	return indexByLabel[max]
}

func pickUpCups(input []int, startIndex int, numToPickUp int) ([]int, []int) {
	cups := make([]int, 0)
	indexesRemoved := make([]int, 0)
	for i := startIndex; len(cups) < numToPickUp; i++ {
		if i >= len(input) {
			i = len(input) % i
		}
		val := input[i]
		cups = append(cups, val)
		indexesRemoved = append(indexesRemoved, i)
	}

	// remove them
	sort.Ints(indexesRemoved)
	for i := len(indexesRemoved) - 1; i >= 0; i-- {
		input = remove(input, indexesRemoved[i])
	}

	return cups, input
}

func remove(slice []int, s int) []int {
	fmt.Printf("Removing index %d of length %d\n", s, len(slice))
	return append(slice[:s], slice[s+1:]...)
}
