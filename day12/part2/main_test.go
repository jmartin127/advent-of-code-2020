package main

import (
	"fmt"
	"testing"
)

func TestToPolar(t *testing.T) {
	r, a := toPolar(-9, -3)
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
	tests := []struct {
		name      string
		x         int
		y         int
		degrees   int
		expectedX int
		expectedY int
	}{
		{
			name:      "1",
			x:         3,
			y:         5,
			degrees:   90,
			expectedX: -5,
			expectedY: 3,
		},
		{
			name:      "2",
			x:         -3,
			y:         -5,
			degrees:   180,
			expectedX: 3,
			expectedY: 5,
		},
		{
			name:      "3",
			x:         5,
			y:         3,
			degrees:   90,
			expectedX: -3,
			expectedY: 5,
		},
		{
			name:      "4",
			x:         5,
			y:         3,
			degrees:   450,
			expectedX: -3,
			expectedY: 5,
		},
		{
			name:      "5",
			x:         -5,
			y:         -3,
			degrees:   90,
			expectedX: 3,
			expectedY: -5,
		},
		{
			name:      "5",
			x:         -5,
			y:         -3,
			degrees:   360,
			expectedX: -5,
			expectedY: -3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x, y := rotateLeft(tt.x, tt.y, tt.degrees)
			if x != tt.expectedX {
				t.Fatalf("Expected x %d, was %d\n", tt.expectedX, x)
			}
			if y != tt.expectedY {
				t.Fatalf("Expected y %d, was %d\n", tt.expectedY, y)
			}
		})
	}

}

func TestRotateRight(t *testing.T) {
	tests := []struct {
		name      string
		x         int
		y         int
		degrees   int
		expectedX int
		expectedY int
	}{
		{
			name:      "1",
			x:         5,
			y:         3,
			degrees:   90,
			expectedX: 3,
			expectedY: -5,
		},
		{
			name:      "2",
			x:         -3,
			y:         -5,
			degrees:   180,
			expectedX: 3,
			expectedY: 5,
		},
		{
			name:      "3",
			x:         5,
			y:         3,
			degrees:   90,
			expectedX: 3,
			expectedY: -5,
		},
		{
			name:      "4",
			x:         5,
			y:         3,
			degrees:   450,
			expectedX: 3,
			expectedY: -5,
		},
		{
			name:      "5",
			x:         -5,
			y:         -3,
			degrees:   90,
			expectedX: -3,
			expectedY: 5,
		},
		{
			name:      "6",
			x:         -5,
			y:         -3,
			degrees:   360,
			expectedX: -5,
			expectedY: -3,
		},
		{
			name:      "7",
			x:         3,
			y:         -5,
			degrees:   180,
			expectedX: -3,
			expectedY: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x, y := rotateRight(tt.x, tt.y, tt.degrees)
			if x != tt.expectedX {
				t.Fatalf("Expected x %d, was %d\n", tt.expectedX, x)
			}
			if y != tt.expectedY {
				t.Fatalf("Expected y %d, was %d\n", tt.expectedY, y)
			}
		})
	}

}
