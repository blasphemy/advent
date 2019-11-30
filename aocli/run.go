package main

import "github.com/urfave/cli/v2"

import "advent/aocommon"

import "fmt"

var runCMD = &cli.Command{
	Name:      "run",
	Usage:     "run a solution with default input",
	Action:    runAction,
	ArgsUsage: "<year> <day> [<part>]",
	Aliases: []string{
		"r",
		"ru",
	},
}

func runAction(c *cli.Context) error {
	args := c.Args()
	if args.Len() < 2 {
		return cli.ShowCommandHelp(c, "run")
	}
	ys := args.First()
	ds := args.Get(1)
	ps := "0"
	if args.Len() > 2 {
		ps = args.Get(2)
	}
	res, err := runParts(ys, ds, ps)
	if err != nil {
		return err
	}
	for x, y := range res {
		if ps != "1" && ps != "2" {
			fmt.Printf("Part %d\n", x+1)
		}
		fmt.Printf("Answer: %s\n", y.Answer)
		fmt.Printf("Execution Time: %s\n", y.ExecutionTime)
	}
	return nil
}

func runParts(ys, ds, ps string) ([]aocommon.ExecutionResults, error) {
	switch ps {
	case "1", "2":
		res, err := aocommon.ExecuteDefault(ys, ds, ps)
		if err != nil {
			return []aocommon.ExecutionResults{}, err
		}
		return []aocommon.ExecutionResults{res}, nil
	default:
		res1, err := aocommon.ExecuteDefault(ys, ds, "1")
		if err != nil {
			return []aocommon.ExecutionResults{}, err
		}
		res2, err := aocommon.ExecuteDefault(ys, ds, "2")
		if err != nil {
			return []aocommon.ExecutionResults{}, err
		}
		return []aocommon.ExecutionResults{
			res1,
			res2,
		}, nil
	}
}
