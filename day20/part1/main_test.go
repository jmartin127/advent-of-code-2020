package main

import "testing"

func TestParseTileID(t *testing.T) {
	v := parseTileID("Tile 2477:")
	if v != 2477 {
		t.Fatal()
	}
}
