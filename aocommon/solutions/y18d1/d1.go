package y18d1

import "advent/aocommon/solutions"
import "strings"

import "strconv"

var Solution = solutions.AOCSolution{
	Year:         18,
	Day:          1,
	DefaultInput: input,
	Answer1Func:  a1,
	Answer2Func:  a2,
}

func a1(i string) string {
	nums, err := parseInput(i)
	if err != nil {
		panic(err)
	}
	n := sumNums(nums)
	return strconv.Itoa(n)
}

func a2(i string) string {
	nums, err := parseInput(i)
	if err != nil {
		panic(err)
	}
	n := seenTwice(nums)
	return strconv.Itoa(n)
}

func parseInput(i string) ([]int, error) {
	split := strings.Split(i, "\n")
	out := []int{}
	for _, x := range split {
		neg := false
		if strings.HasPrefix(x, "-") {
			neg = true
		}
		s := strings.TrimLeft(x, "+-")
		n, err := strconv.Atoi(s)
		if err != nil {
			return []int{}, err
		}
		if neg {
			n = -n
		}
		out = append(out, n)
	}
	return out, nil
}

func sumNums(i []int) int {
	s := 0
	for _, x := range i {
		s += x
	}
	return s
}

func seenTwice(i []int) int {
	seen := make(map[int]bool)
	s := 0
	for true {
		for _, x := range i {
			seen[s] = true
			s += x
			if seen[s] {
				return s
			}
		}
	}
	return 0
}
