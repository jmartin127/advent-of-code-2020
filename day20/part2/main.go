package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type puzzle struct {
	matrix [][]*tile
}

type tile struct {
	id    int
	image []*row
}

type row struct {
	vals []bool
}

func newPuzzle(size int) *puzzle {
	matrix := make([][]*tile, 0)

	for i := 0; i < size; i++ {
		row := make([]*tile, 0)
		for j := 0; j < size; j++ {
			row = append(row, nil)
		}
		matrix = append(matrix, row)
	}

	return &puzzle{
		matrix: matrix,
	}
}

func newTile(id int) *tile {
	image := make([]*row, 0)
	return &tile{
		id:    id,
		image: image,
	}
}

func (t *tile) copy() *tile {
	newImage := make([]*row, 0)
	for _, r := range t.image {
		newImage = append(newImage, r.copy())
	}

	return &tile{
		id:    t.id,
		image: newImage,
	}
}

func (r *row) copy() *row {
	newVals := make([]bool, 0)
	for _, v := range r.vals {
		newVals = append(newVals, v)
	}

	return &row{
		vals: newVals,
	}
}

func (t *tile) addRow(r *row) {
	t.image = append(t.image, r)
}

func (t *tile) print() {
	for _, r := range t.image {
		for _, v := range r.vals {
			if v {
				fmt.Printf("%s", "#")
			} else {
				fmt.Printf("%s", ".")
			}
		}
		fmt.Println()
	}
}

/*
Tile 2477:
....#...#.
#..##...#.
...#.....#
..#...#.#.
#.#......#
.#.#######
..#.#...#.
.#.....#..
#..#......
.###.####.

Tile 2609:
*/
func main() {
	file, err := os.Open("/Users/jeff.martin@getweave.com/go/src/github.com/jmartin127/advent-of-code-2020/day20/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	tiles := make([]*tile, 0)
	var currentTile *tile
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "Tile") {
			if currentTile != nil {
				tiles = append(tiles, currentTile)
			}
			tileID := parseTileID(line)
			currentTile = newTile(tileID)
		} else if line == "" {
			// nothing
		} else {
			r := parseRow(line)
			currentTile.addRow(r)
		}
	}
	tiles = append(tiles, currentTile)

	fmt.Printf("num tiles %d\n", len(tiles))
	size := int(math.Sqrt(float64(len(tiles))))
	fmt.Printf("size %d\n", size)

	// compare each tile to each other tile (excluding self). If has exactly 2 matches, it is a corner piece
	answer := 1
	corners := make([]*tile, 0)
	edges := make([]*tile, 0)
	middle := make([]*tile, 0)
	for i := 0; i < len(tiles); i++ {
		var numMatchesForTile int
		t1 := tiles[i]
		for j := 0; j < len(tiles); j++ {
			if i != j {
				t2 := tiles[j]
				if numSharedEdges(t1, t2) > 0 {
					numMatchesForTile++
				}
			}
		}
		fmt.Printf("ID: %d, %d\n", t1.id, numMatchesForTile)
		if numMatchesForTile == 2 {
			fmt.Printf("Found corner! %d\n", t1.id)
			answer *= t1.id
			corners = append(corners, t1)
		} else if numMatchesForTile == 3 {
			edges = append(edges, t1)
		} else if numMatchesForTile == 4 {
			middle = append(middle, t1)
		} else {
			fmt.Printf("TILE IS ODD %d, %d", t1.id, numMatchesForTile)
		}
	}

	// Build it like an actual puzzle!  Organize by corners, edges, middle pieces
	fmt.Printf("Num corners %d\n", len(corners))
	fmt.Printf("Num edges %d\n", len(edges))
	fmt.Printf("Num middle %d\n", len(middle))
	fmt.Printf("Num total %d\n", len(corners)+len(middle)+len(edges))

	// Grab a corner and start building!
	// Add matching edge pieces, the number to add would be: edge_size - 2
	firstCorner := corners[0]
	match := findMatchingEdgePiece(firstCorner, edges)

	fmt.Printf("Starting piece %d\n", firstCorner.id)
	fmt.Printf("Match %d\n", match.id)

}

func buildTopEdge(t *tile, edges []*tile) {

}

func findMatchingEdgePiece(t *tile, edges []*tile) *tile {
	for _, edge := range edges {
		if numSharedEdges(t, edge) > 0 {
			return edge
		}
	}

	return nil
}

