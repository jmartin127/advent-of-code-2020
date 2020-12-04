package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type passport struct {
	contents []field
}

type field struct {
	key   string
	value string
}

var validKeys = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func main() {
	file, err := os.Open("/Users/jeff.martin@getweave.com/go/src/github.com/jmartin127/advent-of-code-2020/day4/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	passports := make([]passport, 0)
	scanner := bufio.NewScanner(file)
	currentPassport := passport{
		contents: make([]field, 0),
	}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			passports = append(passports, currentPassport)
			currentPassport = passport{
				contents: make([]field, 0),
			}
		} else {
			newFields := retrieveFieldsFromLine(line)
			currentPassport.contents = append(currentPassport.contents, newFields...)
		}
	}
	passports = append(passports, currentPassport)

	var numValid int
	for _, passport := range passports {
		fmt.Printf("Passport %+v\n", passport)
		if passport.isValid() {
			numValid++
		}
	}
	fmt.Printf("Valid %d\n", numValid)
}

func (p *passport) isValid() bool {
	var numValid int
	for _, validKey := range validKeys {
		for _, field := range p.contents {
			if field.key == validKey {
				numValid++
			}
		}
	}
	if numValid >= 7 {
		return true
	}
	return false
}

func retrieveFieldsFromLine(line string) []field {
	fields := make([]field, 0)
	pairs := strings.Split(line, " ")
	for _, pair := range pairs {
		values := strings.Split(pair, ":")
		field := field{
			key:   values[0],
			value: values[1],
		}
		fields = append(fields, field)
	}
	return fields
}
