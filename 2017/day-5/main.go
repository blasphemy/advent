package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rows := inputToIntSlice(string(inputBytes))
	rows2 := inputToIntSlice(string(inputBytes))
	t := time.Now()
	answer1 := countStepsOut(rows)
	d1 := time.Since(t)
	t = time.Now()
	answer2 := countStepsOutV2(rows2)
	d2 := time.Since(t)
	fmt.Printf("Part 1: %d (in %s)\n", answer1, d1)
	fmt.Printf("Part 2: %d (in %s)\n", answer2, d2)
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
