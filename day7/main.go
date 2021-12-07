package day7

import (
	"2021/utils"
	"log"
	"math"
	"strings"
)

func Run(part int, test bool) error {
	log.Printf("Running Day 7 part %d", part)
	dataSource := "day7/real.txt"
	if test {
		dataSource = "day7/test.txt"
	}
	if part == 1 {
		return part1(dataSource)
	}
	return part2(dataSource)
}

func part1(dataSource string) error {
	allData := utils.ReadInputFile(dataSource)
	pieces := strings.Split(allData, ",")
	numbers := utils.StringsToIntsNormal(pieces)
	min, max := utils.FindMinAndMax(numbers)

	bestScore := math.MaxInt32
	bestSpot := 0

	for i := min; i <= max; i++ {
		moves := 0
		for _, number := range numbers {
			diff := int(math.Abs(float64(i - number)))
			//log.Printf("%d %d %d", i, number, diff)
			moves += diff
		}
		//log.Println("Total moves", i, moves, bestScore)
		if moves < bestScore {
			//log.Println("new best", i, moves)
			bestScore = moves
			bestSpot = i
		}
	}

	log.Println("Results: ", bestScore, bestSpot)
	return nil
}

func part2(dataSource string) error {
	allData := utils.ReadInputFile(dataSource)
	pieces := strings.Split(allData, ",")
	numbers := utils.StringsToIntsNormal(pieces)
	min, max := utils.FindMinAndMax(numbers)

	bestScore := math.MaxInt32
	bestSpot := 0

	for i := min; i <= max; i++ {
		moves := 0
		for _, number := range numbers {
			diff := int(math.Abs(float64(i - number)))
			diff = diff * (diff + 1) / 2
			//log.Printf("%d %d %d", i, number, diff)
			moves += diff
		}
		//log.Println("Total moves", i, moves, bestScore)
		if moves < bestScore {
			//log.Println("new best", i, moves)
			bestScore = moves
			bestSpot = i
		}
	}

	log.Println("Results: ", bestScore, bestSpot)
	return nil
}
