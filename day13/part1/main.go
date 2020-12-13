package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("/Users/jeff.martin@getweave.com/go/src/github.com/jmartin127/advent-of-code-2020/day13/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	line := scanner.Text()

	// }

	earliest := 1001796
	busIds := []int{37, 41, 457, 13, 17, 23, 29, 431, 19}
	for i := earliest; true; i++ {
		for _, busId := range busIds {
			if i%busId == 0 {
				fmt.Printf("num %d, bus ID %d\n", i, busId)
				os.Exit(1)
			}
		}
	}

}
