package day3

import (
	"2021/utils"
	"log"
	"strconv"
	"strings"
)

type Movement struct {
	Direction string
	Amount    int
}

func Run(part int, test bool) error {
	log.Printf("Running Day 3 part %d", part)
	dataSource := "day3/real.txt"
	if test {
		dataSource = "day3/test.txt"
	}
	if part == 1 {
		return part1(dataSource)
	}
	return part2(dataSource)
}

func part1(dataSource string) error {
	allData := utils.SplitByLine(utils.ReadInputFile(dataSource))
	spaces := len(strings.TrimSpace(allData[0]))
	sumBits := make([]int, spaces)

	for _, line := range allData {
		if line == "" {
			continue
		}
		//log.Println("line: ", line)
		for i := 0; i < spaces; i++ {
			currentBit := string(line[i])
			//log.Println("currentBit: ", currentBit)
			val, _ := strconv.Atoi(currentBit)
			sumBits[i] += val
		}
	}
	//log.Println("sum bits: ", sumBits)

	numData := float64(len(allData))
	avgBits := make([]string, spaces)
	avgBits2 := make([]string, spaces)
	for i, bits := range sumBits {
		averageBits := float64(bits) / numData
		//log.Println("bits: ", bits, numData, averageBits)
		if averageBits >= .5 {
			avgBits[i] = "0"
			avgBits2[i] = "1"
		} else {
			avgBits[i] = "1"
			avgBits2[i] = "0"
		}
	}
	result := intsArrayToInt(avgBits)
	result2 := intsArrayToInt(avgBits2)
	log.Printf("Results: ", avgBits, result, avgBits2, result2, (result * result2))
	return nil
}

func part2(dataSource string) error {
	allData := utils.SplitByLine(utils.ReadInputFile(dataSource))
	cleanData := []string{}
	for _, line := range allData {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		cleanData = append(cleanData, line)
	}

	result, _ := strconv.ParseInt(processPart2(cleanData, "0", "1"), 2, 64)
	result2, _ := strconv.ParseInt(processPart2(cleanData, "1", "0"), 2, 64)

	log.Printf("Results Part 2: ", result, result2, result*result2)
	return nil
}

func processPart2(allData []string, lookFor string, lookFor2 string) string {
	spaces := len(strings.TrimSpace(allData[0]))

	for i := 0; i < spaces; i++ {
		var total float64 = 0
		count := len(allData)

		if count == 1 {
			break
		}

		for _, line := range allData {
			currentBit, _ := strconv.Atoi(string(line[i]))
			total += float64(currentBit)
		}
		avg := total / float64(count)
		lookingFor := lookFor
		if avg >= .5 {
			lookingFor = lookFor2
		}

		newData := []string{}
		for _, line := range allData {
			if string(line[i]) == lookingFor {
				newData = append(newData, line)
			}
		}
		allData = newData
	}
	return allData[0]
}

func intsArrayToInt(vals []string) int64 {
	binaryVal := strings.Join(vals, "")
	result, _ := strconv.ParseInt(binaryVal, 2, 64)
	return result
}
