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
	p1Deck, p2Deck := play(player1Deck, player2Deck, 1)
	var answer int
	if len(p1Deck) > 0 {
		fmt.Println("Player 1 is the winner!")
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
	}
	return result
}

func play(p1Deck, p2Deck []int, gameCounter int) ([]int, []int) {
	if len(p1Deck) == 0 || len(p2Deck) == 0 {
		return p1Deck, p2Deck
	}

	fmt.Printf("=== Game %d ===\n", gameCounter)
	pastStates := make(map[string]bool, 0)

	var roundCounter int
	for true {
		roundCounter++
		// fmt.Printf("\n-- Round %d (Game %d) --\n", roundCounter, gameCounter)
		// fmt.Printf("Player 1's deck: %+v\n", p1Deck)
		// fmt.Printf("Player 2's deck: %+v\n", p2Deck)

		p1Deck, p2Deck, pastStates = playRound(p1Deck, p2Deck, gameCounter, roundCounter, pastStates)
		if len(p1Deck) == 0 || len(p2Deck) == 0 {
			return p1Deck, p2Deck
		}
	}

	return nil, nil
}

func convertDecksToGameState(p1Deck, p2Deck []int) string {
	p1 := strings.Join(convertToStringArray(p1Deck), ",")
	p2 := strings.Join(convertToStringArray(p2Deck), ",")
	state := fmt.Sprintf("%s:%s", p1, p2)
	return state
}

func convertToStringArray(input []int) []string {
	result := make([]string, 0)
	for _, v := range input {
		result = append(result, strconv.Itoa(v))
	}
	return result
}

func playRound(p1Deck, p2Deck []int, gameCounter int, roundCounter int, pastStates map[string]bool) ([]int, []int, map[string]bool) {
	// check for same cards
	newState := convertDecksToGameState(p1Deck, p2Deck)
	if _, ok := pastStates[newState]; ok {
		fmt.Printf("Found past state for round %d of game %d state %s!\n", gameCounter, roundCounter, newState)
		return p1Deck, []int{}, pastStates
	}

	pastStates[newState] = true
	// fmt.Printf("num past %d\n", len(pastStates))

	// get the top card from each deck
	card1, p1Deck := removeFirstElement(p1Deck)
	card2, p2Deck := removeFirstElement(p2Deck)
	// fmt.Printf("Player 1 plays: %d\n", card1)
	// fmt.Printf("Player 2 plays: %d\n", card2)

	// If both players have at least as many cards remaining in their deck as the value of the card they just drew, the winner of the round is determined by playing a new game of Recursive Combat (see below).
	p1Remaining := len(p1Deck)
	p2Remaining := len(p2Deck)
	var roundWinnerIsPlayer1 bool
	if p1Remaining >= card1 && p2Remaining >= card2 {
		newP1Deck := copyNextNCards(p1Deck, card1)
		newP2Deck := copyNextNCards(p2Deck, card2)
		gameCounter++
		subgameDeck1Result, _ := play(newP1Deck, newP2Deck, gameCounter)
		if len(subgameDeck1Result) > 0 {
			roundWinnerIsPlayer1 = true
		}
	} else {
		if card1 > card2 {
			roundWinnerIsPlayer1 = true
		}
	}

	if roundWinnerIsPlayer1 {
		// fmt.Printf("Player 1 wins round %d of game %d!\n", gameCounter, roundCounter)
		p1Deck = append(p1Deck, card1)
		p1Deck = append(p1Deck, card2)
	} else {
		// fmt.Printf("Player 2 wins round %d of game %d!\n", gameCounter, roundCounter)
		p2Deck = append(p2Deck, card2)
		p2Deck = append(p2Deck, card1)
	}
	return p1Deck, p2Deck, pastStates
}

func copyNextNCards(deck []int, n int) []int {
	result := make([]int, 0)
	for i := 0; i < len(deck); i++ {
		result = append(result, deck[i])
		if len(result) == n {
			return result
		}
	}

	return nil
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
