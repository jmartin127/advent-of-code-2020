package main

import (
	"bufio"
	"fmt"
	"os"
)

type cube struct {
	plates []*plate
}

type plate struct {
	rows []*row
}

type row struct {
	vals []bool
}

func newCube(len int) *cube {
	plates := make([]*plate, 0)
	for x := 0; x < len; x++ {
		plates = append(plates, newPlate(len))
	}
	return &cube{
		plates: plates,
	}
}

func newPlate(len int) *plate {
	rows := make([]*row, 0)
	for x := 0; x < len; x++ {
		rows = append(rows, newRow(len))
	}
	return &plate{
		rows: rows,
	}
}

func newRow(len int) *row {
	vals := make([]bool, 0)
	for x := 0; x < len; x++ {
		vals = append(vals, false)
	}
	return &row{
		vals: vals,
	}
}

func main() {
	file, err := os.Open("/Users/jeff.martin@getweave.com/go/src/github.com/jmartin127/advent-of-code-2020/day17/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sp := plate{
		rows: make([]*row, 0),
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		sp.rows = append(sp.rows, readRow(line))
	}

	// initialize the cube
	size := 30
	c := newCube(size)
	fmt.Printf("active %d\n", c.numInActiveState())

	// add the starting plate
	fmt.Printf("active %d\n", c.numInActiveState())
	c.addStartingPlateValues(size/2, &sp)

	// apply cycles
	fmt.Printf("size %d\n", c.size())
	for i := 0; i < 6; i++ {
		c = applyCycle(c)
		//c.printCube()
		fmt.Printf("result %d\n", c.numInActiveState())
	}
}

func applyCycle(c *cube) *cube {
	newCube := c.copy()

	size := len(c.plates)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				numActiveNeigh := c.numActiveNeighbors(i, j, k)
				state := c.plates[k].rows[j].vals[i]
				newState := false
				if state == true && (numActiveNeigh == 2 || numActiveNeigh == 3) {
					newState = true
				} else if state == false && numActiveNeigh == 3 {
					newState = true
				}
				newCube.plates[k].rows[j].vals[i] = newState
			}
		}
	}

	return newCube
}

func (c *cube) addStartingPlateValues(pos int, sp *plate) {
	size := len(sp.rows)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			val := sp.rows[j].vals[i]
			c.plates[pos].rows[j+pos].vals[i+pos] = val
		}
	}
}

func (c *cube) printCube() {
	for z, p := range c.plates {
		fmt.Printf("z=%d\n", z)
		p.printPlate()
	}
}

func (c *cube) size() int {
	return len(c.plates)
}

func (c *cube) numActiveNeighbors(x, y, z int) int {
	//fmt.Printf("num active neighbors of x, y, z %d, %d, %d\n", x, y, z)

	var numActive int
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			for k := z - 1; k <= z+1; k++ {
				if c.inBounds(i, j, k) && c.plates[k].rows[j].vals[i] && !isSelf(x, y, z, i, j, k) {
					numActive++
				}
			}
		}
	}
	return numActive
}

func isSelf(x, y, z int, i, j, k int) bool {
	if x == i && y == j && z == k {
		return true
	}
	return false
}

func (c *cube) inBounds(x, y, z int) bool {
	size := len(c.plates)

	if x < 0 || x > size-1 {
		return false
	}
	if y < 0 || y > size-1 {
		return false
	}
	if z < 0 || z > size-1 {
		return false
	}
	return true
}

func (c *cube) numInActiveState() int {
	var result int
	for _, p := range c.plates {
		result += p.numInActiveState()
	}
	return result
}

func (p *plate) numInActiveState() int {
	var result int
	for _, r := range p.rows {
		result += r.numInActiveState()
	}
	return result
}

func (p *plate) printPlate() {
	for _, r := range p.rows {
		for _, v := range r.vals {
			if v {
				fmt.Printf("%s", "#")
			} else {
				fmt.Printf("%s", ".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (r *row) numInActiveState() int {
	var result int
	for _, v := range r.vals {
		if v {
			result++
		}
	}
	return result
}

func (c *cube) copy() *cube {
	copy := newCube(0)
	for _, p := range c.plates {
		copy.plates = append(copy.plates, p.copy())
	}
	return copy
}

func (p *plate) copy() *plate {
	c := newPlate(0)
	for _, r := range p.rows {
		c.rows = append(c.rows, r.copy())
	}
	return c
}

func (r *row) copy() *row {
	c := newRow(0)
	for _, v := range r.vals {
		c.vals = append(c.vals, v)
	}
	return c
}

func readRow(line string) *row {
	r := &row{
		vals: make([]bool, 0),
	}
	for _, v := range []rune(line) {
		var active bool
		if string(v) == "#" {
			active = true
		}
		r.vals = append(r.vals, active)
	}
	return r
}
