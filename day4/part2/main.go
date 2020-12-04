package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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
	for _, validKey := range validKeys {
		value := p.getFieldValue(validKey)
		if value == "" {
			return false
		}

		fmt.Printf("Checking value %s for key %s\n", value, validKey)

		if validKey == "byr" {
			i, err := strconv.Atoi(value)
			if err != nil {
				return false
			}
			if i < 1920 || i > 2002 {
				return false
			}
		} else if validKey == "iyr" {
			i, err := strconv.Atoi(value)
			if err != nil {
				return false
			}
			if i < 2010 || i > 2020 {
				return false
			}
		} else if validKey == "eyr" {
			i, err := strconv.Atoi(value)
			if err != nil {
				return false
			}
			if i < 2020 || i > 2030 {
				return false
			}
		} else if validKey == "hgt" {
			if !strings.Contains(value, "cm") && !strings.Contains(value, "in") {
				return false
			}

			begin := value[:len(value)-2]
			end := value[len(value)-2:]

			fmt.Printf("Begin %s end %s\n", begin, end)

			num, err := strconv.Atoi(begin)
			if err != nil {
				return false
			}
			if end == "cm" && (num < 150 || num > 193) {
				return false
			}
			if end == "in" && (num < 59 || num > 76) {
				return false
			}
		} else if validKey == "hcl" {
			matched, err := regexp.MatchString(`^#[0-9abcdef]{6}$`, value)
			if err != nil {
				return false
			}
			if !matched {
				return false
			}
			fmt.Printf("matched hcl for value %s\n", value)
		} else if validKey == "ecl" {
			if value != "amb" && value != "blu" && value != "brn" && value != "gry" && value != "grn" && value != "hzl" && value != "oth" {
				return false
			}
		} else if validKey == "pid" {
			if len(value) != 9 {
				return false
			}
			_, err := strconv.Atoi(value)
			if err != nil {
				return false
			}
		}
	}
	return true
}

func (p *passport) getFieldValue(key string) string {
	for _, field := range p.contents {
		if field.key == key {
			return field.value
		}
	}

	return ""
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
