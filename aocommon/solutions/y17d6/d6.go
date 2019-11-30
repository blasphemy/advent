package y17d6

import (
	"advent/aocommon/solutions"
	"log"
	"strconv"
	"strings"
)

const input = "10	3	15	10	5	15	5	15	9	2	5	8	5	2	3	6"

var Solution = solutions.AOCSolution{
	Year:         17,
	Day:          6,
	DefaultInput: input,
	Answer1Func:  a1,
	Answer2Func:  a2,
}

func a1(i string) string {
	istate := inputToArr(i)
	ans1, _ := balance(istate)
	return strconv.Itoa(ans1)
}

func a2(i string) string {
	istate := inputToArr(i)
	_, ans2 := balance(istate)
	return strconv.Itoa(ans2)
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
