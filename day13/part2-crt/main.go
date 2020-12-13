package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

var one = big.NewInt(1)

type busID struct {
	hasValue bool
	id       int
}

func main() {
	file, err := os.Open("/Users/jeff.martin@getweave.com/go/src/github.com/jmartin127/advent-of-code-2020/day13/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var busIDs []busID
	for scanner.Scan() {
		line := scanner.Text()
		busIDs = lineToBusIDs(line)
	}

	// n := []*big.Int{
	// 	big.NewInt(7),
	// 	big.NewInt(13),
	// 	big.NewInt(59),
	// 	big.NewInt(31),
	// 	big.NewInt(19),
	// }
	// a := []*big.Int{
	// 	big.NewInt(0), // 0 --> 7
	// 	big.NewInt(1), // 1 --> 6
	// 	big.NewInt(4), // 4 --> 3
	// 	big.NewInt(6), // 6 --> 1
	// 	big.NewInt(7), // 7 --> 0
	// }

	n, a := generateCrtInput(busIDs)
	answer, err := crt(a, n)
	if err != nil {
		panic(err)
	}
	fmt.Println(answer.Int64() - int64(len(busIDs)-1))
	//fmt.Println(answer.Int64() - int64(busIDs[0].id))
}

func generateCrtInput(ids []busID) ([]*big.Int, []*big.Int) {
	n := make([]*big.Int, 0)
	a := make([]*big.Int, 0)

	max := len(ids) - 1
	for t, id := range ids {
		if id.hasValue {
			n = append(n, big.NewInt(int64(id.id)))
			a = append(a, big.NewInt(int64(max-t)))

			fmt.Printf("Appending %d, %d\n", id.id, max-t)
		} else {
		}
	}

	return n, a
}

func crt(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}

func lineToBusIDs(line string) []busID {
	busIDs := make([]busID, 0)

	vals := strings.Split(line, ",")
	for _, v := range vals {
		var b busID
		if v == "x" {
			b = busID{
				hasValue: false,
			}
		} else {
			id, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			b = busID{
				id:       id,
				hasValue: true,
			}
		}

		busIDs = append(busIDs, b)
	}

	return busIDs
}
