package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	challengeInput := string(inputBytes)
	answer1 := getUniqueLocationsFromInstructions(challengeInput)
	answer2 := getUniqueLocationsFromInstructionsRobo(challengeInput)
	fmt.Printf("Part 1: %d\n", answer1)
	fmt.Printf("Part 2: %d\n", answer2)
}

type coard struct {
	x int
	y int
}

func getUniqueLocationsFromInstructions(input string) int {
	locations := make(map[coard]bool)
	currentCoard := coard{0, 0}
	locations[currentCoard] = true
	directions := strings.Split(input, "")
	for _, x := range directions {
		switch x {
		case ">":
			currentCoard.x++
		case "<":
			currentCoard.x--
		case "^":
			currentCoard.y++
		case "v":
			currentCoard.y--
		}
		locations[currentCoard] = true
	}
	return len(locations)
}

func getUniqueLocationsFromInstructionsRobo(input string) int {
	locations := make(map[coard]bool)
	santaCoard := &coard{0, 0}
	roboCoard := &coard{0, 0}
	locations[*santaCoard] = true
	directions := strings.Split(input, "")
	for y, x := range directions {
		var currentCoard *coard
		if y%2 == 0 {
			currentCoard = santaCoard
		} else {
			currentCoard = roboCoard
		}
		switch x {
		case ">":
			currentCoard.x++
		case "<":
			currentCoard.x--
		case "^":
			currentCoard.y++
		case "v":
			currentCoard.y--
		}
		locations[*currentCoard] = true
	}
	return len(locations)
}
