package main

import (
	"fmt"
	"testing"
)

func TestParseLine(t *testing.T) {
	result := parseLine("1 + (2 * 3) + (4 * (5 + 6))")
	for _, r := range result {
		fmt.Printf("r %+v\n", r)
	}

	t.Fatal()
}

func TestParseLine2(t *testing.T) {
	result := parseLine("4 * (5 + 6)")
	for _, r := range result {
		fmt.Printf("r %+v\n", r)
	}

	t.Fatal()
}

func TestParseLine3(t *testing.T) {
	result := parseLine("5 + 6")
	for _, r := range result {
		fmt.Printf("r %+v\n", r)
	}

	t.Fatal()
}
