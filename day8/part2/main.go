package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	operation string
	argument  int
}

func main() {
	file, err := os.Open("/Users/jeff.martin@getweave.com/go/src/github.com/jmartin127/advent-of-code-2020/day8/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	instructions := make([]instruction, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		instructions = append(instructions, lineToInstruction(line))
	}

	for index, i := range instructions {
		o := i.operation
		if i.operation == "nop" {
			instructions[index].operation = "jmp"
		} else if i.operation == "jmp" {
			instructions[index].operation = "nop"
		}

		if o != i.operation {
			fmt.Printf("changed %s to %s\n", o, i.operation)
		}

		accumulator, finished := runToTermination(instructions)
		if finished {
			fmt.Printf("FINAL %d\n", accumulator)
			break
		}
		instructions[index].operation = o
	}
}

func runToTermination(instructions []instruction) (int, bool) {
	maxLoops := 1000
	count := 0

	var currentIndex int
	var accumulator int
	for {
		count++
		if count > maxLoops {
			return 0, false
		}
		if currentIndex == len(instructions) {
			return accumulator, true
		}
		accumulator, currentIndex = run(instructions, currentIndex, accumulator)
	}
}

func run(allInstructions []instruction, currentIndex int, accumulator int) (int, int) {
	currentInstruction := allInstructions[currentIndex]

	var nextInstructionIndex int
	if currentInstruction.operation == "acc" {
		accumulator = accumulator + currentInstruction.argument
		nextInstructionIndex = currentIndex + 1
	} else if currentInstruction.operation == "jmp" {
		nextInstructionIndex = currentIndex + currentInstruction.argument
	} else if currentInstruction.operation == "nop" {
		nextInstructionIndex = currentIndex + 1
	}

	return accumulator, nextInstructionIndex
}

func lineToInstruction(line string) instruction {
	parts := strings.Split(line, " ")
	operation := parts[0]
	argumentString := parts[1]

	return instruction{
		operation: operation,
		argument:  argToInt(argumentString),
	}
}

func argToInt(arg string) int {
	if string(arg[0]) == "-" {
		v, err := strconv.Atoi(arg[1:])
		if err != nil {
			panic(err)
		}
		return v * -1
	} else {
		v, err := strconv.Atoi(arg[1:])
		if err != nil {
			panic(err)
		}
		return v
	}
}