// 4 orientations:
//   2nd normal
//   2nd rotated
// THEN flip
//   2nd normal
//   2nd rotated
func numSharedEdges(t1 *tile, t2 *tile) int {
	//fmt.Printf("COMPARE %d, %d\n", t1.id, t2.id)

	// a := rotateTile90DegressLeft(t1)
	// a = flipTileHorizontal(a)
	// fmt.Println("A")
	// a.print()
	// b := flipTileHorizontal(t2)
	// fmt.Println("B")
	// b.print()

	// r := compareTiles(a, b)
	// fmt.Printf("R VALUE: %d\n", r)

	if r := compareTiles(t1, t2); r > 0 {
		//fmt.Printf("Num Match: %d\n", r)
		return r
	}
	rotated := rotateTile90DegressLeft(t2)
	if r := compareTiles(t1, rotated); r > 0 {
		//fmt.Printf("Num Match: %d\n", r)
		return r
	}
	rotated = rotateTile90DegressLeft(rotated)
	if r := compareTiles(t1, rotated); r > 0 {
		//fmt.Printf("Num Match: %d\n", r)
		return r
	}
	rotated = rotateTile90DegressLeft(rotated)
	if r := compareTiles(t1, rotated); r > 0 {
		//fmt.Printf("Num Match: %d\n", r)
		return r
	}

	flippedT2 := flipTileHorizontal(t2)
	if r := compareTiles(t1, flippedT2); r > 0 {
		//fmt.Printf("Num Match: %d\n", r)
		return r
	}
	rotated = rotateTile90DegressLeft(flippedT2)
	if r := compareTiles(t1, rotated); r > 0 {
		//fmt.Printf("Num Match: %d\n", r)
		return r
	}
	rotated = rotateTile90DegressLeft(rotated)
	if r := compareTiles(t1, rotated); r > 0 {
		//fmt.Printf("Num Match: %d\n", r)
		return r
	}
	rotated = rotateTile90DegressLeft(rotated)
	if r := compareTiles(t1, rotated); r > 0 {
		//fmt.Printf("Num Match: %d\n", r)
		return r
	}

	//fmt.Printf("Num Match: %d\n", -1)
	return -1
}

func compareTiles(t1 *tile, t2 *tile) int {
	var numMatching int
	if compare1(t1, t2) {
		numMatching++
	}
	if compare2(t1, t2) {
		numMatching++
	}
	if compare3(t1, t2) {
		numMatching++
	}
	if compare4(t1, t2) {
		numMatching++
	}

	return numMatching
}

func compare1(t1 *tile, t2 *tile) bool {
	numColumns := len(t1.image)
	for rowCount, r1 := range t1.image {
		r2 := t2.image[rowCount]

		v1 := r1.vals[numColumns-1]
		v2 := r2.vals[0]
		if v1 != v2 {
			return false
		}
	}

	return true
}

func compare2(t1 *tile, t2 *tile) bool {
	numColumns := len(t1.image)
	for rowCount, r1 := range t1.image {
		r2 := t2.image[rowCount]

		v1 := r1.vals[0]
		v2 := r2.vals[numColumns-1]
		if v1 != v2 {
			return false
		}
	}

	return true
}

func compare3(t1 *tile, t2 *tile) bool {
	numColumns := len(t1.image)
	v1Row := t1.image[len(t1.image)-1]
	v2Row := t2.image[0]
	for i := 0; i < numColumns; i++ {
		v1 := v1Row.vals[i]
		v2 := v2Row.vals[i]
		if v1 != v2 {
			return false
		}
	}

	return true
}

func compare4(t1 *tile, t2 *tile) bool {
	numColumns := len(t1.image)
	v1Row := t1.image[0]
	v2Row := t2.image[len(t2.image)-1]
	for i := 0; i < numColumns; i++ {
		v1 := v1Row.vals[i]
		v2 := v2Row.vals[i]
		if v1 != v2 {
			return false
		}
	}

	return true
}

func flipTileHorizontal(o *tile) *tile {
	//fmt.Printf("\nFLIP!!!\n")
	//o.print()
	t := o.copy()

	result := newTile(t.id)
	for _, r := range t.image {
		newRow := reverseBools(r.vals)
		result.image = append(result.image, &row{
			vals: newRow,
		})
	}

	//fmt.Printf("\nAFTER FLIP!!!\n")
	//result.print()
	return result
}

// flip it, and then swap top/bottom rows
func rotateTile90DegressLeft(o *tile) *tile {
	//fmt.Printf("\nROTATE!!!\n")
	//o.print()
	t := o.copy()

	N := len(t.image)

	// Consider all squares one by one
	for x := 0; x < len(t.image)/2; x++ {
		// Consider elements in group
		// of 4 in current square
		for y := x; y < len(t.image)-x-1; y++ {
			// Store current cell in
			// temp variable
			temp := t.image[x].vals[y]

			// Move values from right to top
			t.image[x].vals[y] = t.image[y].vals[N-1-x]

			// Move values from bottom to right
			t.image[y].vals[N-1-x] = t.image[N-1-x].vals[N-1-y]

			// Move values from left to bottom
			t.image[N-1-x].vals[N-1-y] = t.image[N-1-y].vals[x]

			// Assign temp to left
			t.image[N-1-y].vals[x] = temp
		}
	}

	//fmt.Printf("\nAFTER ROTATE!!!\n")
	//t.print()
	return t
}

func reverseRows(t *tile) {
	for i, j := 0, len(t.image)-1; i < j; i, j = i+1, j-1 {
		t.image[i], t.image[j] = t.image[j], t.image[i]
	}
}

func reverseBools(input []bool) []bool {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
	return input
}

// Tile 2477:
func parseTileID(line string) int {
	parts := strings.Split(line, "Tile ")
	vals := strings.Split(parts[1], ":")
	r, err := strconv.Atoi(vals[0])
	if err != nil {
		panic(err)
	}
	return r
}

// ....#...#.
func parseRow(line string) *row {
	vals := make([]bool, 0)
	for _, r := range []rune(line) {
		if string(r) == "." {
			vals = append(vals, false)
		} else if string(r) == "#" {
			vals = append(vals, true)
		}
	}

	return &row{
		vals: vals,
	}
}
