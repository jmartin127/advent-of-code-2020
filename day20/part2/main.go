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

func (p *puzzle) print() {
	for _, r := range p.matrix {
		for _, t := range r {
			if t != nil {
				t.print()
			}
			fmt.Println()
		}
		fmt.Println()
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
	fmt.Printf("Tile: %d\n", t.id)
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
				numShared, _, _ := numSharedEdges(t1, t2)
				if numShared > 0 {
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
	// Need to orient the corner so that a matching edge is to the right and left
	firstCorner := corners[0]
	fmt.Printf("Starting piece %d\n", firstCorner.id)
	firstCorner = orientStartingPiece(firstCorner, edges)
	firstCorner = rotateTile90DegressLeft(firstCorner)
	firstCorner = rotateTile90DegressLeft(firstCorner)
	firstCorner = rotateTile90DegressLeft(firstCorner)
	firstCorner = flipTileHorizontal(firstCorner)

	p := newPuzzle(size)
	p.matrix[0][0] = firstCorner
	remainingCorners := removeTileFromSlice(firstCorner, corners)
	fmt.Printf("Starting piece\n")
	firstCorner.print()

	// Build out the top edge
	remainingEdges, remainingCorners := buildTopEdge(firstCorner, p, edges, remainingCorners)
	fmt.Printf("Num edges after building first edge %d\n", len(remainingEdges))
	fmt.Printf("Num corners after building first edge %d\n", len(remainingCorners))
	fmt.Printf("Puzzle after doing top edge:\n")
	p.print()

	// // Put the rest of the pieces in the box
	// remaining := make([]*tile, 0)
	// remaining = addToSlice(remaining, middle)
	// remaining = addToSlice(remaining, remainingCorners)
	// remaining = addToSlice(remaining, remainingEdges)

	// // Build each row
	// remaining = addRowToPuzzle(p, 1, remaining)
	// fmt.Printf("Puzzle after doing middle row:\n")
	// p.print()
}

func orientStartingPiece(t *tile, edges []*tile) *tile {
	var count int
	max := 4
	rotatedTile := t

	for true {
		var foundOrientationOne bool
		var foundOrientationThree bool
		for _, e := range edges {
			numMatch, _, orientationID := numSharedEdges(rotatedTile, e)
			if numMatch > 0 {
				fmt.Printf("orientation %d\n", orientationID)
				if orientationID == 1 {
					foundOrientationOne = true
				} else if orientationID == 3 {
					foundOrientationThree = true
				}
			}
		}
		if foundOrientationOne && foundOrientationThree {
			return rotatedTile
		}

		rotatedTile.print()
		rotatedTile = rotateTile90DegressLeft(rotatedTile)
		rotatedTile.print()

		count++
		if count > max {
			fmt.Println("You messed up Jeff")
			os.Exit(1)
		}
	}

	return nil
}

func addRowToPuzzle(p *puzzle, rowIndex int, pieces []*tile) []*tile {
	remainingPieces := pieces

	for i := 0; i < len(p.matrix); i++ {
		match := findMatchingPiece(p.matrix[rowIndex-1][i], remainingPieces)
		p.matrix[rowIndex][i] = match
		remainingPieces = removeTileFromSlice(match, remainingPieces)
	}

	return remainingPieces
}

func addToSlice(o []*tile, toAdd []*tile) []*tile {
	return append(o, toAdd...)
}

// Find edge_size - 2 edge pieces, starting with t, add them to the puzzle
// ... return the updated list of edges
func buildTopEdge(t *tile, p *puzzle, edges []*tile, corners []*tile) ([]*tile, []*tile) {
	remainingEdges := edges
	numEdgesToFind := len(p.matrix) - 2
	lastPiecePlayed := t
	for i := 0; i < numEdgesToFind; i++ {
		match := findMatchingPiece(lastPiecePlayed, remainingEdges)
		p.matrix[0][i+1] = match // set the edge piece in the puzzle
		lastPiecePlayed = match

		// remove the edge from the list of possible edges
		remainingEdges = removeTileFromSlice(match, remainingEdges)
	}

	// add the top-right corner piece
	match := findMatchingPiece(lastPiecePlayed, corners)
	p.matrix[0][len(p.matrix)-1] = match // set the top-right corner piece in the puzzle
	remainingCorners := removeTileFromSlice(match, corners)

	return remainingEdges, remainingCorners
}

func removeTileFromSlice(t *tile, tiles []*tile) []*tile {
	result := make([]*tile, 0)
	for _, t2 := range tiles {
		if t.id != t2.id {
			result = append(result, t2)
		}
	}
	return result
}

func findMatchingPiece(t *tile, tiles []*tile) *tile {
	for _, tile := range tiles {
		numShared, newTile, _ := numSharedEdges(t, tile)
		if numShared > 0 {
			return newTile
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
// Returns:
// 1) the number of shared edges
// 2) the matching tile, oriented properly
// 3) the ID of HOW the match was made (see compareTiles)
func numSharedEdges(t1 *tile, t2 *tile) (int, *tile, int) {
	if r, id := compareTiles(t1, t2); r > 0 {
		return r, t2, id
	}
	rotated := rotateTile90DegressLeft(t2)
	if r, id := compareTiles(t1, rotated); r > 0 {
		return r, rotated, id
	}
	rotated = rotateTile90DegressLeft(rotated)
	if r, id := compareTiles(t1, rotated); r > 0 {
		return r, rotated, id
	}
	rotated = rotateTile90DegressLeft(rotated)
	if r, id := compareTiles(t1, rotated); r > 0 {
		return r, rotated, id
	}

	flippedT2 := flipTileHorizontal(t2)
	if r, id := compareTiles(t1, flippedT2); r > 0 {
		return r, flippedT2, id
	}
	rotated = rotateTile90DegressLeft(flippedT2)
	if r, id := compareTiles(t1, rotated); r > 0 {
		return r, rotated, id
	}
	rotated = rotateTile90DegressLeft(rotated)
	if r, id := compareTiles(t1, rotated); r > 0 {
		return r, rotated, id
	}
	rotated = rotateTile90DegressLeft(rotated)
	if r, id := compareTiles(t1, rotated); r > 0 {
		return r, rotated, id
	}

	return -1, nil, 0
}

// Returns the number of matching edges
// ALSO returns HOW the edges matched (which edge ID)
// 1 == right edge of 1
// 2 == left edge of 1
// 3 == bottom edge of 1
// 4 == top edge of 1
func compareTiles(t1 *tile, t2 *tile) (int, int) {
	//var numMatching int
	if compare1(t1, t2) {
		return 1, 1
	}
	if compare2(t1, t2) {
		return 1, 2
	}
	if compare3(t1, t2) {
		return 1, 3
	}
	if compare4(t1, t2) {
		return 1, 4
	}

	// // Should only ever be at most 1
	// if numMatching > 1 {
	// 	os.Exit(1)
	// }

	return 0, 0
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
