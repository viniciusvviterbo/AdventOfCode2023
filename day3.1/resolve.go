package main

import (
	"bufio"
	"fmt"
	"os"
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
	for idx, char := range line {
		if !unicode.IsDigit(char) && tmpStart != -1 && tmpEnd != -1 { // this is the end of a number
			newNumber := NumberPosition{
				line:     lineIdx,
				startCol: tmpStart,
				endCol:   tmpEnd,
			}
			numbersFound = append(numbersFound, newNumber)
			tmpStart = -1
			tmpEnd = -1
		}

		if !unicode.IsDigit(char) && tmpStart == -1 { // hasnt found a number yet
			continue
		}

		if tmpStart != -1 { // this char is a continuation of a number
			tmpNumber = fmt.Sprintf("%s%c", tmpNumber, char)
			tmpEnd = idx
		} else { // this char is the beggining of a number
			tmpStart = idx
			tmpNumber = fmt.Sprintf("%s%c", tmpNumber, char)
			tmpEnd = idx
		}
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

func numberIsAdjacentToSymbol(numPosition NumberPosition, symsPositions []SymbolPosition) bool {
	numberIsAdjacentToSymbol := false

	for _, sym := range symsPositions {
		if numPosition.line-1 <= sym.line &&
			sym.line <= numPosition.line+1 &&
			numPosition.startCol-1 <= sym.col &&
			sym.col <= numPosition.endCol+1 {
			numberIsAdjacentToSymbol = true
			break
		}
	}

	return numberIsAdjacentToSymbol
}

func getValidNumberPositions(engineMap []string) []NumberPosition {
	numbersPositions := make([]NumberPosition, 0)
	symbolsPositions := make([]SymbolPosition, 0)

	for lineIdx, line := range engineMap {
		numbersPositions = append(numbersPositions, findNumbersInLine(line, lineIdx)...)
		symbolsPositions = append(symbolsPositions, findSymbolsInLine(line, lineIdx)...)
	}

	validNumbers := filter(numbersPositions, func(num NumberPosition) bool {
		return numberIsAdjacentToSymbol(num, symbolsPositions)
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

  fmt.Println(number)
    partNumberSum += number
	}

  fmt.Println(partNumberSum)
	// Check for scanner errors
	if err := johnLennon.Err(); err != nil {
		fmt.Println("Error reading standard input:", err)
	}
}
