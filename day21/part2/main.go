package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	// find intersections
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

	// make a list of inert ingredients
	inert := make(map[string]bool, 0)
	for _, l := range labels {
		for _, i := range l.ingredients {
			if _, ok := possibleAllergens[i]; !ok {
				inert[i] = true
			}
		}
	}
	fmt.Println("Inert")
	for i := range inert {
		fmt.Printf("%s\n", i)
	}

	// remove the inert ingredients from the list
	labelsMissingInert := make([]*label, 0)
	for _, l := range labels {
		labelsMissingInert = append(labelsMissingInert, l.removeInert(inert))
	}
	fmt.Println("Labels without inert")
	for _, l := range labelsMissingInert {
		fmt.Printf("\tLabel: %+v\n", l)
	}

	// key by allergen
	labelsByAllergen = make(map[string][]*label, 0)
	for _, l := range labelsMissingInert {
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

	// print the labels by allergen
	for allergen, labels := range labelsByAllergen {
		fmt.Printf("Allergen %s\n", allergen)
		for _, l := range labels {
			fmt.Printf("\tLabel %s\n", l)
		}
	}

	// Find the initial intersection
	ingredientsByAllergen := make(map[string][]string, 0)
	for allergen, labels := range labelsByAllergen {
		fmt.Printf("Allergen %s\n", allergen)
		intersection := labels[0].ingredients
		for i := 1; i < len(labels); i++ {
			intersection = findIntersection(intersection, labels[i].ingredients)
		}
		ingredientsByAllergen[allergen] = intersection
	}
	fmt.Println("Ingredients by Allergen")
	for allergen, intersection := range ingredientsByAllergen {
		fmt.Printf("Allergen %s\n", allergen)
		for _, i := range intersection {
			fmt.Printf("\tIngredient %s\n", i)
		}
	}

	// Iterate by finding the alergens that have a unique ingredient causing it... THEN remove those from the list of ingredients for the others
	// Use recursion... the base case would be when all alergens have a single ingredient
	result := deduce(ingredientsByAllergen)
	fmt.Printf("%+v\n", result)

	// Get the keys and sort them
	allergens := make([]string, 0)
	for a := range result {
		allergens = append(allergens, a)
	}
	sort.Strings(allergens)

	answer := make([]string, 0)
	for _, a := range allergens {
		ingredients := result[a]
		answer = append(answer, ingredients[0])
	}
	fmt.Printf("Answer: %+v\n", strings.Join(answer, ","))
}

func deduce(ingredientsByAllergen map[string][]string) map[string][]string {
	// Base case... one one ingredient in each list
	finished := true
	for _, i := range ingredientsByAllergen {
		if len(i) > 1 {
			finished = false
			break
		}
	}
	if finished {
		return ingredientsByAllergen
	}

	// Make a list of ones that are already resolved
	canRemove := make(map[string]bool, 0)
	for _, i := range ingredientsByAllergen {
		if len(i) == 1 {
			canRemove[i[0]] = true
		}
	}

	// Remove these ones from the others
	result := make(map[string][]string, 0)
	for a, i := range ingredientsByAllergen {
		if len(i) > 1 {
			result[a] = removeIngredients(i, canRemove)
		} else {
			result[a] = i
		}
	}
	return deduce(result)
}

func removeIngredients(ingredients []string, canRemove map[string]bool) []string {
	result := make([]string, 0)
	for _, i := range ingredients {
		if _, ok := canRemove[i]; !ok {
			result = append(result, i)
		}
	}
	return result
}

func (l *label) removeInert(inert map[string]bool) *label {
	ingredients := make([]string, 0)
	for _, i := range l.ingredients {
		if _, ok := inert[i]; !ok {
			ingredients = append(ingredients, i)
		}
	}

	return &label{
		ingredients: ingredients,
		allergens:   l.allergens,
	}
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
