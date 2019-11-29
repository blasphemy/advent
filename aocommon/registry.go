package aocommon

import "errors"

var funcRegistry map[AOCKey]AOCFunc
var aocInputs map[AOCKey]string //example inputs

func init() {
	funcRegistry = make(map[AOCKey]AOCFunc)
	aocInputs = make(map[AOCKey]string)
	importFuncs()
}

func registerWithInput(year, day, part int, input string, answerfunc AOCFunc) {
	//this one takes ints to make it easier on me
	k := getKey(year, day, part)
	registerAOCFunc(k, answerfunc)
	registerAOCInput(k, input)
}

func registerAOCFunc(key AOCKey, answerfunc AOCFunc) {
	funcRegistry[key] = answerfunc
}

func registerAOCInput(key AOCKey, input string) {
	aocInputs[key] = input
}

func getFunc(key AOCKey) (AOCFunc, error) {
	f, ok := funcRegistry[key]
	if !ok {
		e := errors.New("Function Does Not Exist")
		return nil, e
	}
	return f, nil
}

func getInput(key AOCKey) string {
	i := aocInputs[key]
	return i
}

func AOCAvailable() []AOCKey {
	keys := []AOCKey{}
	for x := range(funcRegistry) {
		keys = append(keys,x)
	}
	return keys
}
