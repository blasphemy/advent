package main

import "github.com/urfave/cli/v2"

import "os"

import "log"

func main() {
	app := &cli.App{
		Name:  "aocli",
		Usage: "Advent Of Code CLI Runner",
		Commands: []*cli.Command{
			runCMD,
			availCMD,
		},
		Version: "0.0.1",
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
