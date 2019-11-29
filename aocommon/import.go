package aocommon

import "aoc2/aocommon/solutions/y2015d1"
import "aoc2/aocommon/solutions/y2015d2"
import "aoc2/aocommon/solutions/y2015d3"

func registerAll() {
	registerSolution(2015,1,y2015d1.Solution)
	registerSolution(2015, 2, y2015d2.Solution)
	registerSolution(2015,3,y2015d3.Solution)
}
