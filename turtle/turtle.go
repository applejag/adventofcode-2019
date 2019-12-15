package turtle

import (
	"image"
	"image/color"
	"math"
)

type Turtle struct {
	Pos         Position
	Orientation float64
	Color       color.Color
	IsPenDown   bool
	UsePattern  bool

	patternIndex byte
	image        *image.RGBA
	size         image.Point
}

type Position struct {
	X, Y float64
}

const patternLength = 6

var pattern = [patternLength]bool{
	false,
	true,
	false,
	true,
	true,
	false,
}

func NewTurtle(i *image.RGBA, starting Position) (t *Turtle) {
	t = &Turtle{
		image:       i,
		Pos:         starting,
		Orientation: 0.0,
		Color:       color.Black,
		IsPenDown:   true,
		UsePattern:  true,
		size:        i.Bounds().Size(),
	}

	return
}

func (t *Turtle) Forward(dist float64) {
	if t.IsPenDown {
		x := t.Pos.X
		y := t.Pos.Y
		dx := math.Cos(t.Orientation)
		dy := math.Sin(t.Orientation)

		for i := 0.; i < dist; i++ {
			if t.UsePattern {
				t.patternIndex = (t.patternIndex + 1) % patternLength
				if pattern[t.patternIndex] {
					// +0.5 for rounding
					t.image.Set(int(x+dx*i+0.5), t.size.Y-int(y+dy*i+0.5), t.Color)
				}
			} else {
				// +0.5 for rounding
				t.image.Set(int(x+dx*i+0.5), t.size.Y-int(y+dy*i+0.5), t.Color)
			}
		}
	}

	x := dist * math.Cos(t.Orientation)
	y := dist * math.Sin(t.Orientation)
	t.Pos = Position{t.Pos.X + x, t.Pos.Y + y}
}

func (t *Turtle) LookRight() {
	t.Orientation = 0
}

func (t *Turtle) LookUp() {
	t.Orientation = math.Pi * 0.5
}

func (t *Turtle) LookLeft() {
	t.Orientation = math.Pi
}

func (t *Turtle) LookDown() {
	t.Orientation = math.Pi * 1.5
}

func (t *Turtle) LookDegrees(degrees float64) {
	const degToRad = 180. / math.Pi
	t.Orientation = degrees * degToRad
}

func (t *Turtle) Turn(radians float64) {
	t.Orientation += radians
}

func (t *Turtle) PenUp() {
	t.IsPenDown = false
}

func (t *Turtle) PenDown() {
	t.IsPenDown = true
	t.patternIndex = 0
}
