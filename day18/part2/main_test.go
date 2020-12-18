package main

import (
	"fmt"
	"testing"
)

// 1 + 2 * 3 + 4 * 5 + 6
func TestReducePlus(t *testing.T) {
	tokens := []*token{
		{
			value: 1,
		},
		{
			isOperator: true,
			operator:   "+",
		},
		{
			value: 2,
		},
		{
			isOperator: true,
			operator:   "*",
		},
		{
			value: 3,
		},
		{
			isOperator: true,
			operator:   "+",
		},
		{
			value: 4,
		},
		{
			isOperator: true,
			operator:   "*",
		},
		{
			value: 5,
		},
		{
			isOperator: true,
			operator:   "+",
		},
		{
			value: 6,
		},
	}
	result := reducePlus(tokens)

	for _, t := range result {
		fmt.Printf("token %+v\n", t)
	}

	t.Fatal()
}
