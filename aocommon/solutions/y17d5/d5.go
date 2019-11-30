package y17d5

import (
	"advent/aocommon/solutions"
	"log"
	"strconv"
	"strings"
)

var Solution = solutions.AOCSolution{
	Year: 17,
	Day: 5,
	DefaultInput: input,
	Answer1Func: a1,
	Answer2Func: a2,
}

func a1(i string) string {
	rows := inputToIntSlice(i)
	ans := countStepsOut(rows)
	return strconv.Itoa(ans)
}

func a2(i string) string {
	rows := inputToIntSlice(i)
	ans := countStepsOutV2(rows)
	return strconv.Itoa(ans)
}

func inputToIntSlice(in string) []int {
	split := strings.Split(in, "\n")
	out := []int{}
	for _, x := range split {
		n, err := strconv.Atoi(x)
		if err != nil {
			log.Fatal(err)
		}
		out = append(out, n)
	}
	return out
}

func countStepsOut(rows []int) int {
	counter := 0
	currentPos := 0
	for currentPos < len(rows) {
		jmpAmount := rows[currentPos]
		rows[currentPos]++
		currentPos = currentPos + jmpAmount
		counter++
	}
	return counter
}

func countStepsOutV2(rows []int) int {
	counter := 0
	currentPos := 0
	for currentPos < len(rows) {
		jmpAmount := rows[currentPos]
		if jmpAmount >= 3 {
			rows[currentPos]--
		} else {
			rows[currentPos]++
		}
		currentPos = currentPos + jmpAmount
		counter++
	}
	return counter
}
