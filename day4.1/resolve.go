package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const TotalRed = 12
const TotalGreen = 13
const TotalBlue = 14

func treatLine(line string) string {
	re := regexp.MustCompile(`(Card \d+:)`)

	newString := re.ReplaceAllString(line, "")

	return newString
}

func commonNumbers(winningNumbers []string, cardNumbers []string) []int {
	uniqueElements := make(map[string]struct{})

	for _, num := range winningNumbers {
		uniqueElements[num] = struct{}{}
	}

	var commonElements []int

	for _, num := range cardNumbers {
		if _, exists := uniqueElements[num]; exists {
			convertedNum, err := strconv.Atoi(num)
			if err != nil {
				panic("Not able to convert number")
			}
			commonElements = append(commonElements, convertedNum)
		}
	}

	return commonElements
}

func getCardValue(line string) int {
	bothCards := strings.Split(line, "|")
	winningNumbers := strings.Fields(bothCards[0])
	cardNumbers := strings.Fields(bothCards[1])

	commonNumbers := commonNumbers(winningNumbers, cardNumbers)
	numCommonNumbers := len(commonNumbers)

	var cardValue int
	if numCommonNumbers > 0 {
		cardValue = 1
	}

	cardValue = int(float64(cardValue) * math.Pow(float64(2), float64(numCommonNumbers-1)))

	return cardValue
}
func main() {
	var totalPoints int

	johnLennon := bufio.NewScanner(os.Stdin)
	for johnLennon.Scan() {
		line := johnLennon.Text()
		if line == "" {
			break
		}

		line = treatLine(line)

		cardValue := getCardValue(line)
		totalPoints += cardValue
	}

	fmt.Println(totalPoints)

	// Check for scanner errors
	if err := johnLennon.Err(); err != nil {
		fmt.Println("Error reading standard input:", err)
	}
}
