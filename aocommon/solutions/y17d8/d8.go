package y17d8

import (
	"advent/aocommon/solutions"
	"strconv"
	"strings"
	"fmt"
)

var Solution = solutions.AOCSolution{
	Year: 17,
	Day: 8,
	DefaultInput: input,
	Answer1Func: a1,
	Answer2Func: a2,
}

type instruction struct {
	target           string
	targetOp         string
	targetAmount     int
	conditionSubject string
	conditionOp      string
	conditionAmount  int
}

func a1(i string) string {
	ans, _ := runInstructions(i)
	return strconv.Itoa(ans)
}

func a2(i string) string {
	_, ans := runInstructions(i)
	return strconv.Itoa(ans)
}

func runInstructions(in string) (int, int) {
	registers := make(map[string]int)
	lines := strings.Split(in, "\n")
	max := 0
	for _, x := range lines {
		inst := lineToInstruction(x)
		if evaluateInstruction(inst, registers) {
			testMax := executeInstruction(inst, registers)
			if testMax > max {
				max = testMax
			}
		}
	}
	return getMax(registers), max
}

func lineToInstruction(in string) instruction {
	out := instruction{}
	fmt.Sscanf(in, "%s %s %d if %s %s %d", &out.target, &out.targetOp, &out.targetAmount, &out.conditionSubject, &out.conditionOp, &out.conditionAmount)
	return out
}

func evaluateInstruction(i instruction, data map[string]int) bool {
	sub := data[i.conditionSubject]
	amt := i.conditionAmount
	switch i.conditionOp {
	case ">":
		return sub > amt
	case "<":
		return sub < amt
	case ">=":
		return sub >= amt
	case "==":
		return sub == amt
	case "<=":
		return sub <= amt
	case "!=":
		return sub != amt
	}
	return false
}

func executeInstruction(i instruction, data map[string]int) int {
	switch i.targetOp {
	case "inc":
		data[i.target] += i.targetAmount
	case "dec":
		data[i.target] -= i.targetAmount
	}
	return data[i.target]
}

func getMax(data map[string]int) int {
	max := 0
	for _, x := range data {
		if x > max {
			max = x
		}
	}
	return max
}
