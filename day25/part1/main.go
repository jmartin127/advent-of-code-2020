package main

import "fmt"

func main() {
	cardPubKey := 5764801
	doorPubKey := 17807724
	// cardPubKey := 335121
	// doorPubKey := 363891

	cardLoopSize := findLoopSize(cardPubKey)
	fmt.Printf("loop size %d\n", cardLoopSize)
	encryptionKey := transformSubjectNumber(doorPubKey, cardLoopSize)
	fmt.Printf("Answer %d\n", encryptionKey)
}

func findLoopSize(pubKey int) int {
	for i := 0; i < 10000000; i++ {
		result := transformSubjectNumber(7, i+1)
		if result == pubKey {
			return i + 1
		}
	}

	return -1
}

/*
The handshake used by the card and the door involves an operation that transforms a subject number. To transform a subject number, start with the value 1. Then, a number of times called the loop size, perform the following steps:

Set the value to itself multiplied by the subject number.
Set the value to the remainder after dividing the value by 20201227.
*/
func transformSubjectNumber(subNum int, loopSize int) int {
	value := 1
	for i := 0; i < loopSize; i++ {
		value = value * subNum
		value = value % 20201227
	}

	return value
}
