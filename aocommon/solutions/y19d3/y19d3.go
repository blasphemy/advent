package y19d3

import (
	"advent/aocommon/solutions"
	"strconv"
	"strings"
)

// types
type dir int

type place struct {
	point
	step int
}

type wire struct {
	point
	places  []place
	counter int
}

type point struct {
	x int
	y int
}

type instruction struct {
	direction dir
	amount    int
}

type instructionSet struct {
	instructions []instruction
}

const (
	UP    = iota
	DOWN  = iota
	LEFT  = iota
	RIGHT = iota
)

var Solution = solutions.AOCSolution{
	Year:         19,
	Day:          3,
	DefaultInput: input,
	Answer1Func:  a1,
	Answer2Func:  a2,
}

func a1(i string) string {
	instructions, err := parseInstructions(i)
	if err != nil {
		panic(err)
	}
	w1 := &wire{point{0, 0}, []place{}, 0}
	w2 := &wire{point{0, 0}, []place{}, 0}
	w1.runInstructionSet(instructions[0])
	w2.runInstructionSet(instructions[1])
	intersections := findIntersections(w1, w2)
	n := -1
	for _, x := range intersections {
		d := getDistance(point{0, 0}, x.point)
		if d < n || n == -1 {
			n = d
		}
	}
	return strconv.Itoa(n)
}

func a2(i string) string {
	instructions, err := parseInstructions(i)
	if err != nil {
		panic(err)
	}
	w1 := &wire{point{0, 0}, []place{}, 0}
	w2 := &wire{point{0, 0}, []place{}, 0}
	w1.runInstructionSet(instructions[0])
	w2.runInstructionSet(instructions[1])
	intersections := findIntersections(w1, w2)
	n := -1
	for _, x := range intersections {
		if x.step < n || n == -1 {
			n = x.step
		}
	}
	return strconv.Itoa(n)
}

func getDistance(p1, p2 point) int {
	a := abs(p1.x - p2.x)
	b := abs(p1.y - p2.y)
	return a + b
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func parseInstructions(i string) ([]instructionSet, error) {
	r := strings.Split(i, "\n")
	out := []instructionSet{}
	for _, x := range r {
		is, err := parseLine(x)
		if err != nil {
			return []instructionSet{}, err
		}
		out = append(out, is)
	}
	return out, nil
}

func parseLine(i string) (instructionSet, error) {
	s := strings.Split(i, ",")
	out := instructionSet{
		instructions: []instruction{},
	}
	for _, x := range s {
		var d dir
		switch x[0] {
		case 'R':
			d = RIGHT
		case 'D':
			d = DOWN
		case 'L':
			d = LEFT
		}
		ns := strings.TrimLeft(x, "RDLU")
		n, err := strconv.Atoi(ns)
		if err != nil {
			return instructionSet{}, err
		}
		w := instruction{
			direction: d,
			amount:    n,
		}
		out.instructions = append(out.instructions, w)
	}
	return out, nil
}

func findIntersections(w1, w2 *wire) []place {
	out := []place{}
	for _, x := range w1.places {
		for _, y := range w2.places {
			if x.point == y.point {
				combinedPoint := x
				combinedPoint.step = x.step + y.step
				out = append(out, combinedPoint)
			}
		}
	}
	return out
}

func (w *wire) runInstructionSet(is instructionSet) {
	for _, x := range is.instructions {
		w.processInstruction(x)
	}
}

func (w *wire) processInstruction(i instruction) {
	var o *int
	n := false
	switch i.direction {
	case RIGHT:
		o = &w.x
	case LEFT:
		o = &w.x
		n = true
	case UP:
		o = &w.y
	case DOWN:
		o = &w.y
		n = true
	}
	for k := 0; k < i.amount; k++ {
		if n {
			*o--
		} else {
			*o++
		}
		w.recordPoint()
	}
}

func (w *wire) recordPoint() {
	w.counter++
	w.places = append(w.places, place{w.point, w.counter})
}