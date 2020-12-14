package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type instructionGroup struct {
	mask *mask
	mem  []*mem
}

type mask struct {
	val []string
}

type mem struct {
	address int
	val     int
}

func NewGroup() instructionGroup {
	return instructionGroup{
		mem: make([]*mem, 0),
	}
}

func main() {
	file, err := os.Open("/Users/jeff.martin@getweave.com/go/src/github.com/jmartin127/advent-of-code-2020/day14/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	groups := make([]instructionGroup, 0)
	scanner := bufio.NewScanner(file)
	currentGroup := NewGroup()
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "mask") {
			if currentGroup.mask != nil {
				groups = append(groups, currentGroup)
				currentGroup = NewGroup()
			}
			m := parseMaskLine(line)
			currentGroup.mask = m
		} else {
			mem := parseMemLine(line)
			currentGroup.mem = append(currentGroup.mem, mem)
		}
	}
	groups = append(groups, currentGroup)

	fmt.Printf("num groups %d\n", len(groups))

	memory := make([]int, 100000)
	for i, g := range groups {
		fmt.Printf("Processing group %d\n", i)
		// loop through each memory update
		for _, m := range g.mem {
			afterMaskValues := applyMask(m.address, g.mask)
			for _, a := range afterMaskValues {
				memory[a] = m.val
				//fmt.Printf("Setting memory index %d to %d\n", a, m.val)
			}
		}
	}

	var result int
	for _, v := range memory {
		result += v
	}
	fmt.Printf("REsult %d\n", result)
}

func applyMask(input int, mask *mask) []int {
	binary := decimalToBinary(input)
	binaryArray := stringToArray(paddedString(binary))
	if len(binaryArray) != 36 {
		os.Exit(1)
	}
	//fmt.Printf("binaryArray %+v\n", binaryArray)
	//fmt.Printf("input %+v\n", binaryArray)
	//fmt.Printf("mask  %+v\n", mask.val)

	result := make([]string, 0)
	for i, m := range mask.val {
		v := binaryArray[i]

		var r string
		if m == "X" {
			r = "X"
		} else if m == "0" {
			r = v
		} else if m == "1" {
			r = "1"
		}
		result = append(result, r)
	}
	//fmt.Printf("result %+v\n", result)

	v := strings.Join(result, "")

	allPossible := allPossibleBinaryStrings([]string{v})

	finalResult := make([]int, 0)
	for _, b := range allPossible {
		d := binaryToDecimal(b)
		//fmt.Printf("Decimal %d, Binary %s\n", d, b)
		finalResult = append(finalResult, d)
	}
	return finalResult
}

// 000000000000000000000000000000X1101X
func allPossibleBinaryStrings(input []string) []string {
	//fmt.Printf("Input to all possible %s\n", input[0])

	for true {
		var numModified int
		result := make([]string, 0)
		for _, v := range input {
			if !strings.Contains(v, "X") { // done with this one
				result = append(result, v)
			} else {
				new := doNextOne(v)
				result = append(result, new...)
				numModified++
			}
		}
		if numModified == 0 {
			return result
		}
		input = result
	}

	return []string{}
}

func doNextOne(v string) []string {
	result := make([]string, 0)
	for i, r := range []rune(v) {
		if string(r) == "X" {
			new := makeTwoNewStrings(i, v)
			result = append(result, new...)
			break
		}
	}
	return result
}

func makeTwoNewStrings(index int, v string) []string {
	one := stringToArray(v)
	two := stringToArray(v)

	one[index] = "0"
	two[index] = "1"

	stringOne := strings.Join(one, "")
	stringTwo := strings.Join(two, "")

	return []string{stringOne, stringTwo}
}

func stringToArray(v string) []string {
	result := make([]string, 0)
	for _, r := range []rune(v) {
		result = append(result, string(r))
	}
	return result
}

func paddedString(v string) string {
	return fmt.Sprintf("%036s", v)
}

func binaryToDecimal(bin string) int {
	n := new(big.Int)
	n, ok := n.SetString(bin, 2)
	if !ok {
		os.Exit(1)
	}

	return int(n.Int64())
}

func decimalToBinary(dec int) string {
	n := int64(dec)
	return strconv.FormatInt(n, 2)
}

// mask = 110X1XX01011X100XX001X00100100X11X10
func parseMaskLine(line string) *mask {
	parts := strings.Split(line, " = ")
	vals := make([]string, 0)
	for _, r := range []rune(parts[1]) {
		vals = append(vals, string(r))
	}

	return &mask{
		val: vals,
	}
}

// mem[36932] = 186083
func parseMemLine(line string) *mem {
	parts := strings.Split(line, " = ")

	v, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	p := strings.Split(line, "[")
	p2 := strings.Split(p[1], "]")
	addresss, err := strconv.Atoi(p2[0])
	if err != nil {
		panic(err)
	}

	return &mem{
		address: addresss,
		val:     v,
	}
}
