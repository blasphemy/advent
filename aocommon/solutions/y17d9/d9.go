package y17d9

import (
	"advent/aocommon/solutions"
	"fmt"
	"io/ioutil"
	"log"
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

func main() {
	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	answer1, answer2 := score(string(inputBytes))
	fmt.Printf("Part 1: %d\nPart 2: %d\n", answer1, answer2)
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
