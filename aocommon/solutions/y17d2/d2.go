package y17d2

import (
	"advent/aocommon/solutions"
	"log"
	"strconv"
	"strings"
)

var Solution = solutions.AOCSolution{
	Year:         17,
	Day:          2,
	DefaultInput: input,
	Answer1Func:  getAnswer1,
	Answer2Func:  getAnswer2,
}

func getAnswer1(input string) string {
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
	return strconv.Itoa(sum)
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

func getAnswer2(input string) string {
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
	return strconv.Itoa(sum)
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
