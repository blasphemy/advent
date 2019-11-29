package aocommon

import "time"

type ExecutionResults struct {
	Answer        int
	ExecutionTime time.Duration
}

func ExecuteDefault(year, day, part string) (ExecutionResults, error) {
	k, err := getKeyFromStrings(year, day, part)
	if err != nil {
		return ExecutionResults{}, err
	}
	f, err := getFunc(k)
	if err != nil {
		return ExecutionResults{}, err
	}
	in := getInput(k)
	sTime := time.Now()
	ans := f(in)
	finTime := time.Since(sTime)
	return ExecutionResults{ans, finTime}, nil
}

func ExecuteInput(year, day, part, input string) (ExecutionResults, error) {
	k, err := getKeyFromStrings(year, day, part)
	if err != nil {
		return ExecutionResults{}, err
	}
	f, err := getFunc(k)
	if err != nil {
		return ExecutionResults{}, err
	}
	startTime := time.Now()
	ans := f(input)
	finTime := time.Since(startTime)
	return ExecutionResults{ans, finTime}, nil
}
