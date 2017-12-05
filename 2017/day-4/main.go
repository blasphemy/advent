package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"time"
)

type version int

const (
	v1 = iota
	v2 = iota
)

func main() {
	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	challengeInput := string(inputBytes)
	rows := strings.Split(challengeInput, "\n")
	t := time.Now()
	answer1 := countSecurePhrases(rows, v1)
	d1 := time.Since(t)
	t = time.Now()
	answer2 := countSecurePhrases(rows, v2)
	d2 := time.Since(t)
	fmt.Printf("Part 1: %d (in %s)\nPart 2: %d (in %s)\n", answer1, d1, answer2, d2)
}

func countSecurePhrases(in []string, v version) int {
	counter := 0
	for _, x := range in {
		if isSecure(x, v) {
			counter++
		}
	}
	return counter
}

func isSecure(in string, v version) bool {
	split := strings.Split(in, " ")
	words := make(map[string]bool)
	for _, x := range split {
		if v == v2 {
			x = sortString(x)
		}
		if words[x] {
			return false
		}
		words[x] = true
	}
	return true
}

func sortString(in string) string {
	split := strings.Split(in, "")
	sort.Strings(split)
	out := ""
	for _, x := range split {
		out = out + x
	}
	return out
}
