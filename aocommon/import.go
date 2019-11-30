package aocommon

import (
	"advent/aocommon/solutions/y2015d1"
	"advent/aocommon/solutions/y17d6"
	"advent/aocommon/solutions/y17d5"
	"advent/aocommon/solutions/y2015d2"
	"advent/aocommon/solutions/y2015d3"
	"advent/aocommon/solutions/y2015d4"
	"advent/aocommon/solutions/y2015d5"
	"advent/aocommon/solutions/y2015d6"
	"advent/aocommon/solutions/y2016d1"
	"advent/aocommon/solutions/y2017d1"
	"advent/aocommon/solutions/y17d2"
	"advent/aocommon/solutions/y17d3"
	"advent/aocommon/solutions/y17d4"
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
	registerSolution(y17d2.Solution)
	registerSolution(y17d3.Solution)
	registerSolution(y17d4.Solution)
	registerSolution(y17d5.Solution)
	registerSolution(y17d6.Solution)
}
