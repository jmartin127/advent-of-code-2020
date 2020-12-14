package main

import (
	"fmt"
	"testing"
)

func TestAllPossibleBinaryStrings(t *testing.T) {
	result := allPossibleBinaryStrings([]string{"000000000000000000000000000000X1101X"})
	for _, r := range result {
		fmt.Printf("result: %+s\n", r)
	}
	t.Fatal()
}
