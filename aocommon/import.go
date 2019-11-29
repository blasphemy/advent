package aocommon

import "aoc2/aocommon/solutions/y2015d1"
import "aoc2/aocommon/solutions/y2015d2"
import "aoc2/aocommon/solutions/y2015d3"
import "aoc2/aocommon/solutions/y2015d4"
import "aoc2/aocommon/solutions/y2015d5"
import "aoc2/aocommon/solutions/y2015d6"

func registerAll() {
	registerSolution(y2015d1.Solution)
	registerSolution(y2015d2.Solution)
	registerSolution(y2015d3.Solution)
	registerSolution(y2015d4.Solution)
	registerSolution(y2015d5.Solution)
	registerSolution(y2015d6.Solution)
}
