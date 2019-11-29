package aocommon

import "errors"
import "aoc2/aocommon/solutions"

var solutionRegistry map[AOCKey]solutions.AOCSolution

func init() {
	solutionRegistry = make(map[AOCKey]solutions.AOCSolution)
	registerAll()
}

func registerSolution(year, day int, solution solutions.AOCSolution) {
	//this one takes ints to make it easier on me
	k := getKey(year, day)
	solutionRegistry[k] = solution
}

func getSolution(key AOCKey) (solutions.AOCSolution, error) {
	f, ok := solutionRegistry[key]
	if !ok {
		e := errors.New("Function Does Not Exist")
		return solutions.AOCSolution{}, e
	}
	return f, nil
}

func AOCAvailable() []AOCKey {
	keys := []AOCKey{}
	for x := range solutionRegistry {
		keys = append(keys, x)
	}
	return keys
}
