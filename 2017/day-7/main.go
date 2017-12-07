package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

type process struct {
	name     string
	children []*process
	weight   int
	parent   *process
}

var (
	processes = make(map[string]*process)
)

func main() {
	startTime := time.Now()
	inputBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	rootName, amount := getAnswers(string(inputBytes))
	duration := time.Since(startTime)
	fmt.Printf("Part 1: %s\nPart 2: %d\nRuntime: %s\n", rootName, amount, duration)
}

func getAnswers(in string) (string, int) {
	splitInput := strings.Split(in, "\n")
	initProcesses(splitInput)
	initChildren(splitInput)
	root := findRootProcess()
	unbalancedTree := root.findUnbalancedNode()
	balanceAmount := unbalancedTree.balance()
	return root.name, balanceAmount
}

func initProcesses(in []string) {
	for _, x := range in {
		split := strings.Split(x, "->")
		s := strings.TrimSuffix(split[0], " ")
		name := ""
		weight := 0
		fmt.Sscanf(s, "%s (%d)", &name, &weight)
		initProcess(name, weight)
	}
}

func initChildren(in []string) {
	workingSet := []string{}
	for _, x := range in {
		split := strings.Split(x, " ")
		if len(split) > 2 {
			workingSet = append(workingSet, x)
		}
	}
	for _, x := range workingSet {
		pSplit := strings.Split(x, " ")
		parent := pSplit[0]
		parent = strings.TrimPrefix(parent, "\t")
		split := strings.Split(x, "->")
		children := strings.Split(strings.TrimPrefix(split[1], " "), " ")
		for _, y := range children {
			childname := strings.TrimSuffix(y, ",")
			parentProcess := processByName(parent)
			childProcess := processByName(childname)
			parentProcess.addChild(childProcess)
		}
	}
}

func (p *process) addChild(c *process) {
	p.children = append(p.children, c)
	c.parent = p
}

func processByName(in string) *process {
	return processes[in]
}

func findRootProcess() *process {
	for _, x := range processes {
		if x.parent == nil {
			return x
		}
	}
	return nil
}

func initProcess(name string, weight int) {
	p := &process{
		name:   name,
		weight: weight,
	}
	processes[name] = p
}

func (p *process) getChildWeight() int {
	sum := 0
	for _, x := range p.children {
		if len(x.children) > 0 {
			sum += x.getChildWeight()
		}
		sum += x.weight
	}
	return sum
}

func (p *process) getTreeWeight() int {
	sum := 0
	sum += p.weight
	sum += p.getChildWeight()
	return sum
}

func (p *process) findUnbalancedTree() *process {
	trees := make(map[int]int)
	for _, x := range p.children {
		weight := x.getTreeWeight()
		trees[weight]++
	}
	if len(trees) <= 1 {
		return p
	}
	minIndex := 0
	minValue := 0
	for x, y := range trees {
		if minIndex == 0 && minValue == 0 {
			minIndex = x
			minValue = y
		}
		if y < minValue {
			minIndex = x
			minValue = y
		}
	}
	for _, x := range p.children {
		if x.getTreeWeight() == minIndex {
			return x
		}
	}
	return nil
}

func (p *process) findUnbalancedNode() *process {
	x := p.findUnbalancedTree()
	if p == x {
		return x
	}
	return x.findUnbalancedNode()
}

func (p *process) balance() int {
	root := findRootProcess()
	initalweight := p.weight
	for p.weight > 0 {
		p.weight--
		if root.findUnbalancedNode() == root {
			return p.weight
		}
	}
	p.weight = initalweight
	for {
		p.weight++
		if root.findUnbalancedNode() == root {
			return p.weight
		}
	}
}
