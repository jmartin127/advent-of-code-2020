package main

import (
	"fmt"
	"testing"
)

// func TestNumSeaMonstersInString(t *testing.T) {
// 	str := ".#.#...#.###...#.##.##..#.#.##.###.#.##.##.#####..##.###.####..#.####.##...#.#..##.##...#..#..###.###.#..####...##..#...#.###...#.##...#.######..###.###.#######..#####...##.#..#..#.#######.####.#..##.########..#..##."
// 	num := numSeaMonstersInString(str)
// 	if num != 2 {
// 		t.Fatalf("expected to find 2, found %d\n", num)
// 	}
// }

func TestSlice(t *testing.T) {
	line := "this is just a test"
	fmt.Printf("%s", line[3:])
	t.Fatal()
}

// index 1
// length(-1)
