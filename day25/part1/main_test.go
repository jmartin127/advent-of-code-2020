package main

import "testing"

func TestTransformNumTimes(t *testing.T) {
	result := transformSubjectNumber(17807724, 8)
	if result != 14897079 {
		t.Fatalf("expected %d, but was %d\n", 14897079, result)
	}
}

func TestTransformNumTimes2(t *testing.T) {
	result := transformSubjectNumber(5764801, 11)
	if result != 14897079 {
		t.Fatalf("expected %d, but was %d\n", 14897079, result)
	}
}

func TestFindLoopSize(t *testing.T) {
	result := findLoopSize(5764801)
	if result != 8 {
		t.Fatalf("expected %d, but was %d\n", 8, result)
	}
}

func TestFindLoopSize2(t *testing.T) {
	result := findLoopSize(17807724)
	if result != 11 {
		t.Fatalf("expected %d, but was %d\n", 11, result)
	}
}
