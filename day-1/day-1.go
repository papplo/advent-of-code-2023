package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func joinDigits(digits []int) (int, error) {
	var digitsAsValue string
	var err error

	for _, digit := range digits {
		digitsAsValue += strconv.Itoa(digit)
	}

	if len(digits) < 1 {
		err = errors.New("no digits found in input")
	}

	result, error := strconv.Atoi(digitsAsValue)
	if error != nil {
		err = errors.New("something else went wrong")
	}

	return result, err
}

func numeralDigit(d string) string {
	switch d {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return "0"
	}
}

func toNumeralDigits(incomingString string) []string {
	possibleDigits := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	var result string = incomingString

	for _, numeral := range possibleDigits {
		findIndex := strings.Index(result, numeral)
		if findIndex > -1 {
			result = strings.Replace(result, numeral, numeralDigit(numeral), -1)
		}
	}

	return strings.Split(result, "")
}

func findFirstAndLastDigit(line string) ([]int, error) {
	var firstLastDigit = make([]int, 2)
	var error error

	//* adjust to take digits as litteral digits */
	numeralCharacters := toNumeralDigits(line)

	for _, lineCharacter := range numeralCharacters {
		digit, err := strconv.Atoi(lineCharacter)
		if err == nil {
			if firstLastDigit[0] == 0 {
				firstLastDigit[0] = digit
			}
			firstLastDigit[1] = digit
		}
		if err != nil {
			error = err
		}
	}
	return firstLastDigit, error
}

func main() {
	// read input
	file, err := os.ReadFile("input.txt")
	if err != nil {
		return
	}

	// declare sum of calculation
	var sumOfAllValues int

	// loop the whole range
	for _, line := range strings.Split(string(file), "\n") {
		mergeLineDigits, _ := findFirstAndLastDigit(line)
		sumOfLine, err := joinDigits(mergeLineDigits)
		if err == nil {
			fmt.Println(sumOfLine)
			sumOfAllValues += sumOfLine
		}

	}

	// write result
	output, err := os.Create("output.txt")
	if err != nil {
		return
	}
	defer output.Close()
	output.WriteString(fmt.Sprint(sumOfAllValues))
	fmt.Printf("AoC: Day 1, part 1: %d", sumOfAllValues)
}
