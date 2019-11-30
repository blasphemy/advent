package y17d3

import (
	"strconv"
	"advent/aocommon/solutions"
	"math"
)

const (
	challengeInput = "361527"
)

var Solution = solutions.AOCSolution{
	Year: 17,
	Day: 3,
	DefaultInput: challengeInput,
	Answer1Func: a1,
	Answer2Func: a2,
}

func a1(i string) string {
	inconvert, _ := strconv.Atoi(i) //bad me
	c := getSpiralCoards(inconvert)
	return strconv.Itoa(getStepsBack(c))
}

func a2(i string) string {
	inconvert, _ := strconv.Atoi(i)
	return strconv.Itoa(getBiggerThanInput(inconvert))
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
		adjPos := getAdjacentPositions(currentCoards)
		for _, x := range adjPos {
			sum = sum + valueMap[x]
		}
		valueMap[currentCoards] = sum
		if sum > i {
			return sum
		}
		counter++
	}
	return 0
}

func getAdjacentPositions(i coards) []coards {
	out := []coards{}
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			c := i
			c.x = c.x + x
			c.y = c.y + y
			if c != i {
				out = append(out, c)
			}
		}
	}
	return out
}
