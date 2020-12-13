package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type busID struct {
	hasValue bool
	id       int
}

func main() {
	file, err := os.Open("/Users/jeff.martin@getweave.com/go/src/github.com/jmartin127/advent-of-code-2020/day13/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var busIDs []busID
	for scanner.Scan() {
		line := scanner.Text()
		busIDs = lineToBusIDs(line)
	}

	for i := 100000000000000; true; i++ {
		conditionMet := true
		for t, busID := range busIDs {
			if busID.hasValue {
				//fmt.Printf("Checking %d\n", busID.id+t)
				if (i+t)%(busID.id) != 0 {
					conditionMet = false
					break
				}
			} else {
				// condition is met
			}
		}

		if conditionMet {
			fmt.Printf("Answer %d\n", i)
			os.Exit(0)
		}
	}

}

func lineToBusIDs(line string) []busID {
	busIDs := make([]busID, 0)

	fmt.Printf("LINE %+v", line)
	vals := strings.Split(line, ",")
	fmt.Printf("Vals %+v", vals)
	for _, v := range vals {
		var b busID
		if v == "x" {
			b = busID{
				hasValue: false,
			}
		} else {
			id, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			b = busID{
				id:       id,
				hasValue: true,
			}
		}

		busIDs = append(busIDs, b)
	}

	return busIDs
}
