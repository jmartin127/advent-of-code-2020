package main

import (
	"fmt"
	"sort"
)

var indexByLabel map[int]int

func main() {
	input := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	//input := []int{1, 2, 3, 4, 8, 7, 5, 9, 6}

	// add 10 million values?
	// for i := 10; i <= 10000000; i++ {
	// 	input = append(input, i)
	// }
	fmt.Println("Done creating array")

	// load them into a map
	fmt.Println("loading map")
	indexByLabel = make(map[int]int, 0)
	for i, v := range input {
		indexByLabel[v] = i
	}
	fmt.Println("Done creating map")

	var currentIndex int
	for i := 0; i < 100; i++ {
		fmt.Printf("\n-- move %d\n", i+1)
		//fmt.Printf("-- cups %+v\n", input)
		input, currentIndex = executeMove(input, currentIndex)
	}

	fmt.Println("\n-- final --")
	fmt.Printf("cups: %+v\n", input)
	fmt.Printf("Current cup %d\n", input[currentIndex])
}

func executeMove(input []int, currentCupIndex int) ([]int, int) {
	//fmt.Printf("Current cup %d\n", input[currentCupIndex])
	currentCupLabel := input[currentCupIndex]
	cupsRemoved, newInput := pickUpCups(input, currentCupIndex+1, 3)
	//fmt.Printf("pick up: %+v\n", cupsRemoved)

	destIndex := determineDestination(currentCupLabel, newInput)
	//fmt.Printf("destination: %d\n", newInput[destIndex])

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

func insertAtIndex(input []int, addingIndex int, value int) []int {

	// update the indexes of the map as needed
	for label, index := range indexByLabel {
		if index >= addingIndex {
			//fmt.Printf("ADDING: Setting %d to %d\n", label, index+1)
			indexByLabel[label] = index + 1
		}
	}

	indexByLabel[value] = addingIndex
	//fmt.Printf("ADDING: Setting value %d to index %d\n", value, addingIndex)

	input = append(input, 0)
	copy(input[addingIndex+1:], input[addingIndex:])
	input[addingIndex] = value

	//fmt.Printf("Input afterward %+v\n", input)

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

	// find the destination
	for i := desired; i > 0; i-- {
		//fmt.Printf("Looking for a cup with label %d\n", i)
		//fmt.Printf("cups %+v\n", cups)
		//fmt.Printf("indexByLabel %+v\n", indexByLabel)
		if v, ok := indexByLabel[i]; ok {
			//fmt.Printf("First return index %d, value %d\n", v, cups[v])
			return v
		}
	}

	// didn't find, return max
	max := -1
	maxIndex := -1
	for i, c := range cups {
		if c > max {
			max = c
			maxIndex = i
		}
	}

	//fmt.Printf("Second return max %d\n", max)
	return maxIndex
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

		// remove from the label map
		delete(indexByLabel, input[i])
	}

	// remove them from the input slice
	sort.Ints(indexesRemoved)
	for i := len(indexesRemoved) - 1; i >= 0; i-- {
		removingIndex := indexesRemoved[i]
		input = remove(input, removingIndex)

		// update the indexes of the map as needed
		for label, index := range indexByLabel {
			if index > removingIndex {
				//fmt.Printf("Setting %d to %d\n", label, index-1)
				indexByLabel[label] = index - 1
			}
		}
	}

	//fmt.Printf("Input afterward %+v\n", input)

	return cups, input
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
