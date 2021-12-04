package day4

import (
	"2021/utils"
	"fmt"
	"log"
	"strings"
)

type Board struct {
	Index          int
	Numbers        [5][5]int
	Marks          [5][5]int
	LastNumber     int
	MarkedComplete bool
}

func (b *Board) Winning() bool {
	for i := 0; i < 5; i++ {
		rowTotal := 0
		columnTotal := 0
		for j := 0; j < 5; j++ {
			rowTotal += b.Marks[i][j]
			columnTotal += b.Marks[j][i]
		}
		//log.Println(rowTotal, columnTotal)
		if rowTotal == 5 || columnTotal == 5 {
			return true
		}
	}
	return false
}

func (b *Board) Score() int {
	total := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.Marks[i][j] == 0 {
				total += b.Numbers[i][j]
			}
		}
	}
	log.Println("score: ", total)
	return total * b.LastNumber
}

func (b *Board) CheckForNumber(number int) {
	b.LastNumber = number
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			//currentNumber := b.Numbers[i][j]
			//log.Println(currentNumber, number, i, j)
			if b.Numbers[i][j] == number {
				//log.Println("found number")
				b.Marks[i][j] = 1
			}
		}
	}
}

func (b *Board) Display() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%02v | %01v \n", b.Numbers[i], b.Marks[i])
	}
	fmt.Print("\n\n")
}

func Run(part int, test bool) error {
	log.Printf("Running Day 4 part %d", part)
	dataSource := "day4/real.txt"
	if test {
		dataSource = "day4/test.txt"
	}
	if part == 1 {
		return part1(dataSource)
	}
	return part2(dataSource)
}

func part1(dataSource string) error {
	allData := utils.SplitByLine(utils.ReadInputFile(dataSource))
	numbers := utils.StringsToInts(strings.Split(allData[0], ","))

	boards := []Board{}
	count := 0
	for i := 2; i < len(allData); i += 6 {
		board := Board{
			Index: count,
		}
		for j := 0; j < 5; j++ {
			numbers := BoardLine(allData[i+j])
			board.Numbers[j] = numbers
		}
		boards = append(boards, board)
		count += 1
	}

	var winningBoard Board
	for _, number := range numbers {
		found := false
		for i, board := range boards {
			board.CheckForNumber(int(number))
			if board.Winning() {
				winningBoard = board
				found = true
				break
			}
			boards[i] = board
		}
		if found {
			break
		}
	}

	log.Printf("Results: ", winningBoard.Score())
	winningBoard.Display()
	return nil
}

func part2(dataSource string) error {
	allData := utils.SplitByLine(utils.ReadInputFile(dataSource))
	numbers := utils.StringsToInts(strings.Split(allData[0], ","))

	boards := []Board{}
	count := 0
	for i := 2; i < len(allData); i += 6 {
		board := Board{
			Index: count,
		}
		for j := 0; j < 5; j++ {
			numbers := BoardLine(allData[i+j])
			board.Numbers[j] = numbers
		}
		boards = append(boards, board)
		count += 1
	}

	var winningBoard Board
	winningBoardCount := 0
	for _, number := range numbers {
		for i, board := range boards {
			if board.MarkedComplete {
				continue
			}
			board.CheckForNumber(int(number))
			if board.Winning() {
				winningBoardCount += 1
				winningBoard = board
				board.MarkedComplete = true
			}
			boards[i] = board
		}
		log.Println("Winning Boards: ", winningBoardCount, len(boards))
		if winningBoardCount == len(boards) {
			break
		}
	}

	log.Printf("Results Part 2: ", winningBoard.Index, winningBoard.Score())
	winningBoard.Display()
	return nil
}

func BoardLine(line string) [5]int {
	var lineNumbers [5]int
	numbers := utils.StringsToInts(strings.Split(line, " "))
	for i := 0; i < 5; i++ {
		lineNumbers[i] = int(numbers[i])
	}
	return lineNumbers
}
