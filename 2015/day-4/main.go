package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	challengeInput := "bgvyzdsv"
	answer1 := getKeyForInput(challengeInput, "00000")
	fmt.Printf("Part 1: %d\n", answer1)
	answer2 := getKeyForInput(challengeInput, "000000")
	fmt.Printf("Part 2: %d\n", answer2)
}

//This is probably the slowest possible way to do this in go, but oh well.
//Santa will just end up getting an ASIC for mining anyways
func getKeyForInput(i string, searchString string) int {
	counter := 0
	for true {
		testString := i + strconv.Itoa(counter)
		out := md5.Sum([]byte(testString))
		testOut := fmt.Sprintf("%x", out)
		if strings.HasPrefix(testOut, searchString) {
			return counter
		}
		counter++
	}
	return 0
}
