package main

import "testing"

func TestGroupStringToGroup(t *testing.T) {
	a := groupStringToGroup("1 shiny gold bag.")
	if a.num != 1 {
		t.Errorf("Expected 1, got %d\n", a.num)
	}
	if a.color != "shiny gold" {
		t.Errorf("Expected sg, got %s\n", a.color)
	}
}
