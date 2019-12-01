package y18d2

import "advent/aocommon/solutions"

import "strings"

import "strconv"

var Solution = solutions.AOCSolution{
	Year:         18,
	Day:          2,
	DefaultInput: input,
	Answer1Func:  a1,
	Answer2Func:  a2,
}

const sample = `abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab`

const sample2 = `abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz`

func a1(i string) string {
	rows := strings.Split(i, "\n")
	doubles := 0
	triples := 0
	for _, x := range rows {
		if containsDouble(x) {
			doubles++
		}
		if containsTriple(x) {
			triples++
		}
	}
	return strconv.Itoa(triples * doubles)
}

func a2(i string) string {
	for _, x := range strings.Split(i, "\n") {
		for _, y := range strings.Split(i, "\n") {
			if x == y {
				continue
			}
			match := matches(x, y)
			if len(match) > len(x)-2 {
				return string(match)
			}
		}
	}
	return ""
}

func containsDouble(in string) bool {
	letters := make(map[rune]int)
	for _, x := range in {
		letters[x]++
	}
	for _, x := range letters {
		if x == 2 {
			return true
		}
	}
	return false
}

func containsTriple(in string) bool {
	letters := make(map[rune]int)
	for _, x := range in {
		letters[x]++
	}
	for _, x := range letters {
		if x == 3 {
			return true
		}
	}
	return false
}

func matches(s1, s2 string) []rune {
	matches := []rune{}
	for x, y := range s1 {
		if rune(s2[x]) == y {
			matches = append(matches, y)
		}
	}
	return matches
}
