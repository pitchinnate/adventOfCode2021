package day6

import (
	"fmt"
	"log"
)

func Run(part int, test bool) error {
	log.Printf("Running Day 6 part %d", part)
	dataSource := "day6/real.txt"
	if test {
		dataSource = "day6/test.txt"
	}
	if part == 1 {
		return part1(dataSource)
	}
	return part2(dataSource)
}

func part1(dataSource string) error {
	//allData := utils.ReadInputFile(dataSource)
	//pieces := strings.Split(allData, ",")
	//numbers := utils.StringsToIntsNormal(pieces)
	numbers := []int{5}
	days := 40

	for i := 0; i < days; i++ {
		addNumbers := 0
		for index, number := range numbers {
			number -= 1
			if number < 0 {
				numbers[index] = 6
				addNumbers += 1
			} else {
				numbers[index] = number
			}
		}
		for adder := 0; adder < addNumbers; adder++ {
			numbers = append(numbers, 8)
		}
		//fmt.Printf("Day %2d: %v \n", i, numbers)
	}

	log.Println("Results: ", len(numbers))
	return nil
}

func part2(dataSource string) error {
	//allData := utils.ReadInputFile(dataSource)
	//pieces := strings.Split(allData, ",")
	//orgNumbers := utils.StringsToIntsNormal(pieces)
	orgNumbers := []int{1, 2, 3, 4, 5, 6}
	totalFish := 0
	days := 40

	cacheAnswer := map[int]int{}

	for fishNumber, orgNumber := range orgNumbers {
		cached, ok := cacheAnswer[orgNumber]
		if !ok {
			fishCount := 1
			//fmt.Printf("+ add %d fish on day %d \n", 1, 0)
			additions := 1 + int((days-orgNumber)/7)
			//fmt.Printf("Fish %2d: %2d time for %2d additions \n", fishNumber, orgNumber, additions)
			for i := 0; i < additions; i++ {
				fishCount += 1
				currentDay := orgNumber + (i * 7)
				//fmt.Printf("+ add %d fish on day %d p1 \n", 1, currentDay)
				//fmt.Printf("  Added fish on day: %2d\n", currentDay)
				children := calcChildren(currentDay, 1, days)
				//fmt.Printf("    Added fish on: %2d children: %2d \n", currentDay, children)
				fishCount += children
				//fmt.Printf(" %2d fish added \n", fishCount)
			}
			fmt.Printf(" %3d fish added %d \n", fishNumber, fishCount)
			totalFish += fishCount
			cacheAnswer[orgNumber] = fishCount
		} else {
			fmt.Printf(" %3d fish added from cache %d \n", fishNumber, cached)
			totalFish += cached
		}
	}

	log.Println("Results: ", totalFish)
	return nil
}

func calcChildren(currentDay int, level int, totalDays int) int {
	level += 1
	//spaces := make([]string, level*2)
	//padding := strings.Join(spaces, " ")
	children := 0
	additionDay := currentDay + 9
	if additionDay < totalDays {
		//fmt.Printf("%s calc children for %d remaining days: day %d \n", padding, remainingDays, currentDay)

		// account for first addition after 9 days
		//fmt.Printf("%s + add %d fish on day %d p1 \n", padding, 1, currentDay)
		children += 1
		newChildren := calcChildren(additionDay, level, totalDays)
		if newChildren > 0 {
			children += newChildren
			//fmt.Printf("%s + add %d fish on day %d p2 \n", padding, newChildren, currentDay)
			//fmt.Printf("%s adding child with %d remaining days, remaining children created %d \n", padding, additionDay, newChildren)
		}

		// how many more additions will i make after 7 days each
		additions := int((totalDays - additionDay) / 7)
		if additions > 0 {
			//fmt.Printf("%s %d additions remaing after first 9\n", padding, additions)
			for i := 1; i <= additions; i++ {
				currentDay2 := additionDay + (i * 7)
				if currentDay2 >= totalDays {
					continue
				}
				children += 1
				newChildren2 := calcChildren(currentDay2, level, totalDays)
				//fmt.Printf("%s + add %d fish on day %d p3 \n", padding, 1, currentDay2)
				if newChildren2 > 0 {
					//fmt.Printf("%s adding child with %d remaining days, remaining children created %d \n", padding, newDay, newChildren2)
					children += newChildren2
					//fmt.Printf("%s + add %d fish on day %d p4 \n", padding, newChildren2, currentDay2)
				}
			}
		} else {
			//fmt.Printf("%s no more additions\n", padding)
		}
		//fmt.Printf("%s total children added: %d\n", padding, children)
	}
	return children
}
