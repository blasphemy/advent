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
	answer1 := getAnswer1(challengeInput)
	answer2 := getAnswer2(challengeInput)
	fmt.Printf("Answer 1: %d\n", answer1)
	fmt.Printf("Answer 2: %d\n", answer2)
}

func getAnswer1(i string) int {
	floor := 0
	instructions := strings.Split(i, "")
	for _, x := range instructions {
		switch x {
		case "(":
			floor++
		case ")":
			floor--
		}
	}
	return floor
}

func getAnswer2(i string) int {
	floor := 0
	instructions := strings.Split(i, "")
	for y, x := range instructions {
		switch x {
		case "(":
			floor++
		case ")":
			floor--
		}
		if floor == -1 {
			return y + 1
		}
	}
	return 0
}
