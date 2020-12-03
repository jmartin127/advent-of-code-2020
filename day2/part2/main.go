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
		policyPositions := policyParts[0]
		policyLetter := policyParts[1]
		policyPositionsParts := strings.Split(policyPositions, "-")
		policyFirstPosStr := policyPositionsParts[0]
		policySecondPosStr := policyPositionsParts[1]
		policyFirstPos, err := strconv.Atoi(policyFirstPosStr)
		if err != nil {
			panic(err)
		}
		policySecondPos, err := strconv.Atoi(policySecondPosStr)
		if err != nil {
			panic(err)
		}

		passwdRunes := []rune(passwd)

		firstChar := passwdRunes[policyFirstPos-1]
		secondChar := passwdRunes[policySecondPos-1]

		var numMatches int
		if string(firstChar) == policyLetter {
			numMatches++
		}
		if string(secondChar) == policyLetter {
			numMatches++
		}
		if numMatches == 1 {
			numValid++
		}
	}

	fmt.Printf("num valid %d\n", numValid)
}
