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
	input := string(inputBytes)
	answer1 := partOne(inputToIntArray(input))
	answer2 := stringToKnotHash(input)
	fmt.Printf("Part 1: %d\nPart 2: %s\n", answer1, answer2)
}

func stringToKnotHash(in string) string {
	ls := stringToLengthSequence(in)
	sh := sparseHash(ls, 64)
	dh := denseHash(sh)
	return denseHashToString(dh)
}

func partOne(in []int) int {
	l := sparseHash(in, 1)
	return l[0] * l[1]
}

func sparseHash(in []int, rounds int) []int {
	l := initList()
	idx := 0
	skip := 0
	for i := 0; i < rounds; i++ {
		for _, x := range in {
			for idx > len(l) {
				idx -= len(l)
			}
			l = reverse(l, idx, x)
			idx += x + skip
			skip++
		}
	}
	return l
}

func inputToIntArray(in string) []int {
	split := strings.Split(in, ",")
	out := []int{}
	for _, x := range split {
		num, err := strconv.Atoi(x)
		if err != nil {
			log.Fatal(err)
		}
		out = append(out, num)
	}
	return out
}

func initList() []int {
	out := []int{}
	for i := 0; i <= 255; i++ {
		out = append(out, i)
	}
	return out
}

func reverse(in []int, idx int, size int) []int {
	idxlist := []int{}
	currentIdx := idx
	counter := 0
	for {
		if currentIdx > len(in)-1 {
			currentIdx -= len(in)
		}
		idxlist = append(idxlist, currentIdx)
		currentIdx++
		counter++
		if counter >= size {
			break
		}
	}
	reversed := []int{}
	for i := len(idxlist) - 1; i >= 0; i-- {
		reversed = append(reversed, in[idxlist[i]])
	}
	final := in
	for y, x := range idxlist {
		final[x] = reversed[y]
	}
	return final
}

func stringToASCIICodes(in string) []int {
	r := []rune(in)
	out := []int{}
	for _, x := range r {
		out = append(out, int(x))
	}
	return out
}

func stringToLengthSequence(in string) []int {
	add := []int{17, 31, 73, 47, 23}
	ascii := stringToASCIICodes(in)
	return append(ascii, add...)
}

func denseHash(in []int) []int {
	blockSize := 16
	bcounter := 0
	out := []int{}
	workingSet := []int{}
	for _, x := range in {
		workingSet = append(workingSet, x)
		bcounter++
		if bcounter >= blockSize {
			bcounter = 0
			out = append(out, xornums(workingSet))
			workingSet = []int{}
		}
	}
	return out
}

func denseHashToString(in []int) string {
	out := ""
	for _, x := range in {
		out += fmt.Sprintf("%.02x", x)
	}
	return out
}

func xornums(in []int) int {
	working := 0
	for _, x := range in {
		working ^= x
	}
	return working
}
