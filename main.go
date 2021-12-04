package main

import (
	"2021/day1"
	"2021/day2"
	"flag"
	"log"
)

func main() {
	log.Println("Advent of Code 2021")
	day := flag.Int("day", 1, "Select the day # you want to run")
	part := flag.Int("part", 1, "Select if you want to run part 1 or 2")
	testVar := flag.Bool("test", false, "Use test data")
	flag.Parse()

	dayNumber := *day
	runPart := *part
	test := *testVar
	var err error

	switch dayNumber {
	case 1:
		err = day1.Run(runPart, test)
	case 2:
		err = day2.Run(runPart, test)
	}

	if err != nil {
		log.Panicln("Error:", err)
	}
}
