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

	validTickets := make([]*ticket, 0)
	for _, nt := range nearbyTickets {
		valid := true
		for _, v := range nt.vals {
			var validNum bool
			for _, n := range notes {
				for _, be := range n.be {
					if v >= be.start && v <= be.end {
						validNum = true
						break
					}
				}
			}
			if !validNum {
				valid = false
				break
			}
		}
		if valid {
			validTickets = append(validTickets, nt)
		}
	}
	fmt.Printf("num valid %d\n", len(validTickets))

	validFieldsByPosition := make(map[int][]string, 0)
	numPositions := len(validTickets[0].vals)
	for i := 0; i < numPositions; i++ {
		validFields := determinePossibleFieldsForPosition(i, validTickets, notes)
		fmt.Printf("pos %d %+v\n", i, validFields)
		validFieldsByPosition[i] = validFields
	}

	/*
		pos 0 [row]
		pos 1 [class row]
		pos 2 [class row seat]
	*/
	// iteratively find ones which have only one possible spot, and remove those from the other lists
	result := resolveMap(validFieldsByPosition)
	fmt.Printf("resolved\n")
	for k, v := range result {
		fmt.Printf("k, v %d, %s\n", k, v)
	}

	fmt.Println("Computing answer")
	answer := 1
	for k, v := range result {
		if strings.Contains(v, "departure") {
			fmt.Printf("k, v %d, %s\n", k, v)
			fmt.Printf("Answer val %d\n", yourTicket.vals[k])
			answer *= yourTicket.vals[k]
		}
	}

	fmt.Printf("Final answer %d\n", answer)

}

func resolveMap(validFieldsByPosition map[int][]string) map[int]string {
	result := make(map[int]string, 0)
	for true {
		if len(validFieldsByPosition) == 0 {
			break
		}
		var posResolved int
		var name string
		posResolved, name, validFieldsByPosition = resolveOne(validFieldsByPosition)
		result[posResolved] = name
	}

	return result
}

/*
	pos 0 [row]
	pos 1 [class row]
	pos 2 [class row seat]
*/
func resolveOne(validFieldsByPosition map[int][]string) (int, string, map[int][]string) {
	var nameToResolve string
	var positionResolved int
	for pos, validFields := range validFieldsByPosition {
		if len(validFields) == 1 {
			nameToResolve = validFields[0]
			positionResolved = pos
			break
		}
	}
	delete(validFieldsByPosition, positionResolved)

	for pos, validFields := range validFieldsByPosition {
		for i, s := range validFields {
			if s == nameToResolve {
				validFieldsByPosition[pos] = remove(validFields, i)
				break
			}
		}
	}

	return positionResolved, nameToResolve, validFieldsByPosition
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func determinePossibleFieldsForPosition(i int, validTickets []*ticket, notes []*note) []string {
	possibleFields := make([]string, 0)
	for _, n := range notes {
		allValid := true
		for _, vt := range validTickets {
			if !valueSatisfiesNote(vt.vals[i], n) {
				allValid = false
				break
			}
		}
		if allValid {
			possibleFields = append(possibleFields, n.name)
		}
	}

	return possibleFields
}

func valueSatisfiesNote(v int, n *note) bool {
	for _, be := range n.be {
		if v >= be.start && v <= be.end {
			return true
		}
	}

	return false
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
