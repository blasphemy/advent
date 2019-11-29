package solutions 

type AOCSolution struct {
	Answer1Func  AOCFunc
	Answer2Func  AOCFunc
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