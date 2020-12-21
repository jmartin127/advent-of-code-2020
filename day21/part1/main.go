package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type label struct {
	ingredients []string
	allergens   []string
}

func main() {
	file, err := os.Open("/Users/jeff.martin@getweave.com/go/src/github.com/jmartin127/advent-of-code-2020/day21/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	labels := make([]*label, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		labels = append(labels, lineToLabel(line))
	}

	// load into a map, where the key is the allergens
	labelsByAllergen := make(map[string][]*label, 0)
	for _, l := range labels {
		for _, a := range l.allergens {
			if v, ok := labelsByAllergen[a]; ok {
				labelsByAllergen[a] = append(v, l)
			} else {
				newList := make([]*label, 0)
				newList = append(newList, l)
				labelsByAllergen[a] = newList
			}
		}
	}

	possibleAllergens := make(map[string]bool, 0)
	for allergen, labels := range labelsByAllergen {
		fmt.Printf("Allergen %s\n", allergen)
		intersection := labels[0].ingredients
		for i := 1; i < len(labels); i++ {
			intersection = findIntersection(intersection, labels[i].ingredients)
		}
		fmt.Printf("\tIntersection %+v\n", intersection)
		for _, v := range intersection {
			possibleAllergens[v] = true
		}
	}

	var answer int
	for _, l := range labels {
		for _, i := range l.ingredients {
			if _, ok := possibleAllergens[i]; !ok {
				answer++
			}
		}
	}
	fmt.Printf("Answer %d\n", answer)
}

func findIntersection(a1 []string, a2 []string) []string {
	counts := make(map[string]int, 0)
	for _, v := range a1 {
		counts[v] = 1
	}

	result := make([]string, 0)
	for _, v := range a2 {
		if _, ok := counts[v]; ok {
			result = append(result, v)
		}
	}
	return result
}

// mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
func lineToLabel(line string) *label {
	parts := strings.Split(line, " (")

	ingredients := strings.Split(parts[0], " ")

	// contains dairy, fish)
	allertgensStr := parts[1]
	allertgensStr = strings.ReplaceAll(allertgensStr, ")", "")
	allertgensStr = strings.ReplaceAll(allertgensStr, "contains ", "")
	allergens := strings.Split(allertgensStr, ", ")

	return &label{
		ingredients: ingredients,
		allergens:   allergens,
	}
}
