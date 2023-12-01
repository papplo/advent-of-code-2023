package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
		var mergeLineDigits = make([]int, 2)

		for _, lineCharacter := range strings.Split(line, "") {
			digit, err := strconv.Atoi(lineCharacter)
			if err == nil {
				// digit is any digit apparent on line
				if mergeLineDigits[0] == 0 {
					mergeLineDigits[0] = digit
				}
				mergeLineDigits[1] = digit
			}
		}

		fmt.Println(mergeLineDigits[0], mergeLineDigits[1])

		// todo merge and sum
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
