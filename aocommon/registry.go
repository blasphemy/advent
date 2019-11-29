package aocommon

import "errors"

type AOCSolution struct {
	AnswerFunc AOCFunc
	DefaultInput string
}

var solutionRegistry map[AOCKey]AOCSolution

func init() {
	solutionRegistry = make(map[AOCKey]AOCSolution)
	importFuncs()
}

func registerWithInput(year, day, part int, input string, answerfunc AOCFunc) {
	//this one takes ints to make it easier on me
	k := getKey(year, day, part)
	sol := AOCSolution {
		AnswerFunc: answerfunc,
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
	for x := range(solutionRegistry) {
		keys = append(keys,x)
	}
	return keys
}
