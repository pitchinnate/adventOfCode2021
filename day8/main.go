package day8

import (
	"2021/utils"
	"log"
	"strings"
)

type Number struct {
	Number  int
	Length  int
	Letters []string
}

func (n *Number) Match(letters []string) bool {
	if len(letters) != n.Length {
		return false
	}
	log.Printf("Number %d compare %v to %v", n.Number, n.Letters, letters)
	counts := letterCounter(letters, n.Letters, []string{}, []string{})
	log.Print("counts: ", counts)
	singleLetters := getWhereLetterCount(counts, 1)
	log.Print("singleLetters: ", singleLetters)
	return len(singleLetters) == 0
}

func Run(part int, test bool) error {
	log.Printf("Running Day 7 part %d", part)
	dataSource := "day8/real.txt"
	if test {
		dataSource = "day8/test.txt"
	}
	if part == 1 {
		return part1(dataSource)
	}
	return part2(dataSource)
}

func part1(dataSource string) error {
	allData := utils.SplitByLine(utils.ReadInputFile(dataSource))
	count := 0

	for _, line := range allData {
		split := strings.Split(line, "|")
		if len(split) < 2 {
			continue
		}
		pieces := strings.Split(split[1], " ")

		for _, piece := range pieces {
			cleaned := strings.TrimSpace(piece)
			length := len(cleaned)
			if length == 2 || length == 4 || length == 3 || length == 7 {
				count++
			}
		}
	}

	log.Println("Results: ", count)
	return nil
}

