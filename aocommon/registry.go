package aocommon

import "errors"

import "strconv"

var funcRegistry map[string]AOCFunc
var aocInputs map[string]string //example inputs

func InitRegistry() {
	funcRegistry = make(map[string]AOCFunc)
	aocInputs = make(map[string]string)
	importFuncs()
}

func registerWithInput(year,day,part int,input string, answerfunc AOCFunc) {
	//this one takes ints to make it easier on me
	ys := strconv.Itoa(year)
	ds := strconv.Itoa(day)
	ps := strconv.Itoa(part)
	registerAOCFunc(ys,ds,ps, answerfunc)
	registerAOCInput(ys,ds,ps,input)
}

func registerAOCFunc(year,day,part string, answerfunc AOCFunc) {
	key := getKey(year,day,part)
	funcRegistry[key] = answerfunc
}

func registerAOCInput(year,day,part,input string) {
	key := getKey(year,day,part)
	aocInputs[key] = input
}

func getFunc(year,day,part string) (AOCFunc,error) {
	key := getKey(year,day,part)
	f, ok := funcRegistry[key]
	if !ok {
		e := errors.New("Function Does Not Exist")
		return nil, e
	}
	return f, nil
}

func getKey(year,day,part string) string {
	return year + "-" + day + "-" + part
}

func getInput(year,day,part string) string {
	key := getKey(year,day,part)
	i := aocInputs[key]
	return i
}