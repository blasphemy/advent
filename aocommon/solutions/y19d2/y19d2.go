package y19d2

import "advent/aocommon/solutions"

import "strings"

import "strconv"

const needle = 19690720

var Solution = solutions.AOCSolution{
	Year:         19,
	Day:          2,
	DefaultInput: input,
	Answer1Func:  a1,
	Answer2Func:  a2,
}

func a1(i string) string {
	c, err := newComputer(i)
	if err != nil {
		panic(err)
	}
	//restore program
	c.tape[1] = 12
	c.tape[2] = 2
	c.runProgram()
	return strconv.Itoa(c.tape[0])
}

func a2(in string) string {
	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {
			c, err := newComputer(in)
			if err != nil {
				panic(err)
			}
			c.tape[1] = i
			c.tape[2] = j
			c.runProgram()
			if c.tape[0] == needle {
				return strconv.Itoa(100*i + j)
			}
		}
	}
	return ""
}

func newComputer(input string) (*computer, error) {
	n, err := parseInput(input)
	if err != nil {
		return nil, err
	}
	c := &computer{
		tape: n,
		pos:  0,
	}
	return c, nil
}

type computer struct {
	tape []int
	pos  int
}

func (c *computer) opCode() int {
	return c.tape[c.pos]
}

func (c *computer) arg1() int {
	offset := 1
	p := c.tape[c.pos+offset]
	return c.tape[p]
}

func (c *computer) arg2() int {
	offset := 2
	p := c.tape[c.pos+offset]
	return c.tape[p]
}

func (c *computer) register() int {
	offset := 3
	return c.tape[c.pos+offset]
}

func (c *computer) jump4() {
	c.pos = c.pos + 4
}

func (c *computer) runProgram() {
	for c.tick() {
		c.jump4()
	}
}

func (c *computer) tick() bool {
	switch c.opCode() {
	case 1:
		// add
		w := c.arg1() + c.arg2()
		c.tape[c.register()] = w
		return true
	case 2:
		//mult
		w := c.arg1() * c.arg2()
		c.tape[c.register()] = w
		return true
	case 99:
		return false
	}
	return false
}

func parseInput(i string) ([]int, error) {
	s := strings.Split(i, ",")
	out := []int{}
	for _, x := range s {
		n, err := strconv.Atoi(x)
		if err != nil {
			return []int{}, err
		}
		out = append(out, n)
	}
	return out, nil
}
