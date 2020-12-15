package main

import "fmt"

func main() {
	input := []int{19, 0, 5, 1, 10, 13}
	turnLastSpoken := make(map[int][]int, 0)
	for i, v := range input {
		turnLastSpoken[v] = []int{i + 1}
	}

	lastNum := 13
	for turn := len(input) + 1; turn <= 30000000; turn++ {
		var newNum int
		if turns, ok := turnLastSpoken[lastNum]; !ok || len(turns) <= 1 {
			newNum = 0
		} else {
			turns := turnLastSpoken[lastNum]
			last := turns[len(turns)-1]
			secondLast := turns[len(turns)-2]
			diff := last - secondLast
			newNum = diff
		}

		// add the new one
		lastNum = newNum
		if t, ok := turnLastSpoken[newNum]; ok {
			t = append(t, turn)
			turnLastSpoken[newNum] = t
		} else {
			turnLastSpoken[newNum] = []int{turn}
		}
	}

	fmt.Printf("Answer %d\n", lastNum)
}
