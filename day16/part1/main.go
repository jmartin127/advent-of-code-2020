package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type note struct {
	name string
	be   []*beginEnd
}

type beginEnd struct {
	start int
	end   int
}

type ticket struct {
	vals []int
}

func main() {
	file, err := os.Open("/Users/jeff.martin@getweave.com/go/src/github.com/jmartin127/advent-of-code-2020/day16/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	notes := make([]*note, 0)
	scanner := bufio.NewScanner(file)
	var yourTicket ticket
	nearbyTickets := make([]*ticket, 0)
	var parseNearby bool
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
		} else if strings.Contains(line, " or ") {
			note := parseRangeLine(line)
			notes = append(notes, note)
		} else if strings.Contains(line, "your ticket") {
			scanner.Scan()
			nextLine := scanner.Text()
			yourTicket = ticket{
				vals: lineToIntArray(nextLine),
			}
		} else if strings.Contains(line, "nearby tickets") {
			parseNearby = true
		} else if parseNearby {
			nearbyTicket := &ticket{
				vals: lineToIntArray(line),
			}
			nearbyTickets = append(nearbyTickets, nearbyTicket)
		}
	}

	var invalidCount int
	for _, nt := range nearbyTickets {
		for _, v := range nt.vals {
			var valid bool
			for _, n := range notes {
				for _, be := range n.be {
					if v >= be.start && v <= be.end {
						valid = true
						break
					}
				}
			}
			if !valid {
				invalidCount += v
			}
		}
	}

	fmt.Printf("your ticket %+v", yourTicket)
	fmt.Printf("invalid %d\n", invalidCount)
}

// 7,1,14
func lineToIntArray(line string) []int {
	vals := strings.Split(line, ",")
	r := make([]int, 0)
	for _, v := range vals {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		r = append(r, i)
	}
	return r
}

// class: 1-3 or 5-7
func parseRangeLine(line string) *note {
	parts := strings.Split(line, ": ")

	bes := strings.Split(parts[1], " or ")

	beginEnds := make([]*beginEnd, 0)
	for _, secondPart := range bes {
		vals := strings.Split(secondPart, "-")
		start, err := strconv.Atoi(vals[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(vals[1])
		if err != nil {
			panic(err)
		}
		be := &beginEnd{
			start: start,
			end:   end,
		}
		beginEnds = append(beginEnds, be)
	}

	return &note{
		name: parts[0],
		be:   beginEnds,
	}
}
