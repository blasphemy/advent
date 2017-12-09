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
		log.Fatal(err)
	}
	answer1, answer2 := score(string(inputBytes))
	fmt.Printf("Part 1: %d\nPart 2: %d\n", answer1, answer2)
}

func score(in string) (int, int) {
	score := 0
	index := 0
	depth := 0
	garbage := false
	garbageScore := 0
	s := strings.Split(in, "")
	for {
		switch s[index] {
		case "!":
			index++
			break
		case "<":
			if !garbage {
				garbage = true
				break
			} else {
				garbageScore++
			}
		case ">":
			garbage = false
		case "{":
			if !garbage {
				depth++
			} else {
				garbageScore++
			}
		case "}":
			if !garbage {
				score += depth
				depth--
			} else {
				garbageScore++
			}
		default:
			if garbage {
				garbageScore++
			}
		}
		index++
		if index >= len(s) {
			break
		}
	}
	return score, garbageScore
}
