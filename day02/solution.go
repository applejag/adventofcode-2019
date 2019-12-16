package day02

import (
	"fmt"

	"github.com/jilleJr/adventofcode-2019/intcode"
	"github.com/jilleJr/adventofcode-2019/util"
)

func step(codes []int, current int) (newCurrent int) {
	switch codes[current] {
	case 1:
		var a, b, r int = intcode.ValidateBinaryOp(codes, current, "ADD")
		codes[r] = codes[a] + codes[b]
		return current + 4
	case 2:
		var a, b, r int = intcode.ValidateBinaryOp(codes, current, "MUL")
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
		intcode.PrintOpCodes(codes, current)
		fmt.Scanln(new(string))
		current = step(codes, current)
	}
}

func walkFast(codes []int) {
	var current int = 0

	for current != -1 {
		intcode.PrintOpCode(codes, current)
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
	if codes, err := util.ReadIntegers("day02/input.txt"); err != nil {
		panic(err)
	} else {
		part2(codes)
	}
}
