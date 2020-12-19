package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type rule struct {
	id          string
	subsetOne   []string
	subsetTwo   []string
	resolvedOne []string
	resolvedTwo []string
}

func main() {
	file, err := os.Open("/Users/jeff.martin@getweave.com/go/src/github.com/jmartin127/advent-of-code-2020/day19/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	messages := make([]string, 0)
	scanner := bufio.NewScanner(file)
	rules := make([]*rule, 0)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, ":") {
			if strings.Contains(line, "a") || strings.Contains(line, "b") {
				rules = append(rules, parseTerminalRule(line))
			} else {
				rules = append(rules, parseRule(line))
			}
		} else if line == "" {
			// nothing
		} else {
			messages = append(messages, line)
		}
	}

	for _, r := range rules {
		fmt.Printf("rule %+v\n", r)
	}

	result := reduceRuleset(rules)
	ruleZero := getRule("0", result)

	valid := make(map[string]bool, 0)
	for _, m := range messages {
		for _, v := range ruleZero.resolvedOne {
			if v == m {
				valid[m] = true
				break
			}
		}
		for _, v := range ruleZero.resolvedTwo {
			if v == m {
				valid[m] = true
				break
			}
		}
	}
	fmt.Printf("Answer: %d\n", len(valid))

	// Now check for subsets of 31/42
	rule31 := getRule("31", result)
	rule42 := getRule("42", result)

	for _, m := range messages {
		if messageMatchesNewRule(m, rule42.getResolved(), rule31.getResolved()) {
			valid[m] = true
		}
	}

	fmt.Printf("Final Answer: %d\n", len(valid))
}

func messageMatchesNewRule(message string, m1 map[string]bool, m2 map[string]bool) bool {
	var size int
	for k := range m1 {
		size = len(k)
		break
	}

	if len(message)%size != 0 {
		return false
	}

	var numM1 int
	var numM2 int
	var foundInSecond bool
	subs := splitMessageToLength(message, size)
	for _, s := range subs {
		_, inM1 := m1[s]
		_, inM2 := m2[s]

		if !inM1 && !inM2 {
			return false
		}

		if inM1 {
			numM1++
		}

		if inM2 {
			numM2++
			foundInSecond = true
		}
		if inM1 && foundInSecond {
			return false
		}
	}

	if numM2 > numM1 || numM1 == 0 || numM2 == 0 || numM1 == numM2 {
		return false
	}

	return true
}

func splitMessageToLength(message string, length int) []string {
	result := make([]string, 0)
	sub := ""
	for _, r := range []rune(message) {
		sub += string(r)
		if len(sub) == length {
			result = append(result, sub)
			sub = ""
		}
	}

	return result
}

func (r *rule) getResolved() map[string]bool {
	result := make(map[string]bool, 0)
	for _, v := range r.resolvedOne {
		result[v] = true
	}
	for _, v := range r.resolvedTwo {
		result[v] = true
	}

	return result
}

func getRule(id string, rules []*rule) *rule {
	for _, r := range rules {
		if r.id == id {
			return r
		}
	}
	return nil
}

func reduceRuleset(rules []*rule) []*rule {
	// base case
	allResolved := true
	resolved := make(map[string]*rule, 0)
	for _, r := range rules {
		if !r.isResolved() {
			allResolved = false
		} else {
			resolved[r.id] = r
		}
	}
	if allResolved {
		return rules
	}

	for _, r := range rules {
		substituteForRule(resolved, r)
	}

	return reduceRuleset(rules)
}

func substituteForRule(resolved map[string]*rule, r *rule) {
	if r.isResolved() {
		return
	}

	if len(r.resolvedOne) == 0 {
		ok, newEntries := resolveSubset(r.subsetOne, resolved)
		if ok {
			r.resolvedOne = newEntries
		}
	}

	if len(r.resolvedTwo) == 0 && len(r.subsetTwo) > 0 {
		ok, newEntries := resolveSubset(r.subsetTwo, resolved)
		if ok {
			r.resolvedTwo = newEntries
		}
	}
}

// subset 2 3
// 2: aa | bb
// 3: ab | ba
//
// OR just a single number
func resolveSubset(subset []string, resolved map[string]*rule) (bool, []string) {
	if len(subset) == 1 {
		resolvedLeft, ok := resolved[subset[0]]
		if !ok {
			return false, []string{}
		}
		result := make([]string, 0)
		result = append(result, resolvedLeft.resolvedOne...)
		result = append(result, resolvedLeft.resolvedTwo...)
		return true, result
	}

	resolvedLeft, ok := resolved[subset[0]]
	if !ok {
		return false, []string{}
	}

	resolvedRight, ok := resolved[subset[1]]
	if !ok {
		return false, []string{}
	}

	result := make([]string, 0)

	for _, v := range resolvedLeft.resolvedOne {
		for _, v2 := range resolvedRight.resolvedOne {
			result = append(result, v+v2)
		}
	}

	for _, v := range resolvedLeft.resolvedOne {
		for _, v2 := range resolvedRight.resolvedTwo {
			result = append(result, v+v2)
		}
	}

	for _, v := range resolvedLeft.resolvedTwo {
		for _, v2 := range resolvedRight.resolvedOne {
			result = append(result, v+v2)
		}
	}

	for _, v := range resolvedLeft.resolvedTwo {
		for _, v2 := range resolvedRight.resolvedTwo {
			result = append(result, v+v2)
		}
	}

	return true, result
}

func (r *rule) isResolved() bool {
	if len(r.subsetOne) > 0 && len(r.resolvedOne) == 0 {
		return false
	}
	if len(r.subsetTwo) > 0 && len(r.resolvedTwo) == 0 {
		return false
	}
	return true
}

// 4: "a"
func parseTerminalRule(line string) *rule {
	parts := strings.Split(line, ": ")
	letterPart := parts[1]
	letter := letterPart[1 : len(letterPart)-1]
	return &rule{
		id:          parts[0],
		resolvedOne: []string{letter},
	}
}

// 4: 116 66
// 48: 29 66 | 129 116
func parseRule(line string) *rule {
	parts := strings.Split(line, ": ")
	idString := parts[0]
	rest := parts[1]

	subRules := strings.Split(rest, " | ")

	subsetTwo := []string{}
	if len(subRules) > 1 {
		subsetTwo = parseSubrule(subRules[1])
	}

	return &rule{
		id:        idString,
		subsetOne: parseSubrule(subRules[0]),
		subsetTwo: subsetTwo,
	}
}

// 129 116
func parseSubrule(line string) []string {
	fmt.Printf("Parsing line %s\n", line)
	parts := strings.Split(line, " ")

	result := []string{parts[0]}
	if len(parts) > 1 {
		result = append(result, parts[1])
	}

	return result
}
