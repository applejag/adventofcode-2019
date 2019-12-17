package intcode

import (
	"fmt"
	"strconv"
)

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

func PrintOpCode(codes []int, i int) {
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

func PrintOpCodes(codes []int, current int) {
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

		PrintOpCode(codes, i)
	}

	fmt.Print("\nCURRENT:\n")
	PrintOpCode(codes, current)
	fmt.Print("\n")
}

func ValidateBinaryOp(codes []int, current int, op string) (a int, b int, r int) {
	a = ResolveAddress(1, codes, current, op)
	b = ResolveAddress(2, codes, current, op)
	r = ResolveAddress(3, codes, current, op)
	return
}

func ResolveValue(param int, mode int, codes []int, current int, op string) int {
	count := len(codes)

	if param < 0 || param >= count {
		panic(fmt.Errorf("%v:%s\tparameter %v: parameter index out of bounds", current, op, param))
	}

	switch mode {
	case 0: // position mode
		a := codes[current+param]
		if a < 0 || a >= count {
			panic(fmt.Errorf("%v:%s\tparameter %v: postition mode out of bounds (%v)", current, op, param, a))
		}
		return codes[a]

	case 1: // immediate mode
		return codes[current+param]

	default:
		panic(fmt.Errorf("%v:%s\tunknown intcode parameter mode '%v'", current, op, mode))
	}
}

func ResolveAddress(param int, codes []int, current int, op string) int {
	count := len(codes)

	if param < 0 || param >= count {
		panic(fmt.Errorf("%v:%s\tparameter %v: parameter index out of bounds", current, op, param))
	}

	// position mode
	a := codes[current+param]
	if a < 0 || a >= count {
		panic(fmt.Errorf("%v:%s\tparameter %v: postition mode out of bounds (%v)", current, op, param, a))
	}
	return a
}

func ResolveModes(code int) (instr, mode1, mode2, mode3 int) {
	instr = code % 100
	mode1 = (code / 100) % 10
	mode2 = (code / 1000) % 10
	mode3 = (code / 1000) % 10
	return
}
