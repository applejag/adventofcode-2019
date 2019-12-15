package day0

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readLines(file string) ([]string, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var builder strings.Builder
	if _, err := builder.Write(b); err != nil {
		return nil, err
	}

	lines := strings.Split(builder.String(), "\n")
	return lines, nil
}

func calcFuelForMass(mass int) int {
	fuel := mass/3 - 2
	fmt.Printf("%v\t=> %v\n", mass, fuel)
	return fuel
}

func _calcFuelForMassRecursive(mass int, sum int) int {
	fuel := mass/3 - 2

	if fuel <= 0 {
		return 0
	}

	if sum == 0 {
		fmt.Print(fuel)
	} else {
		fmt.Printf(", %v", fuel)
	}

	return fuel + _calcFuelForMassRecursive(fuel, sum+fuel)
}

func calcFuelForMassRecursive(mass int) int {
	fmt.Printf("%v\t=> ", mass)
	fuel := _calcFuelForMassRecursive(mass, 0)
	fmt.Printf(" => %v\n", fuel)

	return fuel
}

// Solution of the advent days' pussles
func Solution() {
	var sum int

	if lines, err := readLines("input.txt"); err != nil {
		panic(err)
	} else {
		for _, line := range lines {
			if len(line) == 0 {
				continue
			}

			if mass, err := strconv.Atoi(line); err != nil {
				panic(err)
			} else {
				sum += calcFuelForMassRecursive(mass)
			}
		}
	}

	fmt.Printf("Sum = %v\n", sum)
}
