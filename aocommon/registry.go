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

func registerWithInput(year, day int, input string, answer1func AOCFunc, answer2func AOCFunc) {
	//this one takes ints to make it easier on me
	k := getKey(year, day)
	sol := AOCSolution{
		Answer1Func:  answer1func,
		Answer2Func:  answer2func,
		DefaultInput: input,
	}
	solutionRegistry[k] = sol
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
