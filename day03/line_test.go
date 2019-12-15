package day03

import (
	"kalle/adventofcode/util"
	"testing"
)

func TestIntersectsMatches(t *testing.T) {
	var intersections = [][2]Line{
		[2]Line{MakeLine(10, 10, 10, false), MakeLine(15, 5, 10, true)},
		[2]Line{MakeLine(15, 5, 10, true), MakeLine(10, 10, 10, false)},
		// Test negatives
		[2]Line{MakeLine(20, 10, -10, false), MakeLine(15, 15, -10, true)},
		// Test mix
		[2]Line{MakeLine(20, 10, -10, false), MakeLine(15, 5, 10, true)},
	}
	var want = util.Vector2{X: 15, Y: 10}
	for i, lines := range intersections {
		if got, ok := lines[0].IntersectLine(lines[1]); !ok {
			t.Errorf("expected intersection %v between %v and %v, got none", i, lines[0], lines[1])
		} else if got != want {
			t.Errorf("expected intersection %v between %v and %v to be at %v, got %v", i, lines[0], lines[1], want, got)
		}
	}
}

func TestIntersectsMisses(t *testing.T) {
	var misses = [][2]Line{
		// Overlapping
		[2]Line{MakeLine(10, 10, 10, true), MakeLine(15, 10, 10, true)},
		[2]Line{MakeLine(10, 10, 10, false), MakeLine(10, 5, 10, false)},
		// On the edge
		[2]Line{MakeLine(10, 10, 10, false), MakeLine(10, 10, 10, true)},
		[2]Line{MakeLine(10, 10, 5, false), MakeLine(15, 5, 5, true)},
	}
	for i, lines := range misses {
		if got, ok := lines[0].IntersectLine(lines[1]); ok {
			t.Errorf("expected no intersection on %v between %v and %v, got %v", i, lines[0], lines[1], got)
		}
	}
}

func TestIntersectionSingle(t *testing.T) {
	var (
		horizontalList = LineList{
			Line{pos: util.Vector2{X: 12, Y: 3}, length: 3, vertical: false},
			Line{pos: util.Vector2{X: 0, Y: 0}, length: 10, vertical: false},
			Line{pos: util.Vector2{X: -5, Y: -5}, length: 0, vertical: false},
		}
		verticalLine = Line{pos: util.Vector2{X: 5, Y: -5}, length: 10, vertical: true}
		want         = util.Vector2{X: 5, Y: 0}
	)

	got := horizontalList.Intersections(verticalLine)

	if got == nil || len(got) == 0 {
		t.Fatalf("expected 1 match, got %v", got)
	}

	if len(got) > 1 {
		t.Fatalf("expected 1 match, got %v matches: %v", len(got), got)
	}

	if v := got[0]; v != want {
		t.Fatalf("expected vector %v, got %v", want, v)
	}
}

func TestIntersectPointHorizontal(t *testing.T) {
	var (
		line   = MakeLine(-5, 10, 10, false)
		points = []util.Vector2{
			util.MakeVector2(0, 10),
			util.MakeVector2(-5, 10),
			util.MakeVector2(5, 10),
		}
		wants = []int{
			5,
			0,
			10,
		}
	)

	for i, point := range points {
		if got, ok := line.IntersectPoint(point); !ok {
			t.Errorf("expected %v to intersect %v, but did not", point, line)
		} else if want := wants[i]; got != want {
			t.Errorf("expected %v to intersect %v at distance %v, but got %v", point, line, want, got)
		}
	}
}

func TestIntersectPointVertical(t *testing.T) {
	var (
		line   = MakeLine(10, -5, 10, true)
		points = []util.Vector2{
			util.MakeVector2(10, 0),
			util.MakeVector2(10, -5),
			util.MakeVector2(10, 5),
		}
		wants = []int{
			5,
			0,
			10,
		}
	)

	for i, point := range points {
		if got, ok := line.IntersectPoint(point); !ok {
			t.Errorf("expected %v to intersect %v, but did not", point, line)
		} else if want := wants[i]; got != want {
			t.Errorf("expected %v to intersect %v at distance %v, but got %v", point, line, want, got)
		}
	}
}

func TestIntersectPointNoHorizontal(t *testing.T) {
	var (
		line   = MakeLine(10, 10, 10, false)
		points = []util.Vector2{
			util.MakeVector2(9, 10),
			util.MakeVector2(15, 9),
			util.MakeVector2(15, 11),
			util.MakeVector2(21, 10),
		}
	)

	for _, point := range points {
		if got, ok := line.IntersectPoint(point); ok {
			t.Errorf("expected %v to not intersect %v, but did at %v", point, line, got)
		}
	}
}

func TestIntersectPointNoVertical(t *testing.T) {
	var (
		line   = MakeLine(10, 10, 10, true)
		points = []util.Vector2{
			util.MakeVector2(10, 9),
			util.MakeVector2(9, 15),
			util.MakeVector2(11, 15),
			util.MakeVector2(10, 21),
		}
	)

	for _, point := range points {
		if got, ok := line.IntersectPoint(point); ok {
			t.Errorf("expected %v to not intersect %v, but did at %v", point, line, got)
		}
	}
}

func containsVector2(list []util.Vector2, elem util.Vector2) bool {
	for _, item := range list {
		if item == elem {
			return true
		}
	}
	return false
}

func TestIntersectionMultiple(t *testing.T) {
	var (
		horizontalList = LineList{
			Line{pos: util.Vector2{X: 12, Y: 3}, length: 3, vertical: false},
			Line{pos: util.Vector2{X: 0, Y: 0}, length: 10, vertical: false},
			Line{pos: util.Vector2{X: -3, Y: -3}, length: 10, vertical: false},
		}
		verticalLine = Line{pos: util.Vector2{X: 5, Y: -5}, length: 10, vertical: true}
		want1        = util.Vector2{X: 5, Y: 0}
		want2        = util.Vector2{X: 5, Y: -3}
	)

	got := horizontalList.Intersections(verticalLine)

	if got == nil || len(got) != 2 {
		t.Fatalf("expected 2 matches, got %v: %v", len(got), got)
	}

	if !containsVector2(got, want1) {
		t.Fatalf("expected vector %v to be in list, got %v", want1, got)
	}

	if !containsVector2(got, want2) {
		t.Fatalf("expected vector %v to be in list, got %v", want2, got)
	}
}

func TestIntersectionNone(t *testing.T) {
	var (
		horizontalList = LineList{
			Line{pos: util.Vector2{X: 12, Y: 3}, length: 3, vertical: false},
			Line{pos: util.Vector2{X: 0, Y: 0}, length: 10, vertical: false},
			Line{pos: util.Vector2{X: -3, Y: -3}, length: 10, vertical: false},
		}
		verticalLine = Line{pos: util.Vector2{X: 50, Y: -5}, length: 10, vertical: true}
	)

	got := horizontalList.Intersections(verticalLine)

	if len(got) != 0 {
		t.Fatalf("expected 0 matches, got %v: %v", len(got), got)
	}
}
