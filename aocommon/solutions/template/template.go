package template

import "advent/aocommon/solutions"

var Solution = solutions.AOCSolution{
	Year:         99,
	Day:          99,
	DefaultInput: "test",
	Answer1Func:  a1,
	Answer2Func:  a1,
}

func a1(i string) string {
	return i
}
