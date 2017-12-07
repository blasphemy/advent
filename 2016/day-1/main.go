package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type place struct {
	x int
	y int
}

func main() {
	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	answer1 := findDistanceV1(string(inputBytes))
	answer2 := findDistanceV2(string(inputBytes))
	fmt.Printf("Part 1: %d\nPart 2: %d\n", answer1, answer2)
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
