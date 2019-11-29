package aocommon

import "aoc2/aocommon/solutions/y2015d1"
import "aoc2/aocommon/solutions/y2015d2"

func importFuncs() {
	RegisterSolution(2015, 1, AOCSolution{
		Answer1Func: y2015d1.Answer1,
		Answer2Func: y2015d1.Answer2,
		DefaultInput: y2015d1.Challenge,
	})
	RegisterSolution(2015, 2, AOCSolution{
		Answer1Func: y2015d2.Answer1,
		Answer2Func: y2015d2.Answer2,
		DefaultInput: y2015d2.Challenge,
	})
}
