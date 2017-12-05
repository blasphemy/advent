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
		log.Fatal(err)
	}
	rows := inputToIntSlice(string(inputBytes))
	rows2 := inputToIntSlice(string(inputBytes))
	answer1 := countStepsOut(rows)
	fmt.Printf("Part 1: %d\n", answer1)
	answer2 := countStepsOutV2(rows2)
	fmt.Printf("Part 2: %d\n", answer2)
}

func inputToIntSlice(in string) []int {
	split := strings.Split(in, "\n")
	out := []int{}
	for _, x := range split {
		n, err := strconv.Atoi(x)
		if err != nil {
			log.Fatal(err)
		}
		out = append(out, n)
	}
	return out
}

func countStepsOut(rows []int) int {
	counter := 0
	currentPos := 0
	for currentPos < len(rows) {
		jmpAmount := rows[currentPos]
		rows[currentPos]++
		currentPos = currentPos + jmpAmount
		counter++
	}
	return counter
}

func countStepsOutV2(rows []int) int {
	counter := 0
	currentPos := 0
	for currentPos < len(rows) {
		jmpAmount := rows[currentPos]
		if jmpAmount >= 3 {
			rows[currentPos]--
		} else {
			rows[currentPos]++
		}
		currentPos = currentPos + jmpAmount
		counter++
	}
	return counter
}
