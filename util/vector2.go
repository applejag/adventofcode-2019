package util

// Vector2 is a location in 2D space
type Vector2 struct {
	X, Y int
}

func MakeVector2(x, y int) Vector2 {
	return Vector2{X: x, Y: y}
}

// Add adds two vectors together, returning the result
func (v Vector2) Add(b Vector2) Vector2 {
	return Vector2{v.X + b.X, v.Y + b.Y}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// ManhattanDist returns the manhattan distance of a Vector2
func (v Vector2) ManhattanDist() int {
	return abs(v.X) + abs(v.Y)
}
