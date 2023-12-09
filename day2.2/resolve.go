package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const TotalRed = 12
const TotalGreen = 13
const TotalBlue = 14

func treatLine(line string) string {
	substitutionMap := map[string]string{
		"Game":  "",
		":":     "",
		",":     "",
		";":     "",
		"red":   "r",
		"green": "g",
		"blue":  "b",
	}

	var newString string = strings.Clone(line)

	for key, val := range substitutionMap {
		newString = strings.Replace(
			newString,
			key,
			val,
			-1,
		)
	}

	return newString
}

func valueIsPossible(currValue int, currColor string) bool {
	var valueIsPossible bool
	switch currColor {
	case "r":
		valueIsPossible = currValue <= TotalRed
	case "g":
		valueIsPossible = currValue <= TotalGreen
	case "b":
		valueIsPossible = currValue <= TotalBlue
	}
	return valueIsPossible
}

func lineIsPossible(line string) bool {
	var currValue int

	for idx, field := range strings.Fields(line) {
		if idx == 0 { // field is ID
			continue
		}
		if idx%2 != 0 { // field is a quant
			number, err := strconv.Atoi(field)
			if err != nil {
				panic("Not able to convert number")
			}
			currValue = number
		} else { // field is a color
			currColor := field
			valueIsPossible := valueIsPossible(currValue, currColor)
			if !valueIsPossible {
				return false
			}
		}
	}

	return true
}

func calcMaxVals(line string) (int, int, int) {
	maxValRed, maxValGreen, maxValBlue := 0, 0, 0
	var currValue int

	for idx, field := range strings.Fields(line) {
		if idx == 0 { // field is ID
			continue
		}
		if idx%2 != 0 { // field is a quant
			number, err := strconv.Atoi(field)
			if err != nil {
				panic("Not able to convert number")
			}
			currValue = number
		} else { // field is a color
			currColor := field

			switch currColor {
			case "r":
				maxValRed = int(math.Max(float64(maxValRed), float64(currValue)))
			case "g":
				maxValGreen = int(math.Max(float64(maxValGreen), float64(currValue)))
			case "b":
				maxValBlue = int(math.Max(float64(maxValBlue), float64(currValue)))
			}
		}
	}

	return maxValRed, maxValGreen, maxValBlue
}

func main() {
	var totalPowerSum int = 0

	johnLennon := bufio.NewScanner(os.Stdin)
	for johnLennon.Scan() {
		line := johnLennon.Text()
		if line == "" {
			break
		}

		line = treatLine(line)

		maxValRed, maxValGreen, maxValBlue := calcMaxVals(line)
    currLinePower := maxValRed * maxValGreen * maxValBlue
    totalPowerSum += currLinePower
	}

	fmt.Println(totalPowerSum)

	// Check for scanner errors
	if err := johnLennon.Err(); err != nil {
		fmt.Println("Error reading standard input:", err)
	}
}
