package util

import (
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
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

func integerFieldsFunc(r rune) bool {
	return !unicode.IsDigit(r)
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

	strCodes := strings.FieldsFunc(builder.String(), integerFieldsFunc)
	codes := make([]int, len(strCodes))

	for i, v := range strCodes {
		if op, err := strconv.ParseInt(v, 10, 0); err != nil {
			return nil, err
		} else {
			codes[i] = int(op)
		}
	}

	return codes, nil
}
