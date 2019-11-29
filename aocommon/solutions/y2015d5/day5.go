package y2015d5

import (
	"aoc2/aocommon/solutions"
	"strings"
)

var Solution = solutions.AOCSolution{
	Year: 2015,
	Day: 5,
	DefaultInput: Input,
	Answer1Func: a1,
	Answer2Func: a2,
}

const (
	v1 = iota
	v2 = iota
)

type version int

var (
	badStrings = []string{"ab", "cd", "pq", "xy"}
	vowels     = []string{"a", "e", "i", "o", "u"}
)

func a1(i string) int {
	rows := strings.Split(i, "\n")
	return countNiceStrings(rows, v1)
}

func a2(i string) int {
	rows := strings.Split(i, "\n")
	return countNiceStrings(rows, v2)
}

func countNiceStrings(in []string, v version) int {
	niceCounter := 0
	for _, x := range in {
		if isNice(x, v) {
			niceCounter++
		}
	}
	return niceCounter
}

func isNice(in string, v version) bool {
	if v == v1 {
		if countVowels(in) < 3 {
			return false
		}
		if !hasDoubleLetter(in) {
			return false
		}
		if hasBadString(in) {
			return false
		}
	} else if v == v2 {
		if !hasPair(in) {
			return false
		}
		if !hasSplitPair(in) {
			return false
		}
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

func hasPair(in string) bool {
	split := strings.Split(in, "")
	for x, y := range split {
		if x+1 >= len(split) {
			continue
		}
		potentialPair := y + split[x+1]
		testString := sliceToString(split[x+2:])
		if strings.Contains(testString, potentialPair) {
			return true
		}
	}
	return false
}

func hasSplitPair(in string) bool {
	split := strings.Split(in, "")
	for i := 1; i < len(split)-1; i++ {
		if split[i+1] == split[i-1] {
			return true
		}
	}
	return false
}

func sliceToString(in []string) string {
	out := ""
	for _, x := range in {
		out = out + x
	}
	return out
}
