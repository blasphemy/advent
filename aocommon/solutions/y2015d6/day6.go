package y2015d6

import (
	"aoc2/aocommon/solutions"
	"log"
	"strconv"
	"strings"
)

var Solution = solutions.AOCSolution{
	Year: 2015,
	Day: 6,
	Answer1Func: a1,
	Answer2Func: a2,
	DefaultInput: INPUT,
}

type version int

const (
	v1 = iota
	v2 = iota
)

func a1(i string) int {
	lights := &[1000][1000]int{}
	executeCommands(i, lights, v1)
	return countLit(lights)
}

func a2(i string) int {
	lights := &[1000][1000]int{}
	executeCommands(i, lights, v2)
	return countLit(lights)
}

func turnOnRange(lightarray *[1000][1000]int, v version, x1 int, y1 int, x2 int, y2 int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			if v == v1 {
				lightarray[x][y] = 1
			} else {
				lightarray[x][y]++
			}
		}
	}
}

func turnOffRange(lightarray *[1000][1000]int, v version, x1 int, y1 int, x2 int, y2 int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			if v == v1 {
				lightarray[x][y] = 0
			} else {
				lightarray[x][y]--
				if lightarray[x][y] < 0 {
					lightarray[x][y] = 0
				}
			}
		}
	}
}

func toggleRange(lightarray *[1000][1000]int, v version, x1 int, y1 int, x2 int, y2 int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			if v == v1 {
				if lightarray[x][y] == 0 {
					lightarray[x][y] = 1
				} else {
					lightarray[x][y] = 0
				}
			} else {
				lightarray[x][y] += 2
			}
		}
	}
}

func countLit(lightarray *[1000][1000]int) int {
	counter := 0
	for _, x := range lightarray {
		for _, y := range x {
			counter += y
		}
	}
	return counter
}

func executeCommand(in string, lightarray *[1000][1000]int, v version) {
	in = strings.TrimPrefix(in, "turn ")
	split := strings.Split(in, " ")
	x1, y1 := getPointFromString(split[1])
	x2, y2 := getPointFromString(split[3])
	switch split[0] {
	case "on":
		turnOnRange(lightarray, v, x1, y1, x2, y2)
	case "off":
		turnOffRange(lightarray, v, x1, y1, x2, y2)
	case "toggle":
		toggleRange(lightarray, v, x1, y1, x2, y2)
	}
}

func getPointFromString(in string) (int, int) {
	split := strings.Split(in, ",")
	p1, err := strconv.Atoi(split[0])
	if err != nil {
		log.Fatal(err)
	}
	p2, err := strconv.Atoi(split[1])
	if err != nil {
		log.Fatal(err)
	}
	return p1, p2
}

func executeCommands(in string, lightarray *[1000][1000]int, v version) {
	rows := strings.Split(in, "\n")
	for _, x := range rows {
		executeCommand(x, lightarray, v)
	}
}
