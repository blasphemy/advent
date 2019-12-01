package y19d1

import "advent/aocommon/solutions"

import "strings"

import "strconv"

var Solution = solutions.AOCSolution{
	Year:         19,
	Day:          1,
	DefaultInput: input,
	Answer1Func:  a1,
	Answer2Func:  a2,
}

func a1(i string) string {
	nums,err := parseInput(i)
	if err != nil {
		panic(err)
	}
	masses := []int{}
	for _,x := range nums {
		mass := getMass(x)
		masses= append(masses,mass)
	}
	sum := 0
	for _,x := range masses {
		sum += x
	}
	return strconv.Itoa(sum)
}

func a2(i string) string {
	nums,err := parseInput(i)
	if err != nil {
		panic(err)
	}
	masses := []int{}
	for _,x := range nums {
		mass := getMass2(x)
		masses= append(masses,mass)
	}
	sum := 0
	for _, x := range masses {
		sum += x
	}
	return strconv.Itoa(sum)
}

func parseInput(i string) ([]int,error) {
	rows := strings.Split(i, "\n")
	res := []int{}
	for _,x := range rows {
		n, err := strconv.Atoi(x)
		if err != nil {
			return []int{}, err
		}
		res = append(res, n)
	}
	return res, nil
}

func getMass(i int) int {
	w := i / 3
	w = w - 2
	return w
}

func getMass2(i int) int {
	w := getMass(i)
	x := w
	for w > 0 {
		w = getMass(w)
		if w < 1 {
			w = 0
		}
		x = x + w
	}
	return x
}