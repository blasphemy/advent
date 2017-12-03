package main

import (
	"fmt"
	"math"
)

const (
	challengeInput = 361527
)

func main() {
	coardinates := getSpiralCoards(challengeInput)
	answer1 := getStepsBack(coardinates)
	answer2 := getBiggerThanInput(challengeInput)
	fmt.Printf("Part 1: %d\n", answer1)
	fmt.Printf("Part 2: %d\n", answer2)
}

type coards struct {
	x int
	y int
}

func getSpiralCoards(i int) coards {
	currentCoards := coards{0, 0}
	segCounter := 0
	segCounter2 := 0
	segLength := 1
	/*
		direction 1 = east
		direction 2 = north
		direction 3 = west
		direction 4 = south
	*/
	dir := 1
	counter := 1
	for true {
		if counter == i {
			return currentCoards
		}
		switch dir {
		case 1:
			currentCoards.x++
		case 2:
			currentCoards.y++
		case 3:
			currentCoards.x--
		case 4:
			currentCoards.y--
		}
		segCounter++
		if segCounter == segLength {
			segCounter = 0
			dir++
			if dir == 5 {
				dir = 1
			}
			segCounter2++
			if segCounter2 > 1 {
				segLength++
				segCounter2 = 0
			}
		}
		counter++
	}
	return coards{0, 0}
}

func getStepsBack(c coards) int {
	return int(math.Abs(float64(c.x)) + math.Abs(float64(c.y)))
}

func getBiggerThanInput(i int) int {
	valueMap := make(map[coards]int)
	valueMap[coards{0, 0}] = 1
	counter := 2
	for true {
		var sum int
		currentCoards := getSpiralCoards(counter)
		startingCoards := currentCoards
		currentCoards.x++
		sum = sum + valueMap[currentCoards]
		currentCoards.y++
		sum = sum + valueMap[currentCoards]
		currentCoards.x--
		sum = sum + valueMap[currentCoards]
		currentCoards.x--
		sum = sum + valueMap[currentCoards]
		currentCoards.y--
		sum = sum + valueMap[currentCoards]
		currentCoards.y--
		sum = sum + valueMap[currentCoards]
		currentCoards.x++
		sum = sum + valueMap[currentCoards]
		currentCoards.x++
		sum = sum + valueMap[currentCoards]
		valueMap[startingCoards] = sum
		if sum > i {
			return sum
		}
		counter++
	}
	return 0
}
