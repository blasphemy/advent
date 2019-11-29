package main

import "os"

import "log"

import "aoc2/aocommon"

func main() {
	aocommon.InitRegistry()
	arg := os.Args[1:]
	if len(arg) < 3 {
		log.Println("Please specify an exercise to run.")
		log.Println("./aocli <year> <day> <part>")
		log.Println("./aocli 2015 1 1")
		os.Exit(1)
	}
	year := arg[0]
	day := arg[1]
	part := arg[2]
	ans, dur, err := aocommon.ExecuteDefault(year, day, part)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	log.Printf("Answer: %d", ans)
	log.Printf("Finished in: %s", dur)
}
