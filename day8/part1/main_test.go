package main

import "testing"

func TestArgToInt(t *testing.T) {
	if argToInt("-99") != -99 {
		t.Errorf("Value did not match expected %d\n", argToInt("-99"))
	}
}
