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

func findFirstAndLastDigit(line string) ([]int, error) {
	var firstDigit int
	var lastDigit int

	possibleDigits := []string{
		"zero",
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

	for i := 0; i < len(line); i++ {
		chars := strings.Split(line, "")

		// check if character at index is parsable to an integer
		char, err := strconv.Atoi(chars[i])
		if err == nil {
			if firstDigit == 0 {
				firstDigit = char
			}
			lastDigit = char
		}

		// check if character indexes as first letter in numeral
		for numeralInt, numeral := range possibleDigits {
			findIndex := strings.Index(strings.Join(chars[i:], ""), numeral)
			if findIndex == 0 {
				// found a numeral value at current loop
				if firstDigit == 0 {
					firstDigit = numeralInt
				}
				lastDigit = numeralInt
			}
		}
	}

	res := []int{firstDigit, lastDigit}

	return res, errors.New("")

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
