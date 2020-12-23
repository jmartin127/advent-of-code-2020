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

var (
	nodeByLabel map[int]*node
	absoluteMax int
)

const numMoves = 10000000

func main() {
	//input := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	input := []int{1, 2, 3, 4, 8, 7, 5, 9, 6}

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

	// add 10 million values
	for l := 10; l <= 1000000; l++ {
		n := newNode(l, prevNode)
		prevNode.next = n
		prevNode = n
		if l > absoluteMax {
			absoluteMax = l
		}
	}
	fmt.Println("Done creating list")

	// load the map so we can lookup by label
	fmt.Println("loading map")
	nodeByLabel = make(map[int]*node, 0)
	currentNode := startNode
	for currentNode.next != nil {
		nodeByLabel[currentNode.label] = currentNode
		currentNode = currentNode.next
	}
	nodeByLabel[currentNode.label] = currentNode
	fmt.Println("Done creating map")

	// connect last with first
	startNode.prev = prevNode
	prevNode.next = startNode

	// run the program
	currentCup := startNode
	for i := 0; i < numMoves; i++ {
		currentCup = executeMove(currentCup)
	}

	// create the answer
	cup1 := nodeByLabel[1]
	fmt.Printf("Next after cup 1 %d\n", cup1.next.label)
	fmt.Printf("2nd cup after cup 1 %d\n", cup1.next.next.label)
	fmt.Printf("Answer %d\n", cup1.next.label*cup1.next.next.label)
}

func executeMove(currentCup *node) *node {
	cupsRemoved := pickUpCups(currentCup.next, 3)
	destCup := determineDestination(currentCup, cupsRemoved)
	insertCups(cupsRemoved, destCup)
	return currentCup.next
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
	nextNodeTmp := afterNode.next
	afterNode.next = n
	n.prev = afterNode
	n.next = nextNodeTmp
	nextNodeTmp.prev = n
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

		// skip the cups removed
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

		if n, ok := nodeByLabel[i]; ok {
			return n
		}
	}

	// didn't find, return max
	for i := absoluteMax; i > 0; i-- {
		// skip the cups removed
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

		if n, ok := nodeByLabel[i]; ok {
			return n
		}
	}

	return nil
}

func pickUpCups(startNode *node, numToPickUp int) []*node {
	tmp := startNode.prev
	startNode.prev.next = startNode.next.next.next
	startNode.next.next.next.prev = tmp

	cupsRemoved := make([]*node, 0)
	for n := startNode; len(cupsRemoved) < numToPickUp; n = n.next {
		cupsRemoved = append(cupsRemoved, n)
	}

	for _, cr := range cupsRemoved {
		nullOutNode(cr)
	}

	return cupsRemoved
}

func nullOutNode(n *node) *node {
	n.prev = nil
	n.next = nil
	return n
}
