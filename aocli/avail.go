package main

import "github.com/urfave/cli/v2"
import "fmt"

import "aoc2/aocommon"

var availCMD = &cli.Command{
	Name:   "avail",
	Usage:  "Show available programs",
	Action: availAction,
}

func availAction(c *cli.Context) error {
	fmt.Println("Getting available solutions")
	avail := aocommon.AOCAvailable()
	fmt.Printf("%d Solutions available\n", len(avail))
	for _, y := range avail {
		fmt.Printf("Year %d - Day %d\n", y.Year, y.Day)
	}
	return nil
}
