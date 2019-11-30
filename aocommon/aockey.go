package aocommon

import "strconv"

type keyset struct {
	keys []AOCKey
}

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

func (k *keyset) Len() int {
	return len(k.keys)
}

func (k *keyset) Less(i, j int) bool {
	k1 := k.keys[i]
	k2 := k.keys[j]
	if k1.Year < k2.Year {
		return true
	}
	if (k1.Year == k2.Year) && (k2.Day > k1.Day) {
		return true
	}
	if k1.Year > k2.Year {
		return false
	}
	return false
}

func (k *keyset) Swap(i, j int) {
	k1 := k.keys[i]
	k2 := k.keys[j]
	k.keys[i] = k2
	k.keys[j] = k1
}
