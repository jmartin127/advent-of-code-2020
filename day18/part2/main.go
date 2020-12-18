package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type expression struct {
}

type token struct {
	isOperator    bool
	value         int
	operator      string
	isParen       bool
	parenContents string
}

func main() {
	file, err := os.Open("/Users/jeff.martin@getweave.com/go/src/github.com/jmartin127/advent-of-code-2020/day18/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var finalResult int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := parseLine(line)
		r := evaluateExpression(tokens)
		finalResult += r
	}

	fmt.Printf("Final result %d\n", finalResult)
}

// 1 + (2 * 3) + (4 * (5 + 6))
func evaluateExpression(tokens []*token) int {
	if expressionCanBeSolved(tokens) {
		return solveExpression(tokens)
	}

	newExpr := copy(tokens)
	for i, r := range tokens {
		if r.isParen {
			newToken := &token{
				value: evaluateExpression(parseLine(r.parenContents)),
			}

			newExpr[i] = newToken
		}
	}

	return evaluateExpression(newExpr)
}

func copy(tokens []*token) []*token {
	c := make([]*token, 0)
	for _, t := range tokens {
		newT := &token{
			isOperator:    t.isOperator,
			value:         t.value,
			operator:      t.operator,
			isParen:       t.isParen,
			parenContents: t.parenContents,
		}
		c = append(c, newT)
	}

	return c
}

func expressionCanBeSolved(tokens []*token) bool {
	for _, t := range tokens {
		if t.isParen {
			return false
		}
	}
	return true
}

// 5 + 6 + 2
func solveExpression(tokens []*token) int {
	tokens = reducePlus(tokens)
	result := tokens[0].value
	for i := 1; i < len(tokens); i++ {
		t := tokens[i]
		if !t.isOperator {
			operator := tokens[i-1].operator
			if operator == "+" {
				result = result + t.value
			} else if operator == "*" {
				result = result * t.value
			}
		}
	}

	return result
}

// 2 * 1 + 1 + 2 * 3 + 4 * 5 + 6  -->   2 * 4 * 7 * 11
func reducePlus(tokens []*token) []*token {
	if !hasPlus(tokens) {
		return tokens
	}

	reducedExpression := make([]*token, 0)
	var didReduce bool
	for i := 0; i < len(tokens); i++ {
		t := tokens[i]
		if didReduce {
			reducedExpression = append(reducedExpression, t)
		} else if t.isOperator && t.operator == "+" {
			newVal := tokens[i-1].value + tokens[i+1].value
			newToken := &token{
				value: newVal,
			}
			reducedExpression = append(reducedExpression, newToken)
			didReduce = true
			i++
		} else {
			if i+1 <= len(tokens)-1 && tokens[i+1].isOperator && tokens[i+1].operator == "+" { // next is +
				// nothing
			} else {
				reducedExpression = append(reducedExpression, t)
			}
		}
	}

	return reducePlus(reducedExpression)
}

func hasPlus(tokens []*token) bool {
	for _, t := range tokens {
		if t.isOperator && t.operator == "+" {
			return true
		}
	}
	return false
}

// 1 + 2 * 3 + 4 * 5 + 6
func parseLine(line string) []*token {
	tokens := make([]*token, 0)

	var numParen int
	var parenContents string
	for _, r := range []rune(line) {
		if numParen > 0 && string(r) != ")" {
			parenContents += string(r)
			if string(r) == "(" {
				numParen++
			}
		} else if string(r) == "(" {
			numParen++
		} else if string(r) == ")" {
			numParen--
			if numParen == 0 {
				t := &token{
					isParen:       true,
					parenContents: parenContents,
				}
				tokens = append(tokens, t)
				parenContents = ""
			} else {
				parenContents += string(r)
			}
		} else if string(r) == "+" || string(r) == "*" {
			t := &token{
				isOperator: true,
				operator:   string(r),
			}
			tokens = append(tokens, t)
		} else if string(r) == " " {
			// nothing
		} else { // number
			v, err := strconv.Atoi(string(r))
			if err != nil {
				panic(err)
			}
			t := &token{
				value: v,
			}
			tokens = append(tokens, t)
		}
	}

	return tokens
}
