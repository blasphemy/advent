package aocommon

import "time"

func ExecuteDefault(year,day,part string) (int,error, time.Duration) {
	f,err := GetFunc(year,day,part)
	if err != nil {
		return 0, err, time.Duration(0)
	}
	in := getInput(year,day,part)
	sTime := time.Now()
	ans := f(in)
	finTime := time.Since(sTime)
	return ans, nil, finTime
}