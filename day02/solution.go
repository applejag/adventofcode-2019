package day02

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readOpCodes(file string) ([]int, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var builder strings.Builder
	if _, err := builder.Write(b); err != nil {
		return nil, err
	}

	strCodes := strings.Split(strings.TrimRight(builder.String(), "\x00\n\r\t"), ",")
	codes := make([]int, len(strCodes))

	for i, v := range strCodes {
		if op, err := strconv.ParseUint(v, 10, 0); err != nil {
			return nil, err
		} else {
			codes[i] = int(op)
		}
	}

	return codes, nil
}

func getPrintableTargetOpCode(array []int, index int) string {
	if index < 0 || index > len(array) {
		return "?"
	}
	if val := array[index]; val < 0 || val > len(array) {
		return "?"
	} else {
		return strconv.Itoa(val)
	}
}
func getPrintableOpCode(array []int, index int) string {
	if index < 0 || index > len(array) {
		return "?"
	}
	return strconv.Itoa(array[index])
}

func printOpCode(codes []int, i int) {
	switch codes[i] {
	case 1:
		fmt.Printf("[%v]\t%v\tADD [%v] := [%v] + [%v]\t=== %v + %v\t\n", i, codes[i],
			getPrintableOpCode(codes, i+3),
			getPrintableOpCode(codes, i+1),
			getPrintableOpCode(codes, i+2),
			getPrintableTargetOpCode(codes, i+1),
			getPrintableTargetOpCode(codes, i+2))
	case 2:
		fmt.Printf("[%v]\t%v\tMUL [%v] := [%v] * [%v]\t=== %v * %v\t\n", i, codes[i],
			getPrintableOpCode(codes, i+3),
			getPrintableOpCode(codes, i+1),
			getPrintableOpCode(codes, i+2),
			getPrintableTargetOpCode(codes, i+1),
			getPrintableTargetOpCode(codes, i+2))
	case 99:
		fmt.Printf("[%v]\t%v\tHALT\n", i, codes[i])
	default:
		fmt.Printf("[%v]\t%v\n", i, codes[i])
	}
}

func printOpCodes(codes []int, current int) {
	count := len(codes)
	currentOp := codes[current]
	a := codes[current+1]
	b := codes[current+2]
	r := codes[current+3]

	for i := 0; i < count; i++ {
		if i == current {
			fmt.Print("->")
		}
		if i == a || i == b {
			if currentOp == 1 {
				fmt.Print("+")
			} else if currentOp == 2 {
				fmt.Print("*")
			}
		}
		if (currentOp == 1 || currentOp == 2) && i == r {
			fmt.Print("=")
		}

		printOpCode(codes, i)
	}

	fmt.Print("\nCURRENT:\n")
	printOpCode(codes, current)
	fmt.Print("\n")
}

func validateOp(codes []int, current int, op string) (a int, b int, r int) {
	count := len(codes)
	a = codes[current+1]
	b = codes[current+2]
	r = codes[current+3]
	if a < 0 || a >= count {
		panic(fmt.Errorf("%s at %v LHS index out of bounds (%v)", op, current, a))
	}
	if b < 0 || b >= count {
		panic(fmt.Errorf("%s at %v RHS index out of bounds (%v)", op, current, b))
	}
	if r < 0 || r >= count {
		panic(fmt.Errorf("%s at %v result index out of bounds (%v)", op, current, r))
	}
	return
}

func step(codes []int, current int) (newCurrent int) {
	switch codes[current] {
	case 1:
		var a, b, r int = validateOp(codes, current, "ADD")
		codes[r] = codes[a] + codes[b]
		return current + 4
	case 2:
		var a, b, r int = validateOp(codes, current, "MUL")
		codes[r] = codes[a] * codes[b]
		return current + 4
	case 99:
		return -1
	default:
		panic(fmt.Errorf("Unexpected opcode at %v (%v)", current, codes[current]))
	}
}

func walkSlow(codes []int) {
	var current int = 0

	for current != -1 {
		printOpCodes(codes, current)
		fmt.Scanln(new(string))
		current = step(codes, current)
	}
}

func walkFast(codes []int) {
	var current int = 0

	for current != -1 {
		printOpCode(codes, current)
		current = step(codes, current)
	}
}

func walkSilent(codes []int) {
	var current int = 0

	for current != -1 {
		current = step(codes, current)
	}
}

func part1(codes []int) {
	codes[1] = 12
	codes[2] = 2

	walkFast(codes)
	fmt.Printf("\nResult at index 0: %v\n", codes[0])
}

func part2(codes []int) {
	const goal = 19690720

	original := make([]int, len(codes))
	copy(original, codes)

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			codes[1] = noun
			codes[2] = verb
			walkSilent(codes)
			fmt.Printf("verb %v\tnoun %v\t=> %v\n", verb, noun, codes[0])
			if codes[0] == goal {
				fmt.Printf("\nResult %v at verb=%v noun=%v\n", codes[0], verb, noun)
				return
			} else {
				copy(codes, original)
			}
		}
	}

	panic(fmt.Errorf("found no matching combination"))
}

// Solution of the advent days' pussles
func Solution() {
	if codes, err := readOpCodes("input.txt"); err != nil {
		panic(err)
	} else {
		part2(codes)
	}
}
