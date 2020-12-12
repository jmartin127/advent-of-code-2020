package main

import "testing"

func TestParseInstruction(t *testing.T) {
	i := parseInstruction("F10")
	if i.num != 10 {
		t.Fatalf("expected 10, was %d", i.num)
	}
	if i.direction != "F" {
		t.Fatalf("Expected F was %s\n", i.direction)
	}

	i = parseInstruction("N3")
	if i.num != 3 {
		t.Fatalf("expected 3, was %d", i.num)
	}
	if i.direction != "N" {
		t.Fatalf("Expected N was %s\n", i.direction)
	}
}
