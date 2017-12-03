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
	challengeInput := string(inputBytes)
	answer1 := getAnswer1(challengeInput)
	answer2 := getAnswer2(challengeInput)
	fmt.Printf("Part 1: %d\n", answer1)
	fmt.Printf("Part 2: %d\n", answer2)
}

func getAnswer1(input string) int {
	rows := strings.Split(input, "\n")
	sum := 0
	for _, x := range rows {
		numString := strings.Split(x, "\t")
		nums, err := sliceAtoi(numString)
		if err != nil {
			log.Fatal(err.Error())
		}
		big, small := maxMin(nums)
		diff := big - small
		sum = sum + diff
	}
	return sum
}

func maxMin(i []int) (int, int) {
	big := 0
	for _, y := range i {
		if y > big {
			big = y
		}
	}
	small := big
	for _, y := range i {
		if y < small {
			small = y
		}
	}
	return big, small
}

func sliceAtoi(i []string) ([]int, error) {
	out := []int{}
	for _, y := range i {
		n, err := strconv.Atoi(y)
		if err != nil {
			return []int{}, err
		}
		out = append(out, n)
	}
	return out, nil
}

func getAnswer2(input string) int {
	rows := strings.Split(input, "\n")
	sum := 0
	for _, x := range rows {
		numString := strings.Split(x, "\t")
		nums, err := sliceAtoi(numString)
		if err != nil {
			log.Fatal(err.Error())
		}
		sum = sum + getDivisibleNums(nums)
	}
	return sum
}

func getDivisibleNums(i []int) int {
	for _, x := range i {
		for _, y := range i {
			if x == y {
				continue
			}
			if x%y == 0 {
				return x / y
			}
		}
	}
	return 0
}
