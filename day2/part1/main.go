package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("/Users/jeff.martin@getweave.com/go/src/github.com/jmartin127/advent-of-code-2020/day2/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numValid int
	for scanner.Scan() {
		line := scanner.Text() // e.g., 16-18 h: hhhhhhhhhhhhhhhhhh
		parts := strings.Split(line, ": ")
		policy := parts[0]
		passwd := parts[1]

		policyParts := strings.Split(policy, " ")
		policyMinMax := policyParts[0]
		policyLetter := policyParts[1]
		policyMinMaxParts := strings.Split(policyMinMax, "-")
		policyMinStr := policyMinMaxParts[0]
		policyMaxStr := policyMinMaxParts[1]
		policyMin, err := strconv.Atoi(policyMinStr)
		if err != nil {
			panic(err)
		}
		policyMax, err := strconv.Atoi(policyMaxStr)
		if err != nil {
			panic(err)
		}

		occurrences := strings.Count(passwd, policyLetter)
		if occurrences <= policyMax && occurrences >= policyMin {
			numValid++
		}
	}

	fmt.Printf("num valid %d\n", numValid)
}
