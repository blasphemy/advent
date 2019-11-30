package y17d4

import (
	"advent/aocommon/solutions"
	"sort"
	"strconv"
	"strings"
)

var Solution = solutions.AOCSolution{
	Year:         17,
	Day:          4,
	DefaultInput: input,
	Answer1Func:  a1,
	Answer2Func:  a2,
}

func a1(i string) string {
	rows := strings.Split(i, "\n")
	ans := countSecurePhrases(rows, v1)
	return strconv.Itoa(ans)
}

func a2(i string) string {
	rows := strings.Split(i, "\n")
	ans := countSecurePhrases(rows, v2)
	return strconv.Itoa(ans)
}

type version int

const (
	v1 = iota
	v2 = iota
)

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
