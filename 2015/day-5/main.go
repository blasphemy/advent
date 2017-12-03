package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var (
	badStrings = []string{"ab", "cd", "pq", "xy"}
	vowels     = []string{"a", "e", "i", "o", "u"}
)

func main() {
	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	challengeInput := string(inputBytes)
	wordRows := strings.Split(challengeInput, "\n")
	answer1 := countNiceStrings(wordRows)
	fmt.Printf("Part 1: %d\n", answer1)
}

func countNiceStrings(in []string) int {
	niceCounter := 0
	for _, x := range in {
		if isNice(x) {
			niceCounter++
		}
	}
	return niceCounter
}

func isNice(in string) bool {
	if countVowels(in) < 3 {
		return false
	}
	if !hasDoubleLetter(in) {
		return false
	}
	if hasBadString(in) {
		return false
	}
	return true
}

func countVowels(in string) int {
	split := strings.Split(in, "")
	vowelCounter := 0
	for _, x := range split {
		for _, y := range vowels {
			if x == y {
				vowelCounter++
			}
		}
	}
	return vowelCounter
}

func hasDoubleLetter(in string) bool {
	split := strings.Split(in, "")
	for x, y := range split {
		if x+1 == len(split) {
			continue
		}
		if y == split[x+1] {
			return true
		}
	}
	return false
}

func hasBadString(in string) bool {
	for _, x := range badStrings {
		if strings.Contains(in, x) {
			return true
		}
	}
	return false
}
