package day03

import "kalle/adventofcode/util"

// Step is one segment in a path
type Step struct {
	direction rune
	length    int
}

// Vector2 creates a vector of direction and length of Step
func (s Step) Vector2() util.Vector2 {
	switch s.direction {
	case 'R':
		return util.Vector2{s.length, 0}
	case 'L':
		return util.Vector2{-s.length, 0}
	case 'U':
		return util.Vector2{0, s.length}
	case 'D':
		return util.Vector2{0, -s.length}
	default:
		return util.Vector2{0, 0}
	}
}
