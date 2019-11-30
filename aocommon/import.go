package aocommon

import (
	"aoc2/aocommon/solutions/y2015d1"
	"aoc2/aocommon/solutions/y2015d2"
	"aoc2/aocommon/solutions/y2015d3"
	"aoc2/aocommon/solutions/y2015d4"
	"aoc2/aocommon/solutions/y2015d5"
	"aoc2/aocommon/solutions/y2015d6"
	"aoc2/aocommon/solutions/y2016d1"
	"aoc2/aocommon/solutions/y2017d1"
)

func registerAll() {
	registerSolution(y2015d1.Solution)
	registerSolution(y2015d2.Solution)
	registerSolution(y2015d3.Solution)
	registerSolution(y2015d4.Solution)
	registerSolution(y2015d5.Solution)
	registerSolution(y2015d6.Solution)
	registerSolution(y2016d1.Solution)
	registerSolution(y2017d1.Solution)
}
