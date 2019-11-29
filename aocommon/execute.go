package aocommon

func ExecuteDefault(year,day,part string) (int,error) {
	f,err := GetFunc(year,day,part)
	if err != nil {
		return 0, err
	}
	in := getInput(year,day,part)
	ans := f(in)
	return ans, nil
}