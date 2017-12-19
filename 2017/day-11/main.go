package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type position struct {
	x int
	y int
}

func main() {
	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	answer1, answer2 := processInstructions(string(inputBytes))
	fmt.Printf("Part 1: %d\nPart 2: %d\n", answer1, answer2)
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
