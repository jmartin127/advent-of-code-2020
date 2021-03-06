package main

import (
	"bufio"
	"fmt"
	"os"
)

type person struct {
	selections []rune
}

type group struct {
	people []person
}

func newGroup() group {
	return group{
		people: []person{},
	}
}

/*
zvxc <-- group
dv
vh
xv
jvem

mxfhdeyikljnz <-- 2nd group
vwzbjmsrgq
*/
func main() {
	file, err := os.Open("/Users/jeff.martin@getweave.com/go/src/github.com/jmartin127/advent-of-code-2020/day6/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	groups := make([]group, 0)
	currentGroup := newGroup()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			groups = append(groups, currentGroup)
			currentGroup = newGroup()
		} else {
			p := person{
				selections: []rune(line),
			}
			currentGroup.people = append(currentGroup.people, p)
		}
	}
	groups = append(groups, currentGroup)

	var totalCount int
	for _, g := range groups {
		totalCount = totalCount + numYes(g)
	}

	fmt.Printf("Count %d\n", totalCount)
}

func numYes(g group) int {
	a := make(map[rune]bool, 0)
	for _, p := range g.people {
		for _, s := range p.selections {
			a[s] = true
		}
	}

	keys := make([]rune, 0)
	for k := range a {
		keys = append(keys, k)
	}

	return len(keys)
}
