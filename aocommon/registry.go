package aocommon

import "errors"

type AOCSolution struct {
	Answer1Func  AOCFunc
	Answer2Func  AOCFunc
	DefaultInput string
}

var solutionRegistry map[AOCKey]AOCSolution

func init() {
	solutionRegistry = make(map[AOCKey]AOCSolution)
	importFuncs()
}

func (s *AOCSolution) Answer(part int) AOCFunc {
	switch part {
	case 1:
		return s.Answer1Func
	case 2:
		return s.Answer2Func
	default:
		return s.Answer1Func
	}
}

func RegisterSolution(year, day int, solution AOCSolution) {
	//this one takes ints to make it easier on me
	k := getKey(year, day)
	solutionRegistry[k] = solution
}

func getSolution(key AOCKey) (AOCSolution, error) {
	f, ok := solutionRegistry[key]
	if !ok {
		e := errors.New("Function Does Not Exist")
		return AOCSolution{}, e
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
