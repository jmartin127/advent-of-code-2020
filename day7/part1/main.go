package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type bag struct {
	color     string
	bagGroups []bagGroup
}

type bagGroup struct {
	num   int
	color string
}

/*
light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain    1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain    3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain     5 faded blue bags, 6 dotted black bags.
faded blue bags contain     no other bags.
dotted black bags contain   no other bags.
*/
func main() {
	file, err := os.Open("/Users/jeff.martin@getweave.com/go/src/github.com/jmartin127/advent-of-code-2020/day7/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	allBags := make([]bag, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		bag := lineToBag(line)
		allBags = append(allBags, bag)
	}

	// find all bags that can contain a gold bag
	final := make(map[string]bool, 0)
	uqPcs := findUniqueParentColors([]string{"shiny gold"}, allBags)
	for true {
		if len(uqPcs) == 0 {
			break
		} else {
			keys := make([]string, 0)
			for k := range uqPcs {
				final[k] = true
				keys = append(keys, k)
			}
			uqPcs = findUniqueParentColors(keys, allBags)
		}
	}

	fmt.Printf("Length %d\n", len(final))

}

func findUniqueParentColors(colors []string, allBags []bag) map[string]bool {
	parentColors := make(map[string]bool, 0)
	for _, b := range allBags {
		for _, bg := range b.bagGroups {
			for _, c := range colors {
				if bg.color == c {
					parentColors[b.color] = true
				}
			}
		}
	}
	return parentColors
}

func lineToBag(line string) bag {
	components := strings.Split(line, " bags contain ")
	color := components[0]
	whatIsContained := components[1]

	bagGroups := make([]bagGroup, 0)
	if strings.Contains(whatIsContained, "no other bags") {
		bagGroups = []bagGroup{}
	} else if strings.Contains(whatIsContained, ", ") { // 1 bright white bag, 2 muted yellow bags
		groupStrings := strings.Split(whatIsContained, ", ")
		for _, gs := range groupStrings {
			bg := groupStringToGroup(gs)
			bagGroups = append(bagGroups, bg)
		}
	} else { // 1 shiny gold bag.
		gs := whatIsContained
		bg := groupStringToGroup(gs)
		bagGroups = append(bagGroups, bg)
	}

	return bag{
		color:     color,
		bagGroups: bagGroups,
	}
}

// 1 shiny gold bag.
// 2 muted yellow bags
func groupStringToGroup(groupString string) bagGroup {

	re := regexp.MustCompile("^.*(\\d) (.+) bag.*$")
	match := re.FindStringSubmatch(groupString)

	num, err := strconv.Atoi(match[1])
	if err != nil {
		panic(err)
	}
	return bagGroup{
		num:   num,
		color: match[2],
	}
}
