package main

import "github.com/urfave/cli/v2"

import "aoc2/aocommon"

import "fmt"

var runCMD = &cli.Command{
	Name: "run",
	Usage: "run a solution with default input",
	Action: runAction,
	ArgsUsage: "<year> <day> <part>",
}

func runAction(c *cli.Context) error {
	args := c.Args()
	if args.Len() < 3 {
		return cli.ShowCommandHelp(c,"run")
	}
	ys := args.First()
	ds := args.Get(1)
	ps := args.Get(2)
	res, err := aocommon.ExecuteDefault(ys,ds,ps)
	if err != nil {
		return err
	}
	fmt.Printf("Answer: %d\n", res.Answer)
	fmt.Printf("Execution Time: %s\n", res.ExecutionTime)
	return nil
}