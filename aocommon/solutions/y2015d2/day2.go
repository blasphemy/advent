package y2015d2

import (
	"advent/aocommon/solutions"
	"log"
	"sort"
	"strconv"
	"strings"
)

var Solution = solutions.AOCSolution{
	Year:         15,
	Day:          2,
	Answer1Func:  getTotalAreaFromInput,
	Answer2Func:  getTotalRibbonFromInput,
	DefaultInput: Challenge,
}

type dimensions struct {
	l int
	w int
	h int
}

func getTotalAreaFromInput(i string) string {
	rows := strings.Split(i, "\n")
	sum := 0
	for _, x := range rows {
		d, err := inputToDimensions(strings.TrimPrefix(x, "\t"))
		if err != nil {
			log.Fatal(err.Error())
		}
		a := d.getArea()
		sum = sum + a
	}
	return strconv.Itoa(sum)
}

func inputToDimensions(i string) (dimensions, error) {
	split := strings.Split(i, "x")
	out := dimensions{}
	l, err := strconv.Atoi(split[0])
	if err != nil {
		return out, err
	}
	w, err := strconv.Atoi(split[1])
	if err != nil {
		return out, err
	}
	h, err := strconv.Atoi(split[2])
	if err != nil {
		return out, err
	}
	out = dimensions{l, w, h}
	return out, nil
}

func (d dimensions) getArea() int {
	area1 := d.l * d.w
	area2 := d.w * d.h
	area3 := d.h * d.l
	areas := []int{area1, area2, area3}
	smallest := 0
	out := 0
	for _, x := range areas {
		if smallest == 0 {
			smallest = x
		}
		if x < smallest {
			smallest = x
		}
		out = out + x*2
	}
	out = out + smallest
	return out
}

func getTotalRibbonFromInput(i string) string {
	rows := strings.Split(i, "\n")
	sum := 0
	for _, x := range rows {
		d, err := inputToDimensions(strings.TrimPrefix(x, "\t"))
		if err != nil {
			log.Fatal(err.Error())
		}
		a := d.getRibbon()
		sum = sum + a
	}
	return strconv.Itoa(sum)
}

func (d dimensions) getRibbon() int {
	s := []int{d.w, d.l, d.h}
	sort.Ints(s)
	r := s[0] + s[0] + s[1] + s[1]
	r = r + d.l*d.w*d.h
	return r
}
