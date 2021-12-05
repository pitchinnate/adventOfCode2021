package day5

import (
	"2021/utils"
	"log"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Line struct {
	Point1 Point
	Point2 Point
}

func (l *Line) GetPoints(includeDiagonals bool) []Point {
	points := []Point{l.Point1, l.Point2}
	if l.Point1.Y == l.Point2.Y {
		min, max := utils.FindMinAndMax([]int{l.Point1.X, l.Point2.X})
		if max-min > 1 {
			for i := min + 1; i < max; i++ {
				points = append(points, Point{
					X: i,
					Y: l.Point1.Y,
				})
			}
		}
	} else if l.Point1.X == l.Point2.X {
		min, max := utils.FindMinAndMax([]int{l.Point1.Y, l.Point2.Y})
		if max-min > 1 {
			for i := min + 1; i < max; i++ {
				points = append(points, Point{
					X: l.Point1.X,
					Y: i,
				})
			}
		}
	} else {
		minX, maxX := utils.FindMinAndMax([]int{l.Point1.X, l.Point2.X})
		diffx := maxX - minX
		minY, maxY := utils.FindMinAndMax([]int{l.Point1.X, l.Point2.X})
		diffy := maxY - minY

		diffX := l.Point1.X - l.Point2.X
		diffY := l.Point1.Y - l.Point2.Y

		incX := -1
		incY := -1
		if diffX < 0 {
			incX = 1
		}
		if diffY < 0 {
			incY = 1
		}

		if !includeDiagonals || diffy != diffx {
			return []Point{}
		}
		if maxX-minX > 1 {
			for y := 1; y < diffx; y++ {
				points = append(points, Point{
					X: l.Point1.X + (y * incX),
					Y: l.Point1.Y + (y * incY),
				})
			}
			//log.Println("Diaganol Line: ", points)
		}
	}
	return points
}

func Run(part int, test bool) error {
	log.Printf("Running Day 4 part %d", part)
	dataSource := "day5/real.txt"
	if test {
		dataSource = "day5/test.txt"
	}
	if part == 1 {
		return part1(dataSource)
	}
	return part2(dataSource)
}

func part1(dataSource string) error {
	allData := utils.SplitByLine(utils.ReadInputFile(dataSource))
	lines, max := parseData(allData)
	grid := makeGrid(max, 0)

	for _, line := range lines {
		linePoints := line.GetPoints(false)
		for _, lp := range linePoints {
			grid[lp.Y][lp.X] += 1
		}
	}

	log.Println("Results: ")
	finalCount := 0
	for y := 0; y <= max; y++ {
		//fmt.Println(grid[y])
		for x := 0; x <= max; x++ {
			if grid[y][x] > 1 {
				finalCount += 1
			}
		}
	}
	log.Println("Results: ", finalCount)
	return nil
}

func part2(dataSource string) error {
	allData := utils.SplitByLine(utils.ReadInputFile(dataSource))
	lines, max := parseData(allData)
	grid := makeGrid(max, 0)

	for _, line := range lines {
		linePoints := line.GetPoints(true)
		for _, lp := range linePoints {
			grid[lp.Y][lp.X] += 1
		}
	}

	log.Println("Results: ")
	finalCount := 0
	for y := 0; y <= max; y++ {
		//fmt.Println(grid[y])
		for x := 0; x <= max; x++ {
			if grid[y][x] > 1 {
				finalCount += 1
			}
		}
	}
	log.Println("Results: ", finalCount)
	return nil
}

func makeGrid(max int, fill int) [][]int {
	grid := make([][]int, max+1)
	for y := 0; y <= max; y++ {
		grid[y] = make([]int, max+1)
		for x := 0; x <= max; x++ {
			grid[y][x] = fill
		}
	}

	return grid
}

func parseData(lines []string) ([]Line, int) {
	newLines := []Line{}
	allNumbers := []int{}
	for _, line := range lines {
		cleanLine := strings.TrimSpace(line)
		if cleanLine == "" {
			continue
		}
		pieces := strings.Split(cleanLine, "->")
		point1 := strings.Split(strings.TrimSpace(pieces[0]), ",")
		point1Cords := utils.StringsToIntsNormal(point1)
		point2 := strings.Split(strings.TrimSpace(pieces[1]), ",")
		point2Cords := utils.StringsToIntsNormal(point2)

		newLines = append(newLines, Line{
			Point1: Point{
				X: point1Cords[0],
				Y: point1Cords[1],
			},
			Point2: Point{
				X: point2Cords[0],
				Y: point2Cords[1],
			},
		})
		allNumbers = append(allNumbers, point1Cords...)
		allNumbers = append(allNumbers, point2Cords...)
	}
	_, max := utils.FindMinAndMax(allNumbers)
	return newLines, max
}
