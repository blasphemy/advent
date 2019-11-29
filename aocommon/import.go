package aocommon

import "aoc2/aocommon/solutions/y2015d1"
import "aoc2/2015/day-2"

func importFuncs() {
	registerWithInput(2015, 1, 1, y2015d1.Challenge, y2015d1.Answer1)
	registerWithInput(2015, 1, 2, y2015d1.Challenge, y2015d1.Answer2)
	registerWithInput(2015, 2, 1, day2.Challenge, day2.Answer1)
	registerWithInput(2015, 2, 2, day2.Challenge, day2.Answer2)
}
