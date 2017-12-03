package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	inputString := string(inputBytes)
	answer1 := getNumberPart1(inputString)
	answer2 := getNumberPart2(inputString)
	fmt.Printf("Part 1 Answer: %d\n", answer1)
	fmt.Printf("Part 2 Answer: %d\n", answer2)
}

func getNumberPart1(input string) int {
	pieces := strings.Split(input, "")
	sum := 0
	for x, y := range pieces {
		if x == len(pieces)-1 {
			x = 0
		} else {
			x = x + 1
		}
		if y == pieces[x] {
			n, err := strconv.Atoi(y)
			if err != nil {
				log.Fatal(err.Error())
			}
			sum = sum + n
		}
	}
	return sum
}

func getNumberPart2(input string) int {
	pieces := strings.Split(input, "")
	sum := 0
	for x, y := range pieces {
		x = x + len(pieces)/2
		if x >= len(pieces) {
			x = x - len(pieces)
		}
		if y == pieces[x] {
			n, err := strconv.Atoi(y)
			if err != nil {
				log.Fatal(err.Error())
			}
			sum = sum + n
		}
	}
	return sum
}
