package day03

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"github.com/jilleJr/adventofcode-2019/turtle"
	"github.com/jilleJr/adventofcode-2019/util"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const imageOriginX, imageOriginY = 1024., 1024.
const imageScale = 0.1
const sqrt2 = 1.41421356237

func parseStep(step string) Step {
	direction := rune(step[0])
	lenString := string(step[1:])
	if len, err := strconv.ParseUint(lenString, 10, 0); err != nil {
		panic(err)
	} else {
		return Step{
			direction,
			int(len),
		}
	}
}

func parseLines(array []string) [][]Step {
	var output [2][]Step
	count := 0

	for _, v := range array {
		if v != "" {
			if count == 2 {
				panic(fmt.Errorf("too many instruction lines"))
			}
			line := strings.Split(v, ",")
			steps := make([]Step, len(line))
			for i, s := range line {
				steps[i] = parseStep(s)
			}
			output[count] = steps
			count++
		}
	}

	return output[:2]
}

func readPaths(file string) ([][]Step, error) {
	barr, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var builder strings.Builder

	for _, b := range barr {
		if b == '\n' || unicode.IsPrint(rune(b)) {
			err := builder.WriteByte(b)
			if err != nil {
				return nil, err
			}
		}
	}

	return parseLines(strings.Split(builder.String(), "\n")), nil
}

func walk(pos util.Vector2, step Step) util.Vector2 {
	return pos.Add(step.Vector2())
}

func drawLines(t *turtle.Turtle, lines LineList) {
	t.UsePattern = true
	t.PenDown()
	for _, line := range lines {
		positive := line.length > 0
		switch {
		case !line.vertical && positive:
			t.LookRight()
		case line.vertical && positive:
			t.LookUp()
		case !line.vertical && !positive:
			t.LookLeft()
		case line.vertical && !positive:
			t.LookDown()
		}
		t.Forward(math.Abs(float64(line.length)) * imageScale)
	}
}

func drawLinesUntilIntersection(t *turtle.Turtle, lines LineList, inter Intersection) {
	t.UsePattern = true
	t.PenDown()
	pos := inter.Pos
	for _, line := range lines {
		positive := line.length > 0
		switch {
		case !line.vertical && positive:
			t.LookRight()
		case line.vertical && positive:
			t.LookUp()
		case !line.vertical && !positive:
			t.LookLeft()
		case line.vertical && !positive:
			t.LookDown()
		}

		if dist, ok := line.IntersectPoint(pos); ok {
			t.Forward(float64(dist) * imageScale)
			break
		} else {
			t.Forward(math.Abs(float64(line.length)) * imageScale)
		}
	}
}

func walkSet(set []Step) LineList {
	var (
		pos   util.Vector2
		lines []Line
	)

	for _, step := range set {
		line := Line{
			vertical: step.direction == 'U' || step.direction == 'D',
			length:   step.length,
			pos:      pos,
		}

		if step.direction == 'L' || step.direction == 'D' {
			line.length = -line.length
		}

		lines = append(lines, line)
		pos = walk(pos, step)
	}

	return lines
}

func distanceToIntersection(lines LineList, point util.Vector2) int {
	var traveled = 0

	for _, line := range lines {
		if dist, ok := line.IntersectPoint(point); ok {
			return traveled + dist
		}

		if line.length < 0 {
			traveled -= line.length
		} else {
			traveled += line.length
		}
	}

	panic(fmt.Errorf("list of lines expected to intersect with %v, but did not", point))
}

func scoreIntersections(list1, list2 LineList, points []util.Vector2) []Intersection {
	var scored []Intersection

	for _, point := range points {
		var list1Dist = distanceToIntersection(list1, point)
		var list2Dist = distanceToIntersection(list2, point)
		fmt.Printf("to %v, dist is %v+%v = %v\n", point, list1Dist, list2Dist, list1Dist+list2Dist)
		scored = append(scored, Intersection{Pos: point, Score: list1Dist + list2Dist})
	}

	return scored
}

func drawCrossAtVector(t *turtle.Turtle, pos util.Vector2) {
	const crossSide = 10.
	const crossSideDiv2 = crossSide * .5
	const crossDiagonal = crossSide * sqrt2

	t.UsePattern = false

	t.Pos = turtle.Position{X: float64(pos.X)*imageScale - crossSideDiv2 + imageOriginX, Y: float64(pos.Y)*imageScale + crossSideDiv2 + imageOriginY}
	t.Orientation = math.Pi * 7 / 4
	t.PenDown()
	t.Forward(crossDiagonal)
	t.PenUp()

	t.Pos = turtle.Position{X: float64(pos.X)*imageScale - crossSideDiv2 + imageOriginX, Y: float64(pos.Y)*imageScale - crossSideDiv2 + imageOriginY}
	t.Orientation = math.Pi / 4
	t.PenDown()
	t.Forward(crossDiagonal)
	t.PenUp()
}

