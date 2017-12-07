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
	initialState := inputToArr(string(inputBytes))
	answer1, answer2 := balance(initialState)
	fmt.Printf("Part 1: %d\nPart 2: %d\n", answer1, answer2)
}

func balance(in []int) (int, int) {
	patterns := make(map[string]int)
	counter := 0
	for {
		index := chooseBiggest(in)
		amountToBalance := in[index]
		in[index] = 0
		for amountToBalance > 0 {
			index++
			if index >= len(in) {
				index = 0
			}
			in[index]++
			amountToBalance--
		}
		counter++
		arrString := arrToString(in)
		if patterns[arrString] > 0 {
			return counter, counter - patterns[arrString]
		}
		patterns[arrString] = counter
	}
}

func arrToString(in []int) string {
	out := ""
	for _, x := range in {
		out += strconv.Itoa(x)
	}
	return out
}

func chooseBiggest(in []int) int { //index number
	max := 0
	index := 0
	for x, y := range in {
		if y > max {
			index = x
			max = y
		}
	}
	return index
}

func inputToArr(in string) []int {
	split := strings.Split(in, "\t")
	arr := []int{}
	for _, x := range split {
		num, err := strconv.Atoi(x)
		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, num)
	}
	return arr
}
