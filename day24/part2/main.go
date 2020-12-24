package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	size        = 1000
	refLocation = size / 2
)

type instruction struct {
	directions []string
}

type tile struct {
	flipped bool // false == white, true == black
}

func (t *tile) flip() {
	t.flipped = !t.flipped
}

func (t *tile) copy() *tile {
	return &tile{
		flipped: t.flipped,
	}
}

type floor struct {
	tiles [][]*tile
}

func newFloor(size int) *floor {
	tiles := make([][]*tile, 0)
	for i := 0; i < size; i++ {
		row := make([]*tile, 0)
		for j := 0; j < size; j++ {
			t := &tile{}
			row = append(row, t)
		}
		tiles = append(tiles, row)
	}

	return &floor{
		tiles: tiles,
	}
}

func (f *floor) countFlipped() int {
	var result int
	for i := 0; i < len(f.tiles); i++ {
		for j := 0; j < len(f.tiles); j++ {
			if f.tiles[i][j].flipped {
				result++
			}
		}
	}
	return result
}

func (f *floor) copy() *floor {
	cpy := newFloor(len(f.tiles))
	for i := 0; i < len(f.tiles); i++ {
		for j := 0; j < len(f.tiles); j++ {
			newTile := f.tiles[i][j].copy()
			cpy.tiles[i][j] = newTile
		}
	}
	return cpy
}

func main() {
	file, err := os.Open("/Users/jeff.martin@getweave.com/go/src/github.com/jmartin127/advent-of-code-2020/day24/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	isntructions := make([]*instruction, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		isntructions = append(isntructions, parseLine(line))
	}

	fmt.Println("Instructions")
	for _, ins := range isntructions {
		fmt.Printf("%+v\n", ins)
	}

	fmt.Println("Creating the floor")
	floor := newFloor(size)

	// Run the instructions
	for _, ins := range isntructions {
		followInstruction(floor, ins)
	}

	// Count to get the answer
	answer := floor.countFlipped()
	fmt.Printf("Answer %d\n", answer)

	for i := 0; i < 100; i++ {
		floor = flipTilesForDay(floor)
		fmt.Printf("updated %d\n", floor.countFlipped())
	}
}

func flipTilesForDay(f *floor) *floor {
	// copy the floor
	newFloor := f.copy()

	for i := 2; i < len(f.tiles)-2; i++ {
		for j := 2; j < len(f.tiles)-2; j++ {
			numAdj := numAdjacentBlack(f, j, i)
			t := newFloor.tiles[i][j]
			updateTile(t, numAdj)
		}
	}

	return newFloor
}

func updateTile(t *tile, numAdj int) {
	if t.flipped { // black
		if numAdj == 0 || numAdj > 2 {
			t.flip()
		}
	} else {
		if numAdj == 2 {
			t.flip()
		}
	}
}

func numAdjacentBlack(f *floor, posX, posY int) int {
	var numBlack int

	// se
	if f.tiles[posY-1][posX+1].flipped {
		numBlack++
	}

	// ne
	if f.tiles[posY+1][posX+1].flipped {
		numBlack++
	}

	// nw
	if f.tiles[posY+1][posX-1].flipped {
		numBlack++
	}

	// sw
	if f.tiles[posY-1][posX-1].flipped {
		numBlack++
	}

	// e
	if f.tiles[posY][posX+2].flipped {
		numBlack++
	}

	// w
	if f.tiles[posY][posX-2].flipped {
		numBlack++
	}

	return numBlack
}

func followInstruction(f *floor, ins *instruction) {
	posX := refLocation
	posY := refLocation
	for _, dir := range ins.directions {
		if dir == "se" {
			posX++
			posY--
		} else if dir == "ne" {
			posX++
			posY++
		} else if dir == "nw" {
			posX--
			posY++
		} else if dir == "sw" {
			posX--
			posY--
		} else if dir == "e" {
			posX += 2
		} else if dir == "w" {
			posX -= 2
		} else {
			os.Exit(1)
		}
	}

	f.tiles[posY][posX].flip()
}

/*
sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
*/
func parseLine(line string) *instruction {
	directions := make([]string, 0)
	chars := []rune(line)
	for i := 0; i < len(chars); i++ {
		r := chars[i]
		c := string(r)
		n := ""
		if i+1 <= len(chars)-1 {
			n = string(chars[i+1])
		}
		twoChar := c + n

		if twoChar == "se" || twoChar == "ne" || twoChar == "nw" || twoChar == "sw" {
			directions = append(directions, twoChar)
			i++
		} else if c == "e" || c == "w" {
			directions = append(directions, c)
		} else {
			os.Exit(1)
		}
	}

	return &instruction{
		directions: directions,
	}
}