func newImage() *image.RGBA {
	const width = 2048
	const height = 2048
	img := image.NewRGBA(image.Rect(0, 0, 2048, 2048))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.Black)
		}
	}
	return img
}

func part1(sets [][]Step) {
	img := newImage()
	t := turtle.NewTurtle(img, turtle.Position{X: imageOriginX, Y: imageOriginY})
	t.PenDown()

	lines1 := walkSet(sets[0])
	lines2 := walkSet(sets[1])

	t.Color = color.RGBA{G: 0xff, A: 0xff}
	drawLines(t, lines1)
	t.Pos = turtle.Position{X: imageOriginX, Y: imageOriginY}
	t.Color = color.RGBA{B: 0xff, A: 0xff}
	drawLines(t, lines2)

	t.Color = color.RGBA{R: 0xff, A: 0xff}

	var (
		closest     *util.Vector2
		closestDist int
	)

	t.Color = color.RGBA{R: 0xff, A: 0xff}
	for _, line := range lines2 {
		intersections := lines1.Intersections(line)

		for _, i := range intersections {
			if dist := i.ManhattanDist(); closest == nil || dist < closestDist {
				closest = &i
				closestDist = dist
			}
			drawCrossAtVector(t, i)
		}
	}

	if closest == nil {
		fmt.Println("did not find a single intersection")
	} else {
		fmt.Printf("closest intersection: %v\n", *closest)
		fmt.Printf("manhattan distance of closest: %v\n", closestDist)
	}

	t.Color = color.RGBA{R: 0xff, G: 0xff, A: 0xff}
	drawCrossAtVector(t, util.Vector2{})

	file, _ := os.Create("day03/img.png")
	defer file.Close()

	png.Encode(file, img)
}

func part2(sets [][]Step) {
	img := newImage()
	t := turtle.NewTurtle(img, turtle.Position{X: imageOriginX, Y: imageOriginY})

	lines1 := walkSet(sets[0])
	lines2 := walkSet(sets[1])

	t.Color = color.RGBA{G: 0xff, A: 0xff}
	drawLines(t, lines1)
	t.Pos = turtle.Position{X: imageOriginX, Y: imageOriginY}
	t.Color = color.RGBA{B: 0xff, A: 0xff}
	drawLines(t, lines2)

	t.Color = color.RGBA{R: 0xff, A: 0xff}
	var (
		points []util.Vector2
	)

	t.Color = color.RGBA{R: 0xff, A: 0xff}

	for _, line := range lines2 {
		points = append(points, lines1.Intersections(line)...)

		for _, i := range points {
			drawCrossAtVector(t, i)
		}
	}

	if len(points) == 0 {
		panic(fmt.Errorf("did not find a single intersection"))
	} else {
		fmt.Printf("found %v intersections\n", len(points))
	}

	var (
		intersections  = scoreIntersections(lines1, lines2, points)
		lowestScore    Intersection
		anyLowestScore = false
	)
	for _, inter := range intersections {
		if !anyLowestScore || inter.Score < lowestScore.Score {
			lowestScore = inter
			anyLowestScore = true
		}
	}

	fmt.Printf("intersection with lowest score: %v\n", lowestScore)

	t.Pos = turtle.Position{X: imageOriginX, Y: imageOriginY}
	t.Color = color.RGBA{R: 0xff, G: 0x95, A: 0xff}
	drawLinesUntilIntersection(t, lines1, lowestScore)
	t.Pos = turtle.Position{X: imageOriginX, Y: imageOriginY}
	t.Color = color.RGBA{R: 0xff, B: 0x95, A: 0xff}
	drawLinesUntilIntersection(t, lines2, lowestScore)

	t.Color = color.RGBA{R: 0xff, G: 0xff, A: 0xff}
	drawCrossAtVector(t, util.Vector2{})

	file, _ := os.Create("day03/img.png")
	defer file.Close()

	png.Encode(file, img)
}

// Solution of the advent days' pussles
func Solution() {
	if sets, err := readPaths("day03/input.txt"); err != nil {
		panic(err)
	} else {
		part2(sets)
	}
}
