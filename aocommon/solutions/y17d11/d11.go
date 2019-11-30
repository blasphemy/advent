package y17d11

import (
	"advent/aocommon/solutions"
	"strconv"
	"strings"
)

var Solution = solutions.AOCSolution{
	Year:         17,
	Day:          11,
	DefaultInput: input,
	Answer1Func:  a1,
	Answer2Func:  a2,
}

func a1(i string) string {
	ans, _ := processInstructions(i)
	return strconv.Itoa(ans)
}

func a2(i string) string {
	_, ans := processInstructions(i)
	return strconv.Itoa(ans)
}

type position struct {
	x int
	y int
}

func processInstructions(in string) (int, int) {
	max := 0
	pos := position{0, 0}
	split := strings.Split(in, ",")
	for _, x := range split {
		switch x {
		case "nw":
			pos.y++
			pos.x--
		case "ne":
			pos.x++
		case "se":
			pos.y--
			pos.x++
		case "sw":
			pos.x--
		case "s":
			pos.y--
		case "n":
			pos.y++
		}
		d := getDist(pos)
		if d > max {
			max = d
		}
	}
	return getDist(pos), max
}

func abs(in int) int {
	if in > 0 {
		return in
	}
	return -in
}

func getDist(in position) int {
	return (abs(in.x) + abs(in.y) + abs(in.x+in.y)) / 2
}
