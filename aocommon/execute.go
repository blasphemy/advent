package aocommon

import "time"

func ExecuteDefault(year, day, part string) (int, time.Duration, error) {
	k, err := getKeyFromStrings(year,day,part)
	if err != nil {
		return 0,0, err
	}
	f, err := getFunc(k)
	if err != nil {
		return 0, 0, err
	}
	in := getInput(k)
	sTime := time.Now()
	ans := f(in)
	finTime := time.Since(sTime)
	return ans, finTime, nil
}

func ExecuteInput(year, day, part, input string) (int, time.Duration, error) {
	k, err := getKeyFromStrings(year,day,part)
	if err != nil {
		return 0,0,err
	}
	f, err := getFunc(k)
	if err != nil {
		return 0, 0, err
	}
	startTime := time.Now()
	ans := f(input)
	finTime := time.Since(startTime)
	return ans, finTime, nil
}
