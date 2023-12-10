package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

type NumberPosition struct {
	line     int
	startCol int
	endCol   int
}

type SymbolPosition struct {
	line int
	col  int
}

func filter[T NumberPosition](slice []T, condition func(T) bool) []T {
	var result []T
	for _, elem := range slice {
		if condition(elem) {
			result = append(result, elem)
		}
	}
	return result
}

func printEngineMap(engineMap []string) {
	fmt.Println("=====   Engine Map   =====")
	for _, line := range engineMap {
		fmt.Println(line)
	}
}

func findNumbersInLine(line string, lineIdx int) []NumberPosition {
	var numbersFound []NumberPosition = make([]NumberPosition, 0)

	tmpStart, tmpEnd := -1, -1
	var tmpNumber string

	saveNumber := func() {
		newNumber := NumberPosition{
			line:     lineIdx,
			startCol: tmpStart,
			endCol:   tmpEnd,
		}
		numbersFound = append(numbersFound, newNumber)
		tmpStart = -1
		tmpEnd = -1
	}

	for idx, char := range line {
		runeIsDigit := unicode.IsDigit(char)

		if runeIsDigit {
			if tmpStart != -1 { // this char is a continuation of a number
				tmpNumber = fmt.Sprintf("%s%c", tmpNumber, char)
				tmpEnd = idx
			} else { // this char is the beggining of a number
				tmpStart = idx
				tmpNumber = fmt.Sprintf("%s%c", tmpNumber, char)
				tmpEnd = idx
			}
		} else if tmpStart == -1 { // hasnt found a number yet
			continue
		} else if tmpEnd != -1 { // this is the end of a number
			saveNumber()
		}
	}

	if unicode.IsDigit(rune(line[len(line)-1])) { // the last digit of this line is part of a number
		saveNumber()
	}

	return numbersFound
}

func findSymbolsInLine(line string, lineIdx int) []SymbolPosition {
	var symbolsFound []SymbolPosition = make([]SymbolPosition, 0)

	for idx, char := range line {
		if !unicode.IsDigit(char) && char != '.' { // symbol found
			newSymbol := SymbolPosition{
				line: lineIdx,
				col:  idx,
			}
			symbolsFound = append(symbolsFound, newSymbol)
		}
	}

	return symbolsFound
}

func numberIsAdjacentToGear(numPosition NumberPosition, symsPositions []SymbolPosition, engineMap []string) bool {
	numberIsAdjacentToGear := false

	for _, sym := range symsPositions {
		if numPosition.line-1 <= sym.line &&
			sym.line <= numPosition.line+1 &&
			numPosition.startCol-1 <= sym.col &&
			sym.col <= numPosition.endCol+1 &&
			engineMap[sym.line][sym.col] == '*' {
			numberIsAdjacentToGear = true
			break
		}
	}

	return numberIsAdjacentToGear
}

func numberOfAdjacentNumbers(symPos SymbolPosition, engineMap []string) int {
	var numberOfAdjacentNumbers int

	re := regexp.MustCompile(`\d+`)

	for i := symPos.line - 1; i < 3; i++ {
		matches := re.FindAllString(engineMap[i][symPos.col-1 : symPos.col+1], -1)
		numberOfAdjacentNumbers += len(matches)
	}

	return numberOfAdjacentNumbers
}

func getValidNumberPositions(engineMap []string) []NumberPosition {
	numbersPositions := make([]NumberPosition, 0)
	symbolsPositions := make([]SymbolPosition, 0)

	for lineIdx, line := range engineMap {
		numbersPositions = append(numbersPositions, findNumbersInLine(line, lineIdx)...)
		symbolsPositions = append(symbolsPositions, findSymbolsInLine(line, lineIdx)...)
	}

	validNumbers := filter(numbersPositions, func(num NumberPosition) bool {
		return numberIsAdjacentToGear(num, symbolsPositions, engineMap) //  && numberOfAdjacentNumbers() == 2
	})


	return validNumbers
}

func main() {
	engineMap := []string{}

	johnLennon := bufio.NewScanner(os.Stdin)
	for johnLennon.Scan() {
		line := johnLennon.Text()
		if line == "" {
			break
		}

		engineMap = append(engineMap, line)
	}

	validNumberPositions := getValidNumberPositions(engineMap)

	var partNumberSum int

	for _, numPos := range validNumberPositions {
		newNumber := engineMap[numPos.line][numPos.startCol : numPos.endCol+1]
		number, err := strconv.Atoi(newNumber)
		if err != nil {
			panic("Not able to convert number")
		}

		partNumberSum += number
	}

	fmt.Println(partNumberSum)
	// Check for scanner errors
	if err := johnLennon.Err(); err != nil {
		fmt.Println("Error reading standard input:", err)
	}
}
