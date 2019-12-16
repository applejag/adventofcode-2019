package util

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// ReadLines returns content of a file split into lines
func ReadLines(file string) ([]string, error) {
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

// ReadIntegers returns array of integers from a file of comma separated numbers
func ReadIntegers(file string) ([]int, error) {
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
