package main

import (
	"bufio"
	"fmt"
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

func main() {
  var totalIds int = 0

	johnLennon := bufio.NewScanner(os.Stdin)
	for johnLennon.Scan() {
		line := johnLennon.Text()
		if line == "" {
			break
		}

		line = treatLine(line)

		var lineIsPossible bool = lineIsPossible(line)

    if lineIsPossible {
      possibleLineId, err := strconv.Atoi(strings.Fields(line)[0])
			if err != nil {
				panic("Not able to convert number")
			}
      totalIds += possibleLineId
    }

	}

  fmt.Println(totalIds)

	// Check for scanner errors
	if err := johnLennon.Err(); err != nil {
		fmt.Println("Error reading standard input:", err)
	}
}
