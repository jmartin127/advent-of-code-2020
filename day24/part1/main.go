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
	fmt.Printf("Anser %d\n", answer)
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
