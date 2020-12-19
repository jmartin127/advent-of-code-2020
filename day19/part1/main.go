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

	// rules = replaceRules(rules)
	// for _, r := range rules {
	// 	fmt.Printf("rule %+v\n", r)
	// }

	for _, m := range messages {
		fmt.Printf("message %s\n", m)
	}

	result := reduceRuleset(rules)
	var ruleZero *rule
	for _, r := range result {
		//fmt.Printf("Rule %+v\n", r)
		if r.id == "0" {
			ruleZero = r
		}
	}

	var numMatches int
	for _, m := range messages {
		hasMatch := false
		for _, v := range ruleZero.resolvedOne {
			if v == m {
				hasMatch = true
				break
			}
		}
		for _, v := range ruleZero.resolvedTwo {
			if v == m {
				hasMatch = true
				break
			}
		}
		if hasMatch {
			numMatches++
		}
	}
	fmt.Printf("Answer: %d\n", numMatches)
}

// func replaceRules(rules []*rule) []*rule {

// 	newRules := make([]*rule, 0)
// 	for _, r := range rules {
// 		if r.id == "8" { // replace with 42 contents
// 			fmt.Printf("BAD RULE %+v\n", r)
// 			newRules = append(newRules, r)
// 		} else {
// 			newRules = append(newRules, r)
// 		}
// 	}
// 	return newRules
// }

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
		//fmt.Printf("Rule %+v\n", r)
	}

	return reduceRuleset(rules)
}

func substituteForRule(resolved map[string]*rule, r *rule) {
	if r.isResolved() {
		return
	}

	//fmt.Printf("RULE!!! %+v", r)

	if len(r.resolvedOne) == 0 {
		//fmt.Printf("Resolving subset: %+v. Resolved: %+v\n", r.subsetOne, r.resolvedOne)
		ok, newEntries := resolveSubset(r.subsetOne, resolved)
		if ok {
			r.resolvedOne = newEntries
		}
	}

	if len(r.resolvedTwo) == 0 && len(r.subsetTwo) > 0 {
		//fmt.Printf("Resolving subset: %+v. Resolved: %+v\n", r.subsetTwo, r.resolvedTwo)
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
	//fmt.Printf("RESOLVE SUBSET %+v\n", subset)

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
	//fmt.Printf("RESOLVE LEFT %+v\n", resolvedLeft)
	//fmt.Printf("RESOLVE RIGHT %+v\n", resolvedRight)

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
