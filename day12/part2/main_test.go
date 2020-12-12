package main

import (
	"fmt"
	"testing"
)

func TestToPolar(t *testing.T) {
	r, a := toPolar(10, 4)
	fmt.Printf("r a %f %f\n", r, a)
	x, y := toCartesian(r, a)
	fmt.Printf("x, y %d, %d\n", x, y)

	if x != 5 {
		t.Fatal("x wasn't 5")
	}
	if y != 3 {
		t.Fatal("y wasn't 3")
	}
}

func TestRotateLeft(t *testing.T) {
	x, y := rotateRight(-9, -3, 180)

	fmt.Printf("x, y %d, %d\n", x, y)

	if x != 3 {
		t.Fatalf("was %d\n", x)
	}
	if y != -9 {
		t.Fatalf("was %d\n", y)
	}
}
