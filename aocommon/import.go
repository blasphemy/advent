package aocommon

import "aoc2/aocommon/solutions/y2015d1"
import "aoc2/aocommon/solutions/y2015d2"
import "aoc2/aocommon/solutions/y2015d3"

func registerAll() {
	registerSolution(y2015d1.Solution)
	registerSolution(y2015d2.Solution)
	registerSolution(y2015d3.Solution)
}
