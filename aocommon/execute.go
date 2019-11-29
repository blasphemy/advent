package aocommon

import "time"

func ExecuteDefault(year,day,part string) (int,time.Duration,error) {
	f,err := getFunc(year,day,part)
	if err != nil {
		return 0, 0, err
	}
	in := getInput(year,day,part)
	sTime := time.Now()
	ans := f(in)
	finTime := time.Since(sTime)
	return ans, finTime, nil
}

func ExecuteInput(year,day,part,input string) (int,time.Duration,error) {
	f, err := getFunc(year,day,part)
	if err != nil {
		return 0, 0, err
	}
	startTime := time.Now()
	ans := f(input)
	finTime := time.Since(startTime)
	return ans, finTime, nil
}