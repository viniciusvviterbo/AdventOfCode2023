package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var writtenOutNumbers map[string]string

func init() {
	writtenOutNumbers = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
}

const maxNumberLength = 5

func convertWholeLine(line string) string {
	var newLine string

	for startWord, endWord := 0, 0; startWord < len(line) && endWord + 1 <= len(line); {
    fmt.Println(startWord, endWord)
		if unicode.IsDigit(rune(line[endWord])) {
			newLine = fmt.Sprintf("%s%c", newLine, rune(line[endWord]))
			endWord++
			startWord = endWord
		} else {
			val, ok := writtenOutNumbers[string(line[startWord:endWord+1])]
			if ok {
				newLine = newLine + val
				startWord = endWord + 1
			}
			endWord++
		}
	}

	fmt.Println(newLine)

	return newLine
}

func getNumerals(line string) string {
	var fullNumber string

	line = convertWholeLine(line)

	for _, char := range line {
		if unicode.IsDigit(char) {
			fullNumber = fmt.Sprintf("%s%c", fullNumber, char)
		}
	}

	return fmt.Sprintf("%c%c", fullNumber[0], fullNumber[len(fullNumber)-1])
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
