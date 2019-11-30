package aocommon

import "time"

import "strconv"

type ExecutionResults struct {
	Answer        string
	ExecutionTime time.Duration
}

func ExecuteDefault(year, day, part string) (ExecutionResults, error) {
	k, err := getKeyFromStrings(year, day)
	if err != nil {
		return ExecutionResults{}, err
	}
	sol, err := getSolution(k)
	if err != nil {
		return ExecutionResults{}, err
	}
	pint, err := strconv.Atoi(part)
	if err != nil {
		return ExecutionResults{}, err
	}
	sTime := time.Now()
	ans := sol.Answer(pint)(sol.DefaultInput)
	finTime := time.Since(sTime)
	return ExecutionResults{ans, finTime}, nil
}

func ExecuteInput(year, day, part, input string) (ExecutionResults, error) {
	k, err := getKeyFromStrings(year, day)
	if err != nil {
		return ExecutionResults{}, err
	}
	sol, err := getSolution(k)
	if err != nil {
		return ExecutionResults{}, err
	}
	pint, err := strconv.Atoi(part)
	if err != nil {
		return ExecutionResults{}, err
	}
	startTime := time.Now()
	ans := sol.Answer(pint)(input)
	finTime := time.Since(startTime)
	return ExecutionResults{ans, finTime}, nil
}