func part2(dataSource string) error {
	allData := utils.SplitByLine(utils.ReadInputFile(dataSource))
	//count := 0

	numbers := []Number{
		Number{0, 6, []string{"a", "b", "c", "e", "f", "g"}},
		Number{1, 2, []string{"c", "f"}},
		Number{2, 5, []string{"a", "c", "d", "e", "g"}},
		Number{3, 5, []string{"a", "c", "d", "f", "g"}},
		Number{4, 4, []string{"b", "c", "d", "f"}},
		Number{5, 5, []string{"a", "b", "d", "f", "g"}},
		Number{6, 6, []string{"a", "b", "d", "e", "f", "g"}},
		Number{7, 3, []string{"a", "c", "f"}},
		Number{8, 7, []string{"a", "b", "c", "d", "e", "f", "g"}},
		Number{9, 6, []string{"a", "b", "c", "d", "f", "g"}},
	}
	log.Print(numbers)

	for _, line := range allData {
		combos := map[int][]string{}
		combosSplit := map[int]map[int][]string{}
		translations := map[string]string{
			"a": "",
			"b": "",
			"c": "",
			"d": "",
			"e": "",
			"f": "",
			"g": "",
		}

		split := strings.Split(line, "|")
		if len(split) < 2 {
			continue
		}
		pieces := strings.Split(strings.TrimSpace(split[0]), " ")

		for _, piece := range pieces {
			cleaned := strings.TrimSpace(piece)
			length := len(cleaned)
			if length == 0 {
				continue
			}
			val, ok := combos[length]
			if !ok {
				combos[length] = []string{cleaned}
			} else {
				val = append(val, cleaned)
				combos[length] = val
			}
		}

		for index, com := range combos {
			for index2, val := range com {
				_, ok := combosSplit[index]
				if !ok {
					combosSplit[index] = map[int][]string{}
				}

				combosSplit[index][index2] = strings.Split(val, "")
			}
		}

		// can easily get letter a by comparing 1 and 7
		diffLettersIn17 := letterCounter(combosSplit[3][0], combosSplit[2][0], []string{}, []string{})
		translations["a"] = getWhereLetterCount(diffLettersIn17, 1)[0]

		diffLettersIn1478 := letterCounter(combosSplit[3][0], combosSplit[2][0], combosSplit[4][0], combosSplit[7][0])
		//log.Println("diff letters 1, 4, 7, 8:", diffLettersIn1478)
		EorG := getWhereLetterCount(diffLettersIn1478, 1)
		//log.Println("E or G is: ", EorG)

		diffLettersIn069 := letterCounter(combosSplit[6][0], combosSplit[6][1], combosSplit[6][2], []string{})
		//log.Println("diff letters 0 6 9:", diffLettersIn069) // 2s = d, c, e | 3s = a, b, f, g
		AorBorForG := getWhereLetterCount(diffLettersIn069, 3)

		diffLettersIn235 := letterCounter(combosSplit[5][0], combosSplit[5][1], combosSplit[5][2], []string{})
		//log.Println("diff letters 2 3 5:", diffLettersIn235) // 1s = b, e | 2s = c, f | 3s = a, d, g
		BorE := getWhereLetterCount(diffLettersIn235, 1)
		AorDorG := getWhereLetterCount(diffLettersIn235, 3)
		//log.Println("B or E is: ", BorE)

		diff := letterCounter(BorE, EorG, []string{}, []string{})
		translations["e"] = getWhereLetterCount(diff, 2)[0]

		diff = letterCounter(EorG, []string{translations["e"]}, []string{}, []string{})
		translations["g"] = getWhereLetterCount(diff, 1)[0]

		diff = letterCounter(BorE, []string{translations["e"]}, []string{}, []string{})
		translations["b"] = getWhereLetterCount(diff, 1)[0]

		diff = letterCounter(AorDorG, []string{translations["a"], translations["g"]}, []string{}, []string{})
		translations["d"] = getWhereLetterCount(diff, 1)[0]

		diff = letterCounter(AorBorForG, []string{translations["a"], translations["b"], translations["g"]}, []string{}, []string{})
		translations["f"] = getWhereLetterCount(diff, 1)[0]

		diff = letterCounter(combosSplit[3][0], []string{translations["a"], translations["f"]}, []string{}, []string{})
		translations["c"] = getWhereLetterCount(diff, 1)[0]

		//log.Println(lettersIn1, lettersIn4, lettersIn7, lettersIn8)
		log.Println("Translations: ", translations)
		log.Println("\n\n")

		pieces2 := strings.Split(strings.TrimSpace(split[1]), " ")
		log.Println("pieces2", pieces2)
		myNumbers := []Number{}
		for _, piece := range pieces2 {
			cleaned := strings.TrimSpace(piece)
			length := len(cleaned)
			if length == 0 {
				continue
			}
			translated := translateLetters(strings.Split(cleaned, ""), translations)
			log.Println("Original", cleaned, "translated", translated)
			log.Println("===========================================")
			for _, number := range numbers {
				if number.Match(translated) {
					myNumbers = append(myNumbers, number)
					break
				}
			}
			log.Println("\n")

		}
		log.Println("numbers: ", myNumbers)

		break
	}

	log.Println("Results: ")
	return nil
}

func translateLetters(orignal []string, translations map[string]string) []string {
	letters := []string{}

	for _, letter := range orignal {
		letters = append(letters, translations[letter])
	}

	return letters
}

func letterCounter(list1 []string, list2 []string, list3 []string, list4 []string) map[string]int {
	letterCounts := map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
		"e": 0,
		"f": 0,
		"g": 0,
	}
	for _, letter1 := range list1 {
		letterCounts[letter1] += 1
	}
	for _, letter1 := range list2 {
		letterCounts[letter1] += 1
	}
	if len(list3) > 0 {
		for _, letter1 := range list3 {
			letterCounts[letter1] += 1
		}
	}
	if len(list4) > 0 {
		for _, letter1 := range list4 {
			letterCounts[letter1] += 1
		}
	}
	return letterCounts
}

func differentLetter(list1 []string, list2 []string, list3 []string, list4 []string) []string {
	letterCounts := letterCounter(list1, list2, list3, list4)
	log.Println("letter counts:", letterCounts)

	diff := []string{}
	for letter, count := range letterCounts {
		if count == 1 {
			diff = append(diff, letter)
		}
	}

	return diff
}

func getWhereLetterCount(letterCounts map[string]int, val int) []string {
	lettersFound := []string{}
	for letter, count := range letterCounts {
		if count == val {
			lettersFound = append(lettersFound, letter)
		}
	}
	return lettersFound
}
