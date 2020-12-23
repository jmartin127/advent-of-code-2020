package main

import (
	"fmt"
)

type node struct {
	prev  *node
	next  *node
	label int
}

func newNode(label int, prevNode *node) *node {
	return &node{
		label: label,
		prev:  prevNode,
	}
}

var nodeByLabel map[int]*node

var absoluteMax int

func main() {
	input := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	//input := []int{1, 2, 3, 4, 8, 7, 5, 9, 6}

	// initialize
	prevNode := newNode(input[0], nil)
	startNode := prevNode
	for i := 1; i < len(input); i++ {
		l := input[i]
		n := newNode(l, prevNode)
		prevNode.next = n
		prevNode = n

		if l > absoluteMax {
			absoluteMax = l
		}
	}

	// add 10 million values?
	// TODO set aboslute max
	// for i := 10; i <= 10000000; i++ {
	// 	input = append(input, i)
	// }
	fmt.Println("Done creating array")

	// load them into a map
	fmt.Println("loading map")
	nodeByLabel = make(map[int]*node, 0)
	currentNode := startNode
	for currentNode.next != nil {
		fmt.Printf("Adding to map %d\n", currentNode.label)
		nodeByLabel[currentNode.label] = currentNode
		currentNode = currentNode.next
	}
	nodeByLabel[currentNode.label] = currentNode
	fmt.Println("Done creating map")

	// connect last with first
	startNode.prev = prevNode
	prevNode.next = startNode
	fmt.Printf("Start node %+v\n", startNode)
	fmt.Printf("Last node %+v\n", prevNode)

	currentCup := startNode
	for i := 0; i < 100; i++ {
		fmt.Printf("\n-- move %d\n", i+1)
		fmt.Printf("-- cups ")
		printCups(currentCup)
		fmt.Println()
		currentCup = executeMove(currentCup)
	}

	fmt.Println("\n-- final --")
	fmt.Printf("cups: \n")
	printCups(currentCup)
	fmt.Println()
	fmt.Printf("Current cup %d\n", currentCup.label)
}

func printCups(n *node) {
	startLabel := n.label
	current := n
	for true {
		fmt.Printf("%d ", current.label)
		current = current.next
		if current.label == startLabel {
			break
		}
	}
}

func executeMove(currentCup *node) *node {
	fmt.Printf("Current cup %d\n", currentCup.label)

	//fmt.Println("called pickup")
	cupsRemoved := pickUpCups(currentCup.next, 3)
	//: %+v\n", cupsRemoved)

	//fmt.Println("called determine")
	destCup := determineDestination(currentCup, cupsRemoved)
	fmt.Printf("destination: %d\n", destCup.label)

	//fmt.Println("called insert")
	insertCups(cupsRemoved, destCup)

	// determine a new currentIndex
	// if the insert happened before the currentIndex, need to add 3
	newCurrentNode := currentCup.next
	fmt.Printf("New current node %+v\n", newCurrentNode)

	return newCurrentNode
}

/*
The crab places the cups it just picked up so that they are immediately clockwise of the destination cup. They keep the same order as when they were picked up.
*/
func insertCups(cupsRemoved []*node, destCup *node) {
	insertAfterNode(destCup, cupsRemoved[0])
	insertAfterNode(cupsRemoved[0], cupsRemoved[1])
	insertAfterNode(cupsRemoved[1], cupsRemoved[2])
}

func insertAfterNode(afterNode *node, n *node) {
	//fmt.Printf("ADDING: Setting value %d to index %d\n", value, addingIndex)
	nextNodeTmp := afterNode.next
	afterNode.next = n
	n.prev = afterNode
	n.next = nextNodeTmp
	//fmt.Printf("Input afterward %+v\n", input)
}

/*
The crab selects a destination cup: the cup with a label equal to the current cup's label minus one.
If this would select one of the cups that was just picked up, the crab will keep subtracting one until
it finds a cup that wasn't just picked up. If at any point in this process the value goes below the
lowest value on any cup's label, it wraps around to the highest value on any cup's label instead.

Returns the destination index
*/
func determineDestination(currentCup *node, cupsRemoved []*node) *node {
	desired := currentCup.label - 1

	// find the destination
	for i := desired; i > 0; i-- {
		//fmt.Printf("Looking for a cup with label %d\n", i)
		//fmt.Printf("cups %+v\n", cups)
		//fmt.Printf("indexByLabel %+v\n", indexByLabel)

		// skip the cups removed
		var wasRemoved bool
		for _, cr := range cupsRemoved {
			// cr %d\n", cr.label)
			if cr.label == i {
				wasRemoved = true
				break
			}
		}
		if wasRemoved {
			continue
		}

		if n, ok := nodeByLabel[i]; ok {
			fmt.Printf("First return value %+v\n", n)
			return n
		}
	}

	// didn't find, return max
	for i := absoluteMax; i > 0; i-- {
		fmt.Printf("checking max %d\n", i)
		fmt.Printf("cups removed\n")
		printList(cupsRemoved)
		var wasRemoved bool
		for _, cr := range cupsRemoved {
			if cr.label == i {
				wasRemoved = true
				break
			}
		}
		if wasRemoved {
			continue
		}

		fmt.Printf("Checking map for label %d\n", i)
		if n, ok := nodeByLabel[i]; ok {
			fmt.Printf("Second return max %+v\n", n)
			return n
		}
	}

	fmt.Println("Uh oh... returning nil")
	return nil
}

func printList(cups []*node) {
	for _, c := range cups {
		fmt.Printf("cup %+v\n", c)
	}
}

func pickUpCups(startNode *node, numToPickUp int) []*node {
	cupsRemoved := make([]*node, 0)
	for n := startNode; len(cupsRemoved) < numToPickUp; n = n.next {
		// node %+v\n", n)
		cupsRemoved = append(cupsRemoved, n)
	}

	for _, cr := range cupsRemoved {
		remove(cr)
	}

	//fmt.Printf("Input afterward %+v\n", input)

	return cupsRemoved
}

func remove(n *node) *node {
	n.prev.next = n.next // set the next of the previous to the next
	n.next.prev = n.prev // set thep previous of the next to the previous
	n.prev = nil
	n.next = nil

	return n
}
