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
