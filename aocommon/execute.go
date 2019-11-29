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
	sol, err := getSolution(k)
	if err != nil {
		return ExecutionResults{}, err
	}
	sTime := time.Now()
	ans := sol.AnswerFunc(sol.DefaultInput)
	finTime := time.Since(sTime)
	return ExecutionResults{ans, finTime}, nil
}

func ExecuteInput(year, day, part, input string) (ExecutionResults, error) {
	k, err := getKeyFromStrings(year, day, part)
	if err != nil {
		return ExecutionResults{}, err
	}
	sol, err := getSolution(k)
	if err != nil {
		return ExecutionResults{}, err
	}
	startTime := time.Now()
	ans := sol.AnswerFunc(input)
	finTime := time.Since(startTime)
	return ExecutionResults{ans, finTime}, nil
}
