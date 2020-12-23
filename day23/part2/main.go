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

	// load them into a map
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
	fmt.Printf("Start node %+v\n", startNode)
	fmt.Printf("Last node %+v\n", prevNode)

	currentCup := startNode
	for i := 0; i < numMoves; i++ {
		currentCup = executeMove(currentCup)
	}

	// find cup 1
	var cup1 *node
	for true {
		if currentCup.label == 1 {
			cup1 = currentCup
			break
		}
		currentCup = currentCup.next
	}

	fmt.Printf("Next after cup 1 %d\n", cup1.next.label)
	fmt.Printf("2nd cup after cup 1 %d\n", cup1.next.next.label)
	fmt.Printf("Anwer %d\n", cup1.next.label*cup1.next.next.label)
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
	//fmt.Printf("ADDING: Setting value %d to index %d\n", value, addingIndex)
	nextNodeTmp := afterNode.next
	afterNode.next = n
	n.prev = afterNode
	n.next = nextNodeTmp
	nextNodeTmp.prev = n
	//fmt.Printf("Seeting previous for node %+v to %+v\n", n, afterNode)
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
			//fmt.Printf("First return value %+v\n", n)
			return n
		}
	}

	// didn't find, return max
	for i := absoluteMax; i > 0; i-- {
		//fmt.Printf("checking max %d\n", i)
		//fmt.Printf("cups removed\n")
		//printList(cupsRemoved)
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

		//fmt.Printf("Checking map for label %d\n", i)
		if n, ok := nodeByLabel[i]; ok {
			//fmt.Printf("Second return max %+v\n", n)
			return n
		}
	}

	//fmt.Println("Uh oh... returning nil")
	return nil
}

func printList(cups []*node) {
	for _, c := range cups {
		fmt.Printf("cup %+v\n", c)
	}
}

func pickUpCups(startNode *node, numToPickUp int) []*node {
	//fmt.Printf("Removing starting with node %+v\n", startNode)
	//fmt.Printf("Previous %+v\n", startNode.prev)

	tmp := startNode.prev
	//fmt.Printf("Setting the next of %+v to %+v\n", startNode.prev, startNode.next.next.next)
	startNode.prev.next = startNode.next.next.next
	startNode.next.next.next.prev = tmp

	cupsRemoved := make([]*node, 0)
	for n := startNode; len(cupsRemoved) < numToPickUp; n = n.next {
		//fmt.Printf("adding removed cup %+v\n", n)
		cupsRemoved = append(cupsRemoved, n)
	}

	for _, cr := range cupsRemoved {
		nullOutNode(cr)
	}

	//fmt.Printf("Input afterward %+v\n", input)

	return cupsRemoved
}

func nullOutNode(n *node) *node {
	n.prev = nil
	n.next = nil

	return n
}
