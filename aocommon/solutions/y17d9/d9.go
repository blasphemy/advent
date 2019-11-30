package y17d9

import (
	"advent/aocommon/solutions"
	"strconv"
	"strings"
)

var Solution = solutions.AOCSolution{
	Year:         17,
	Day:          9,
	DefaultInput: input,
	Answer1Func:  a1,
	Answer2Func:  a2,
}

func a1(in string) string {
	ans, _ := score(in)
	return strconv.Itoa(ans)
}

func a2(in string) string {
	_, ans := score(in)
	return strconv.Itoa(ans)
}

func score(in string) (int, int) {
	score := 0
	index := 0
	depth := 0
	garbage := false
	garbageScore := 0
	s := strings.Split(in, "")
	for {
		switch s[index] {
		case "!":
			index++
			break
		case "<":
			if !garbage {
				garbage = true
				break
			} else {
				garbageScore++
			}
		case ">":
			garbage = false
		case "{":
			if !garbage {
				depth++
			} else {
				garbageScore++
			}
		case "}":
			if !garbage {
				score += depth
				depth--
			} else {
				garbageScore++
			}
		default:
			if garbage {
				garbageScore++
			}
		}
		index++
		if index >= len(s) {
			break
		}
	}
	return score, garbageScore
}
