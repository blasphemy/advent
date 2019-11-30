package solutions

type AOCFunc func(string) string

type AOCSolution struct {
	Year         int
	Day          int
	Answer1Func  func(string) string
	Answer2Func  func(string) string
	DefaultInput string
}

func (s *AOCSolution) Answer(part int) AOCFunc {
	switch part {
	case 1:
		return s.Answer1Func
	case 2:
		return s.Answer2Func
	default:
		return s.Answer1Func
	}
}
