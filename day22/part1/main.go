package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("/Users/jeff.martin@getweave.com/go/src/github.com/jmartin127/advent-of-code-2020/day22/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	player1Deck := make([]int, 0)
	player2Deck := make([]int, 0)
	var addToPlayer2Deck bool
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "P") {
			// nothing
		} else if line == "" {
			addToPlayer2Deck = true
		} else {
			if addToPlayer2Deck {
				player2Deck = append(player2Deck, parseLine(line))
			} else {
				player1Deck = append(player1Deck, parseLine(line))
			}
		}
	}

	// play!
	p1Deck, p2Deck := play(player1Deck, player2Deck)
	var answer int
	if len(p1Deck) > 0 {
		answer = determineScore(p1Deck)
	} else {
		answer = determineScore(p2Deck)
	}
	fmt.Printf("Answer %d\n", answer)
}

func determineScore(winningDeck []int) int {
	var result int
	var count int
	for i := len(winningDeck) - 1; i >= 0; i-- {
		count++
		result += ((count) * winningDeck[i])
		fmt.Printf("Count %d\n", count)
		fmt.Printf("Num %d\n", winningDeck[i])
	}
	return result
}

func play(p1Deck, p2Deck []int) ([]int, []int) {
	for true {
		p1Deck, p2Deck = playRound(p1Deck, p2Deck)
		if len(p1Deck) == 0 || len(p2Deck) == 0 {
			return p1Deck, p2Deck
		}
	}

	return nil, nil
}

func playRound(p1Deck, p2Deck []int) ([]int, []int) {
	// get the top card from each deck
	card1, p1Deck := removeFirstElement(p1Deck)
	card2, p2Deck := removeFirstElement(p2Deck)

	if card1 > card2 {
		p1Deck = append(p1Deck, card1)
		p1Deck = append(p1Deck, card2)
	} else {
		p2Deck = append(p2Deck, card2)
		p2Deck = append(p2Deck, card1)
	}

	return p1Deck, p2Deck
}

func removeFirstElement(a []int) (int, []int) {
	x, a := a[0], a[1:]
	return x, a
}

/*
Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10
*/
func parseLine(line string) int {
	i, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	return i
}
