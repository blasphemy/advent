package aocommon

import "strconv"

type AOCKey struct {
	Year int
	Day  int
}

func getKeyFromStrings(year, day string) (AOCKey, error) {
	y, err := strconv.Atoi(year)
	if err != nil {
		return AOCKey{}, err
	}
	d, err := strconv.Atoi(day)
	if err != nil {
		return AOCKey{}, err
	}
	return getKey(y, d), nil
}

func getKey(year, day int) AOCKey {
	return AOCKey{
		Year: year,
		Day:  day,
	}
}
