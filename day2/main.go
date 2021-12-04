package day2

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
	log.Printf("Running Day 2 part %d", part)
	dataSource := "day2/real.txt"
	if test {
		dataSource = "day2/test.txt"
	}
	if part == 1 {
		return part1(dataSource)
	}
	return part2(dataSource)
}

func part1(dataSource string) error {
	allData := getData(dataSource)
	xCord := 0
	yCord := 0
	for _, movement := range allData {
		switch movement.Direction {
		case "forward":
			xCord += movement.Amount
		case "up":
			yCord -= movement.Amount
		case "down":
			yCord += movement.Amount
		}
	}
	log.Printf("Results: xCord: %d yCord: %d total: %d", xCord, yCord, (xCord * yCord))
	return nil
}

func part2(dataSource string) error {
	allData := getData(dataSource)
	xCord := 0
	yCord := 0
	aim := 0
	for _, movement := range allData {
		switch movement.Direction {
		case "forward":
			xCord += movement.Amount
			yCord += (movement.Amount * aim)
		case "up":
			aim -= movement.Amount
		case "down":
			aim += movement.Amount
		}
	}
	log.Printf("Results Part 2: xCord: %d yCord: %d total: %d", xCord, yCord, (xCord * yCord))
	return nil
}

func getData(dataSource string) []Movement {
	movements := []Movement{}
	allData := utils.SplitByLine(utils.ReadInputFile(dataSource))
	for _, data := range allData {
		if data == "" {
			continue
		}
		data = strings.TrimSpace(data)
		pieces := strings.Split(data, " ")
		amount, _ := strconv.Atoi(pieces[1])
		movements = append(movements, Movement{
			Direction: pieces[0],
			Amount:    amount,
		})
	}
	return movements
}
