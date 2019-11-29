package aocommon

import "aoc2/aocommon/solutions/y2015d1"
import "aoc2/aocommon/solutions/y2015d2"

func importFuncs() {
	registerWithInput(2015, 1, y2015d1.Challenge, y2015d1.Answer1, y2015d1.Answer2)
	registerWithInput(2015, 2, y2015d2.Challenge, y2015d2.Answer1, y2015d2.Answer2)
}
