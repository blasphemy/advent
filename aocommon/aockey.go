package aocommon

import "strconv"

type AOCKey struct {
	Year int
	Day  int
	Part int
}

func getKeyFromStrings(year, day, part string) (AOCKey,error) {
	y, err := strconv.Atoi(year)
	if err != nil {
		return AOCKey{}, err
	}
	d, err := strconv.Atoi(day)
	if err != nil {
		return AOCKey{}, err
	}
	p, err := strconv.Atoi(part)
	if err != nil {
		return AOCKey{}, err
	}
	return getKey(y,d,p), nil
}

func getKey(year,day,part int) AOCKey {
	return AOCKey{
		Year: year,
		Day:  day,
		Part: part,
	}
}