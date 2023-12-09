package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func getNumerals(line string) string {
	var fullNumber string

	for _, char := range line {
		if unicode.IsDigit(char) {
			fullNumber = fmt.Sprintf("%s%c", fullNumber, char)
		}
	}

  return fmt.Sprintf("%c%c", fullNumber[0], fullNumber[len(fullNumber) - 1])
}

func main() {
  totalSum := 0

	johnLennon := bufio.NewScanner(os.Stdin)
	for johnLennon.Scan() {
		line := johnLennon.Text()

    if line == "" {
      break
    }

    var numerals string = getNumerals(line)

    number, err := strconv.Atoi(numerals)
    if err != nil {
      panic("Not able to convert number")
    }

    totalSum += number
	}

  fmt.Println(totalSum)

	// Check for scanner errors
	if err := johnLennon.Err(); err != nil {
		fmt.Println("Error reading standard input:", err)
	}
}

