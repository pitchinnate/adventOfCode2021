package day1

import (
	"2021/utils"
	"log"
	"math"
)

func Run(part int, test bool) error {
	log.Printf("Running Day 1 part %d", part)
	dataSource := "day1/real.txt"
	if test {
		dataSource = "day1/test.txt"
	}
	if part == 1 {
		return part1(dataSource)
	}
	return part2(dataSource)
}

func part1(dataSource string) error {
	allData := utils.StringsToInts(utils.SplitByLine(utils.ReadInputFile(dataSource)))
	increases := findIncreases(allData)
	log.Println("Increases: ", increases)
	return nil
}

func part2(dataSource string) error {
	allData := utils.StringsToInts(utils.SplitByLine(utils.ReadInputFile(dataSource)))
	modifiedData := []int64{}
	length := len(allData)
	for index, data := range allData {
		if index > length-3 {
			break
		}
		modifiedData = append(modifiedData, data+allData[index+1]+allData[index+2])
	}
	increases := findIncreases(modifiedData)
	log.Println("Increases Part 2: ", increases)
	return nil
}

func findIncreases(allData []int64) int {
	increases := 0
	var previousVal int64 = math.MaxInt64
	for _, currentValue := range allData {
		if currentValue > previousVal {
			increases += 1
		}
		previousVal = currentValue
	}
	return increases
}
