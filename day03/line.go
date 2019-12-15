package day03

import "kalle/adventofcode/util"

import "fmt"

// Line is a 2D point with a length
type Line struct {
	pos      util.Vector2
	length   int
	vertical bool
}

// LineList is a collection of lines
type LineList []Line

// MakeLine returns a new line
func MakeLine(x, y, length int, vertical bool) Line {
	return Line{pos: util.Vector2{X: x, Y: y}, length: length, vertical: vertical}
}

func minMax(a, b int) (min, max int) {
	if a > b {
		return b, a
	}
	return a, b
}

// IntersectLine returns true if the two lines intersects and where
func (line Line) IntersectLine(other Line) (util.Vector2, bool) {
	if line.vertical == other.vertical {
		return util.Vector2{}, false
	}
	if line.vertical {
		// line.vertical = true
		// other.vertical = false
		var otherMinX, otherMaxX = minMax(other.pos.X+other.length, other.pos.X)
		if line.pos.X >= otherMaxX ||
			line.pos.X <= otherMinX {
			return util.Vector2{}, false
		}
		var lineMinY, lineMaxY = minMax(line.pos.Y+line.length, line.pos.Y)
		if lineMaxY < other.pos.Y ||
			lineMinY > other.pos.Y {
			return util.Vector2{}, false
		}
		return util.Vector2{X: line.pos.X, Y: other.pos.Y}, true
	}
	return other.IntersectLine(line)
}

func (line Line) String() string {
	switch {
	case line.vertical && line.length < 0:
		return fmt.Sprintf("{%v, %v-%v}", line.pos.X, line.pos.Y, -line.length)
	case line.vertical && line.length > 0:
		return fmt.Sprintf("{%v, %v+%v}", line.pos.X, line.pos.Y, line.length)
	case !line.vertical && line.length > 0:
		return fmt.Sprintf("{%v+%v, %v}", line.pos.X, line.length, line.pos.Y)
	case !line.vertical && line.length < 0:
		return fmt.Sprintf("{%v-%v, %v}", line.pos.X, -line.length, line.pos.Y)
	default:
		return fmt.Sprintf("{%v, %v}", line.pos.X, line.pos.Y)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

// IntersectPoint checks wether a line crosses a point, and if so at what distance
func (line Line) IntersectPoint(point util.Vector2) (distance int, ok bool) {
	if line.vertical {
		if line.pos.X != point.X {
			return 0, false
		}

		minY, maxY := minMax(line.pos.Y, line.pos.Y+line.length)
		if point.Y >= minY && point.Y <= maxY {
			return abs(point.Y - line.pos.Y), true
		}
	} else {
		if line.pos.Y != point.Y {
			return 0, false
		}

		minX, maxX := minMax(line.pos.X, line.pos.X+line.length)
		if point.X >= minX && point.X <= maxX {
			return abs(point.X - line.pos.X), true
		}
	}

	return 0, false
}

// Intersections returns the points of any intersecting lines.
func (c LineList) Intersections(a Line) []util.Vector2 {
	var intersections []util.Vector2

	for _, b := range c {
		if i, ok := a.IntersectLine(b); ok {
			intersections = append(intersections, i)
		}
	}

	return intersections
}
