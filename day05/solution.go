package day05

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jilleJr/adventofcode-2019/intcode"
	"github.com/jilleJr/adventofcode-2019/util"
)

func step(codes []int, current int) (newCurrent int) {
	var instr, mode1, mode2, _ = intcode.ResolveModes(codes[current])
	switch instr {
	case 1:
		// add
		var (
			a    = intcode.ResolveValue(1, mode1, codes, current, "ADD")
			b    = intcode.ResolveValue(2, mode2, codes, current, "ADD")
			addr = intcode.ResolveAddress(3, codes, current, "ADD")
		)
		codes[addr] = a + b
		return current + 4
	case 2:
		// multiply
		var (
			a    = intcode.ResolveValue(1, mode1, codes, current, "MUL")
			b    = intcode.ResolveValue(2, mode2, codes, current, "MUL")
			addr = intcode.ResolveAddress(3, codes, current, "MUL")
		)
		codes[addr] = a * b
		return current + 4
	case 3:
		// input
		var (
			addr   = intcode.ResolveAddress(1, codes, current, "IN")
			reader = bufio.NewReader(os.Stdin)
		)

		fmt.Printf("%v:IN\tenter input: ", current)
		if input, err := reader.ReadString('\n'); err != nil {
			panic(err)
		} else if n, err := strconv.Atoi(strings.Trim(input, "\r\n\t ")); err != nil {
			panic(err)
		} else {
			codes[addr] = n
		}

		return current + 2
	case 4:
		// output
		var (
			a = intcode.ResolveValue(1, mode1, codes, current, "OUT")
		)

		if a == 0 {
			fmt.Printf("%v:OUT\toutput: 0\ttest passed\n", current)
		} else {
			if codes[current+2] == 99 {
				fmt.Printf("%v:OUT\toutput: %v\n", current, a)
			} else {
				fmt.Println(fmt.Errorf("%v:ERR\toutput: %v\tnon-zero test result", current, a))
			}
		}

		return current + 2

	case 5:
		// jump-if-true
		var (
			cond = intcode.ResolveValue(1, mode1, codes, current, "JMP")
			addr = intcode.ResolveValue(2, mode2, codes, current, "JMP")
		)
		if cond != 0 {
			return addr
		} else {
			return current + 3
		}

	case 6:
		// jump-if-false
		var (
			cond = intcode.ResolveValue(1, mode1, codes, current, "JPF")
			addr = intcode.ResolveValue(2, mode2, codes, current, "JPF")
		)
		if cond == 0 {
			return addr
		} else {
			return current + 3
		}

	case 7:
		// less-than
		var (
			a    = intcode.ResolveValue(1, mode1, codes, current, "LT")
			b    = intcode.ResolveValue(2, mode2, codes, current, "LT")
			addr = intcode.ResolveAddress(3, codes, current, "LT")
		)
		if a < b {
			codes[addr] = 1
		} else {
			codes[addr] = 0
		}
		return current + 4

	case 8:
		// equals
		var (
			a    = intcode.ResolveValue(1, mode1, codes, current, "EQ")
			b    = intcode.ResolveValue(2, mode2, codes, current, "EQ")
			addr = intcode.ResolveAddress(3, codes, current, "EQ")
		)
		if a == b {
			codes[addr] = 1
		} else {
			codes[addr] = 0
		}
		return current + 4

	case 99:
		// end
		fmt.Printf("%v:END\tend of execution reached", current)
		return -1
	default:
		panic(fmt.Errorf("%v:?\tunsupported instruction %v", current, instr))
	}
}

func walk(codes []int) {
	var current int = 0

	for current != -1 {
		current = step(codes, current)
	}
}

func part1(codes []int) {
	walk(codes)
	fmt.Printf("\nresult at index 0: %v\n", codes[0])
}

// Solution of the advent days' pussles
func Solution() {
	if codes, err := util.ReadIntegers("day05/input.txt"); err != nil {
		panic(err)
	} else {
		part1(codes)
	}
}
