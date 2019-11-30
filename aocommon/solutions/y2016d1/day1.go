package y2016d1

import (
	"advent/aocommon/solutions"
	"log"
	"math"
	"strconv"
	"strings"
)

var Solution = solutions.AOCSolution{
	Year: 16,
	Day: 1,
	DefaultInput: INPUT,
	Answer1Func: findDistanceV1,
	Answer2Func: findDistanceV2,
}

const INPUT = "R4, R5, L5, L5, L3, R2, R1, R1, L5, R5, R2, L1, L3, L4, R3, L1, L1, R2, R3, R3, R1, L3, L5, R3, R1, L1, R1, R2, L1, L4, L5, R4, R2, L192, R5, L2, R53, R1, L5, R73, R5, L5, R186, L3, L2, R1, R3, L3, L3, R1, L4, L2, R3, L5, R4, R3, R1, L1, R5, R2, R1, R1, R1, R3, R2, L1, R5, R1, L5, R2, L2, L4, R3, L1, R4, L5, R4, R3, L5, L3, R4, R2, L5, L5, R2, R3, R5, R4, R2, R1, L1, L5, L2, L3, L4, L5, L4, L5, L1, R3, R4, R5, R3, L5, L4, L3, L1, L4, R2, R5, R5, R4, L2, L4, R3, R1, L2, R5, L5, R1, R1, L1, L5, L5, L2, L1, R5, R2, L4, L1, R4, R3, L3, R1, R5, L1, L4, R2, L3, R5, R3, R1, L3"

type place struct {
	x int
	y int
}

func findDistanceV1(in string) int {
	dir := 0
	x := 0
	y := 0
	words := strings.Split(in, " ")
	for _, z := range words {
		z = strings.TrimSuffix(z, ",")
		if strings.HasPrefix(z, "L") {
			z = strings.TrimPrefix(z, "L")
			dir--
		} else {
			z = strings.TrimPrefix(z, "R")
			dir++
		}
		if dir > 3 {
			dir = 0
		}
		if dir < 0 {
			dir = 3
		}
		amount, err := strconv.Atoi(z)
		if err != nil {
			log.Fatal(err)
		}
		switch dir {
		case 0:
			y += amount
		case 1:
			x += amount
		case 2:
			y -= amount
		case 3:
			x -= amount
		}
	}
	return int(math.Abs(float64(x) + math.Abs(float64(y))))
}

func findDistanceV2(in string) int {
	dir := 0
	x := 0
	y := 0
	places := make(map[place]bool)
	places[place{x, y}] = true
	words := strings.Split(in, " ")
	for _, z := range words {
		z = strings.TrimSuffix(z, ",")
		if strings.HasPrefix(z, "L") {
			z = strings.TrimPrefix(z, "L")
			dir--
		} else {
			z = strings.TrimPrefix(z, "R")
			dir++
		}
		if dir > 3 {
			dir = 0
		}
		if dir < 0 {
			dir = 3
		}
		amount, err := strconv.Atoi(z)
		if err != nil {
			log.Fatal(err)
		}
		var use *int
		switch dir {
		case 0:
			use = &y
		case 1:
			use = &x
		case 2:
			use = &y
			amount = -amount
		case 3:
			use = &x
			amount = -amount
		}
		for amount != 0 {
			if amount > 0 {
				amount--
				*use++
			}
			if amount < 0 {
				amount++
				*use--
			}
			p := place{x, y}
			if places[p] {
				return int(math.Abs(float64(x) + math.Abs(float64(y))))
			}
			places[p] = true
		}
	}
	return 0
}
