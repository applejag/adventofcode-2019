package day04

import (
	"fmt"
	"strconv"
)

const inputLower = 152085
const inputUpper = 670283

func part1(password int) bool {
	str := fmt.Sprint(password)
	var lastN int
	var twoAdjacent = false

	for i, r := range str {
		n, err := strconv.Atoi(string(r))
		if err != nil {
			panic(fmt.Errorf("unable to parse '%v' as number in password '%v'", r, password))
		}

		// all digits must increase
		if i > 0 && n < lastN {
			return false
		}

		// two digits must follow each other some time
		if i > 0 && rune(str[i-1]) == r {
			twoAdjacent = true
		}

		lastN = n
	}

	if !twoAdjacent {
		return false
	}

	return true
}

func part2(password int) bool {
	str := fmt.Sprint(password)
	var lastN int
	var twoAdjacent = false
	var adjacent = 0
	var onlyTwoAdjacent = false

	for i, r := range str {
		n, err := strconv.Atoi(string(r))
		if err != nil {
			panic(fmt.Errorf("unable to parse '%v' as number in password '%v'", r, password))
		}

		// all digits must increase
		if i > 0 && n < lastN {
			return false
		}

		// two digits must follow each other some time
		if i > 0 && rune(str[i-1]) == r {
			twoAdjacent = true
			adjacent++
		} else {
			if adjacent == 2 {
				onlyTwoAdjacent = true
			}
			adjacent = 1
		}

		lastN = n
	}

	if !twoAdjacent {
		return false
	}

	// at least one group must have reached only two adjacent
	if !onlyTwoAdjacent && adjacent != 2 {
		return false
	}

	return true
}

// Solution of the advent days' pussles
func Solution() {
	var count = 0

	for password := inputLower; password <= inputUpper; password++ {
		var valid = part2(password)

		if valid {
			count++
			fmt.Printf("valid password: %v\n", password)
		}
	}

	fmt.Printf("%v valid passwords\n", count)
}
